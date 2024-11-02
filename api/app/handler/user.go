package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
	"sipamit-be/api/app/repo"
	"sipamit-be/internal/pkg/context"
	"sipamit-be/internal/pkg/doc"
	"sipamit-be/internal/pkg/log"
	"sipamit-be/internal/pkg/util"
)

type userForm struct {
	FullName string `form:"full_name" json:"full_name"`
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func newUserForm(c echo.Context) (*userForm, error) {
	f := new(userForm)
	if err := c.Bind(f); err != nil {
		log.Errorf("Failed to bind user form: %v", err)
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	if f.FullName == "" && f.Username == "" && f.Password == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Please fill provided field")
	}

	return f, nil
}

type UserHandler struct {
	userRepo *repo.UserCollRepository
}

func NewUserHandler(e *echo.Echo, db *mongo.Database) *UserHandler {
	u := &UserHandler{
		userRepo: repo.NewUserRepository(db),
	}

	group := e.Group("/api", context.Handler, context.SuperAdminOnly)

	group.POST("/user", u.create)
	group.GET("/users", u.find)
	group.GET("/user/:username", u.detail)
	group.PUT("/user/:username", u.editUserProfile)
	group.DELETE("/user/:username", u.delete)

	return u
}

// create
// @Tags User
// @Summary Add new admin user
// @ID user-create
// @Security ApiKeyAuth
// @Router /api/user [POST]
// @Param body body userForm true "User Form"
// @Produce json
// @Success 200
func (h *UserHandler) create(c echo.Context) error {
	nc := c.(*context.Context)

	if !nc.Claims.IsSuperAdmin() {
		return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
	}

	f, err := newUserForm(c)
	if err != nil {
		return err
	}

	if f.FullName == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Full Name is required")
	}
	if f.Username == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Username is required")
	}
	if f.Password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Password is required")
	}

	_, err = h.userRepo.FindByUsername(f.Username)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to find user by username: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusBadRequest, "Username already exists")
	}

	user := &repo.User{
		ID:        bson.NewObjectID(),
		FullName:  f.FullName,
		Username:  f.Username,
		Password:  util.CryptPassword(f.Password),
		Role:      doc.AdminRole,
		Inserted:  nc.Claims.ByAt(),
		IsDeleted: false,
	}

	err = h.userRepo.InsertOne(user)
	if err != nil {
		log.Errorf("Failed to insert user: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, user)
}

// find
// @Tags User
// @Summary Get all users
// @ID user-find
// @Security ApiKeyAuth
// @Param q query string false "Search by fullname"
// @Param page query int false "Page number pagination" default(1)
// @Param limit query int false "Limit pagination" default(10)
// @Param sort query string false "Sort" enums(asc,desc)
// @Router /api/users [GET]
// @Produce json
// @Success 200
func (h *UserHandler) find(c echo.Context) error {
	nc := c.(*context.Context)

	if !nc.Claims.IsSuperAdmin() {
		return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
	}

	cq := util.NewCommonQuery(c)

	users, err := h.userRepo.FindAll(cq)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to find all users: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "Users not found")
	}

	totalUser, err := h.userRepo.CountQuery(cq)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to count users: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "Users not found")
	}

	result := util.MakeResult(users, totalUser, cq.Page, cq.Limit)
	return c.JSON(http.StatusOK, result)
}

// detail
// @Tags User
// @Summary Get user detail
// @ID user-detail
// @Security ApiKeyAuth
// @Router /api/user/{username} [GET]
// @Param username path string true "username"
// @Produce json
// @Success 200
func (h *UserHandler) detail(c echo.Context) error {
	nc := c.(*context.Context)

	if !nc.Claims.IsSuperAdmin() {
		return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
	}

	username := c.Param("username")
	if username == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Username is required")
	}

	user, err := h.userRepo.FindByUsername(username)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to find user by username: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}
	return c.JSON(http.StatusOK, user)
}

// editUserProfile
// @Tags User
// @Summary Edit user profile
// @ID user-edit
// @Security ApiKeyAuth
// @Router /api/user/{username} [PUT]
// @Param username path string true "username"
// @Param body body userForm true "User Form"
// @Produce json
// @Success 200
func (h *UserHandler) editUserProfile(c echo.Context) error {
	nc := c.(*context.Context)

	username := c.Param("username")
	if username == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Username is required")
	}

	if !nc.Claims.IsSuperAdmin() && nc.Claims.Username != username {
		return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
	}

	f, err := newUserForm(c)
	if err != nil {
		return err
	}

	user, err := h.userRepo.FindByUsername(username)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to find user by username: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	if f.FullName != "" {
		user.FullName = f.FullName
	}

	if f.Username != "" {
		_, err = h.userRepo.FindByUsername(f.Username)
		if err != nil {
			if !errors.Is(err, mongo.ErrNoDocuments) {
				log.Errorf("Failed to find user by username: %v", err)
				return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
			}
			user.Username = f.Username
		} else {
			return echo.NewHTTPError(http.StatusBadRequest, "Username already exists")
		}
	}

	if f.Password != "" {
		user.Password = util.CryptPassword(f.Password)
	}

	user.Updated = nc.Claims.ByAtPtr()
	err = h.userRepo.UpdateOne(user)
	if err != nil {
		log.Errorf("Failed to update user: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, user)
}

// delete
// @Tags User
// @Summary Delete user
// @ID user-delete
// @Security ApiKeyAuth
// @Router /api/user/{username} [DELETE]
// @Param username path string true "username"
// @Produce json
// @Success 200
func (h *UserHandler) delete(c echo.Context) error {
	nc := c.(*context.Context)

	if !nc.Claims.IsSuperAdmin() {
		return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
	}

	username := c.Param("username")
	if username == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Username is required")
	}

	user, err := h.userRepo.FindByUsername(username)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to find user by username: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	user.IsDeleted = true
	err = h.userRepo.UpdateOne(user)
	if err != nil {
		log.Errorf("Failed to delete user: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, user)
}
