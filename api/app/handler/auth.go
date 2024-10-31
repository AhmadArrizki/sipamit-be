package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
	"sipamit-be/api/app/repo"
	"sipamit-be/internal/pkg/context"
	"sipamit-be/internal/pkg/log"
	"sipamit-be/internal/pkg/util"
)

type loginForm struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func newLoginForm(c echo.Context) (*loginForm, error) {
	f := new(loginForm)
	if err := c.Bind(f); err != nil {
		log.Errorf("Failed to bind login form: %v", err)
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	if f.Username == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Username is required")
	}
	if f.Password == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Password is required")
	}
	return f, nil
}

type AuthHandler struct {
	userRepo *repo.UserCollRepository
}

func NewAuthHandler(e *echo.Echo, db *mongo.Database) *AuthHandler {
	h := &AuthHandler{
		userRepo: repo.NewUserRepository(db),
	}

	e.POST("/api/login", h.login)

	return h
}

// login
// @Tags Auth
// @Summary Login
// @ID login
// @Router /api/login [POST]
// @Param body body loginForm true "Login Form"
// @Produce json
// @Success 200
func (h *AuthHandler) login(c echo.Context) error {
	f, err := newLoginForm(c)
	if err != nil {
		return err
	}

	user, err := h.userRepo.FindByUsername(f.Username)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to find user by username: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusBadRequest, "Wrong username/email or password")
	}

	if !util.CheckPassword(user.Password, f.Password) {
		return echo.NewHTTPError(http.StatusBadRequest, "Wrong username/email or password")
	}

	token, err := context.MakeToken(user)
	if err != nil {
		log.Errorf("Failed to make token: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
