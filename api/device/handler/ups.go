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
	"sipamit-be/internal/pkg/util"
)

type upsForm struct {
	Nama       string `form:"nama" json:"nama"`
	Departemen string `form:"departemen" json:"departemen"`
	Tipe       string `form:"tipe" json:"tipe"`
	NoSeri     string `form:"no_seri" json:"no_seri"`
	Lokasi     string `form:"lokasi" json:"lokasi"`
}

func newUPSForm(c echo.Context) (*upsForm, error) {
	f := new(upsForm)
	if err := c.Bind(f); err != nil {
		log.Errorf("Failed to bind ups form: %v", err)
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	if f.Nama == "" && f.Departemen == "" && f.Tipe == "" && f.NoSeri == "" && f.Lokasi == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Please fill provided field")
	}

	return f, nil
}

type UPSHandler struct {
	upsRepo *repo.UPSCollRepository
}

func NewUPSAPIHandler(e *echo.Echo, db *mongo.Database) *UPSHandler {
	h := &UPSHandler{
		upsRepo: repo.NewUPSRepository(db),
	}

	group := e.Group("/api", context.Handler)

	group.GET("/ups", h.findAll)
	group.GET("/ups/:id", h.findOne)

	group.POST("/ups", h.create)

	group.PUT("/ups/:id", h.update)

	group.DELETE("/ups/:id", h.delete)

	return h
}

// findAll
// @Tags Device UPS
// @Summary Get all ups
// @ID get-all-ups
// @Security ApiKeyAuth
// @Param q query string false "Search by nama"
// @Param page query int false "Page number pagination" default(1)
// @Param limit query int false "Limit pagination" default(10)
// @Param sort query string false "Sort" enums(asc,desc)
// @Router /api/ups [GET]
// @Produce json
// @Success 200
func (h *UPSHandler) findAll(c echo.Context) error {
	cq := util.NewCommonQuery(c)

	ups, err := h.upsRepo.FindAll(cq)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get ups: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "UPS not found")
	}

	totalUps, err := h.upsRepo.CountQuery(cq)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to count ups: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "UPS not found")
	}

	result := util.MakeResult(ups, totalUps, cq.Page, cq.Limit)
	return c.JSON(http.StatusOK, result)
}

// findOne
// @Tags Device UPS
// @Summary Get ups by id
// @ID get-ups-by-id
// @Security ApiKeyAuth
// @Router /api/ups/{id} [GET]
// @Produce json
// @Param id path string true "UPS ID"
// @Success 200
func (h *UPSHandler) findOne(c echo.Context) error {
	id := c.Param("id")

	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to get ups: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ups ID")
	}

	ups, err := h.upsRepo.FindOneByID(oId)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get ups: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "UPS not found")
	}
	return c.JSON(http.StatusOK, ups)
}

// create
// @Tags Device UPS
// @Summary Create new ups
// @ID create-new-ups
// @Security ApiKeyAuth
// @Router /api/ups [POST]
// @Produce json
// @Param body body upsForm true "UPS Form"
// @Success 200
func (h *UPSHandler) create(c echo.Context) error {
	nc := c.(*context.Context)

	f, err := newUPSForm(c)
	if err != nil {
		return err
	}

	if f.Lokasi == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Lokasi is required")
	}
	if f.Departemen == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Departemen is required")
	}
	if f.Nama == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Nama is required")
	}
	if f.Tipe == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Tipe is required")
	}
	if f.NoSeri == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "No Seri is required")
	}

	ups := &repo.UPS{
		ID:         bson.NewObjectID(),
		Nama:       f.Nama,
		Departemen: f.Departemen,
		Tipe:       f.Tipe,
		NoSeri:     f.NoSeri,
		Lokasi:     f.Lokasi,
		Inserted:   nc.Claims.ByAt(),
		IsDeleted:  false,
	}

	err = h.upsRepo.InsertOne(ups)
	if err != nil {
		log.Errorf("Failed to create ups: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, ups)
}

// update
// @Tags Device UPS
// @Summary Update ups by id
// @ID update-ups-by-id
// @Security ApiKeyAuth
// @Router /api/ups/{id} [PUT]
// @Produce json
// @Param id path string true "UPS ID"
// @Param body body upsForm true "UPS Form"
// @Success 200
func (h *UPSHandler) update(c echo.Context) error {
	nc := c.(*context.Context)

	id := c.Param("id")
	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to update ups: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ups ID")
	}

	f, err := newUPSForm(c)
	if err != nil {
		return err
	}

	ups, err := h.upsRepo.FindOneByID(oId)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to update ups: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "UPS not found")
	}

	if f.Nama != "" {
		ups.Nama = f.Nama
	}
	if f.Departemen != "" {
		ups.Departemen = f.Departemen
	}
	if f.Tipe != "" {
		ups.Tipe = f.Tipe
	}
	if f.NoSeri != "" {
		ups.NoSeri = f.NoSeri
	}
	if f.Lokasi != "" {
		ups.Lokasi = f.Lokasi
	}

	ups.Updated = nc.Claims.ByAtPtr()
	err = h.upsRepo.UpdateOneByID(oId, ups)
	if err != nil {
		log.Errorf("Failed to update ups: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, ups)
}

// delete
// @Tags Device UPS
// @Summary Delete ups by id
// @ID delete-ups-by-id
// @Security ApiKeyAuth
// @Router /api/ups/{id} [DELETE]
// @Produce json
// @Param id path string true "UPS ID"
// @Success 200
func (h *UPSHandler) delete(c echo.Context) error {
	id := c.Param("id")

	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to delete ups: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ups ID")
	}

	ups, _ := h.upsRepo.FindOneByID(oId)
	if ups == nil {
		return echo.NewHTTPError(http.StatusNotFound, "UPS not found")
	}

	err = h.upsRepo.DeleteOneByID(oId)
	if err != nil {
		log.Errorf("Failed to delete ups: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, "UPS deleted")
}
