package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
	"sipamit-be/api/device/repo"
	"sipamit-be/internal/pkg/context"
	"sipamit-be/internal/pkg/log"
)

type teleponForm struct {
	Lokasi     string `form:"lokasi" json:"lokasi"`
	Departemen string `form:"departemen" json:"departemen"`
	User       string `form:"user" json:"user"`
	Ext        string `form:"ext" json:"ext"`
	Merk       string `form:"merk" json:"merk"`
	Tipe       string `form:"tipe" json:"tipe"`
}

func newTeleponForm(c echo.Context) (*teleponForm, error) {
	f := new(teleponForm)
	if err := c.Bind(f); err != nil {
		log.Errorf("Failed to bind telepon form: %v", err)
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	if f.Lokasi == "" && f.Departemen == "" && f.User == "" && f.Ext == "" && f.Merk == "" && f.Tipe == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Please fill provided field")
	}

	return f, nil
}

type TeleponHandler struct {
	teleponRepo *repo.TeleponCollRepository
}

func NewTeleponAPIHandler(e *echo.Echo, db *mongo.Database) *TeleponHandler {
	h := &TeleponHandler{
		teleponRepo: repo.NewTeleponRepository(db),
	}

	group := e.Group("/api", context.Handler)

	group.GET("/telepons", h.findAll)
	group.GET("/telepon/:id", h.findOne)

	group.POST("/telepon", h.create)

	group.PUT("/telepon/:id", h.update)

	group.DELETE("/telepon/:id", h.delete)

	return h
}

// findAll
// @Tags Device Telepon
// @Summary Get all telepons
// @ID get-all-telepons
// @Security ApiKeyAuth
// @Router /api/telepons [GET]
// @Produce json
// @Success 200
func (h *TeleponHandler) findAll(c echo.Context) error {
	telepons, err := h.teleponRepo.FindAll()
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get telepons: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "Telepons not found")
	}
	return c.JSON(http.StatusOK, telepons)
}

// findOne
// @Tags Device Telepon
// @Summary Get telepon by id
// @ID get-telepon-by-id
// @Security ApiKeyAuth
// @Router /api/telepon/{id} [GET]
// @Produce json
// @Param id path string true "Telepon ID"
// @Success 200
func (h *TeleponHandler) findOne(c echo.Context) error {
	id := c.Param("id")

	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to get telepon: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid telepon ID")
	}

	telepon, err := h.teleponRepo.FindOneByID(oId)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get telepon: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "Telepon not found")
	}
	return c.JSON(http.StatusOK, telepon)
}

// create
// @Tags Device Telepon
// @Summary Create new telepon
// @ID create-new-telepon
// @Security ApiKeyAuth
// @Router /api/telepon [POST]
// @Produce json
// @Param body body teleponForm true "Telepon Form"
// @Success 200
func (h *TeleponHandler) create(c echo.Context) error {
	nc := c.(*context.Context)

	f, err := newTeleponForm(nc)
	if err != nil {
		return err
	}

	if f.Lokasi == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Lokasi is required")
	}
	if f.Departemen == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Departemen is required")
	}
	if f.User == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "User is required")
	}
	if f.Ext == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Ext is required")
	}
	if f.Merk == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Merk is required")
	}
	if f.Tipe == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Tipe is required")
	}

	telepon := &repo.Telepon{
		ID:         bson.NewObjectID(),
		Lokasi:     f.Lokasi,
		Departemen: f.Departemen,
		User:       f.User,
		Ext:        f.Ext,
		Merk:       f.Merk,
		Tipe:       f.Tipe,
		Inserted:   nc.Claims.ByAt(),
		IsDeleted:  false,
	}

	err = h.teleponRepo.InsertOne(telepon)
	if err != nil {
		log.Errorf("Failed to create telepon: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, telepon)
}

// update
// @Tags Device Telepon
// @Summary Update telepon by id
// @ID update-telepon-by-id
// @Security ApiKeyAuth
// @Router /api/telepon/{id} [PUT]
// @Produce json
// @Param id path string true "Telepon ID"
// @Param body body teleponForm true "Telepon Form"
// @Success 200
func (h *TeleponHandler) update(c echo.Context) error {
	nc := c.(*context.Context)

	id := c.Param("id")
	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to update telepon: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid telepon ID")
	}

	f, err := newTeleponForm(nc)
	if err != nil {
		return err
	}

	telepon, err := h.teleponRepo.FindOneByID(oId)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to update telepon: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "Telepon not found")
	}

	if f.Lokasi != "" {
		telepon.Lokasi = f.Lokasi
	}
	if f.Departemen != "" {
		telepon.Departemen = f.Departemen
	}
	if f.User != "" {
		telepon.User = f.User
	}
	if f.Ext != "" {
		telepon.Ext = f.Ext
	}
	if f.Merk != "" {
		telepon.Merk = f.Merk
	}
	if f.Tipe != "" {
		telepon.Tipe = f.Tipe
	}

	telepon.Updated = nc.Claims.ByAtPtr()
	err = h.teleponRepo.UpdateOneByID(oId, telepon)
	if err != nil {
		log.Errorf("Failed to update telepon: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, telepon)
}

// delete
// @Tags Device Telepon
// @Summary Delete telepon by id
// @ID delete-telepon-by-id
// @Security ApiKeyAuth
// @Router /api/telepon/{id} [DELETE]
// @Produce json
// @Param id path string true "Telepon ID"
// @Success 200
func (h *TeleponHandler) delete(c echo.Context) error {
	id := c.Param("id")

	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to delete telepon: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid telepon ID")
	}

	telepon, _ := h.teleponRepo.FindOneByID(oId)
	if telepon == nil {
		return echo.NewHTTPError(http.StatusNotFound, "Telepon not found")
	}

	err = h.teleponRepo.DeleteOneByID(oId)
	if err != nil {
		log.Errorf("Failed to delete telepon: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, "Telepon deleted")
}
