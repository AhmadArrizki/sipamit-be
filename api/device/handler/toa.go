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

type toaForm struct {
	Nama   string `form:"nama" json:"nama"`
	Lokasi string `form:"lokasi" json:"lokasi"`
	Kode   string `form:"kode" json:"kode"`
	Posisi string `form:"posisi" json:"posisi"`
}

func newTOAForm(c echo.Context) (*toaForm, error) {
	f := new(toaForm)
	if err := c.Bind(f); err != nil {
		log.Errorf("Failed to bind toa form: %v", err)
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	if f.Nama == "" && f.Lokasi == "" && f.Kode == "" && f.Posisi == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Please fill provided field")
	}

	return f, nil
}

type TOAHandler struct {
	toaRepo *repo.TOACollRepository
}

func NewTOAAPIHandler(e *echo.Echo, db *mongo.Database) *TOAHandler {
	h := &TOAHandler{
		toaRepo: repo.NewTOARepository(db),
	}

	group := e.Group("/api", context.Handler)

	group.GET("/toas", h.findAll)
	group.GET("/toa/:id", h.findOne)

	group.POST("/toa", h.create)

	group.PUT("/toa/:id", h.update)

	group.DELETE("/toa/:id", h.delete)

	return h
}

// findAll
// @Tags Device TOA
// @Summary Get all toas
// @ID get-all-toas
// @Security ApiKeyAuth
// @Param q query string false "Search by nama"
// @Param page query int false "Page number pagination" default(1)
// @Param limit query int false "Limit pagination" default(10)
// @Param sort query string false "Sort" enums(asc,desc)
// @Router /api/toas [GET]
// @Produce json
// @Success 200
func (h *TOAHandler) findAll(c echo.Context) error {
	cq := util.NewCommonQuery(c)

	toas, err := h.toaRepo.FindAll(cq)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get toas: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "TOAs not found")
	}

	totalToas, err := h.toaRepo.CountQuery(cq)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to count toas: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "TOAs not found")
	}

	result := util.MakeResult(toas, totalToas, cq.Page, cq.Limit)
	return c.JSON(http.StatusOK, result)
}

// findOne
// @Tags Device TOA
// @Summary Get toa by id
// @ID get-toa-by-id
// @Security ApiKeyAuth
// @Router /api/toa/{id} [GET]
// @Produce json
// @Param id path string true "TOA ID"
// @Success 200
func (h *TOAHandler) findOne(c echo.Context) error {
	id := c.Param("id")

	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to get toa: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid toa ID")
	}

	toa, err := h.toaRepo.FindOneByID(oId)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get toa: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "TOA not found")
	}
	return c.JSON(http.StatusOK, toa)
}

// create
// @Tags Device TOA
// @Summary Create new toa
// @ID create-new-toa
// @Security ApiKeyAuth
// @Router /api/toa [POST]
// @Produce json
// @Param body body toaForm true "TOA Form"
// @Success 200
func (h *TOAHandler) create(c echo.Context) error {
	nc := c.(*context.Context)

	f, err := newTOAForm(c)
	if err != nil {
		return err
	}

	if f.Nama == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Nama is required")
	}
	if f.Lokasi == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Lokasi is required")
	}
	if f.Kode == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Kode is required")
	}
	if f.Posisi == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Posisi is required")
	}

	toa := &repo.TOA{
		ID:        bson.NewObjectID(),
		Nama:      f.Nama,
		Lokasi:    f.Lokasi,
		Kode:      f.Kode,
		Posisi:    f.Posisi,
		Inserted:  nc.Claims.ByAt(),
		IsDeleted: false,
	}

	err = h.toaRepo.InsertOne(toa)
	if err != nil {
		log.Errorf("Failed to create toa: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, toa)
}

// update
// @Tags Device TOA
// @Summary Update toa by id
// @ID update-toa-by-id
// @Security ApiKeyAuth
// @Router /api/toa/{id} [PUT]
// @Produce json
// @Param id path string true "TOA ID"
// @Param body body toaForm true "TOA Form"
// @Success 200
func (h *TOAHandler) update(c echo.Context) error {
	nc := c.(*context.Context)

	id := c.Param("id")
	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to update toa: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid toa ID")
	}

	f, err := newTOAForm(c)
	if err != nil {
		return err
	}

	toa, err := h.toaRepo.FindOneByID(oId)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to update toa: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "TOA not found")
	}

	if f.Nama != "" {
		toa.Nama = f.Nama
	}
	if f.Lokasi != "" {
		toa.Lokasi = f.Lokasi
	}
	if f.Kode != "" {
		toa.Kode = f.Kode
	}
	if f.Posisi != "" {
		toa.Posisi = f.Posisi
	}

	toa.Updated = nc.Claims.ByAtPtr()
	err = h.toaRepo.UpdateOneByID(oId, toa)
	if err != nil {
		log.Errorf("Failed to update toa: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, toa)
}

// delete
// @Tags Device TOA
// @Summary Delete toa by id
// @ID delete-toa-by-id
// @Security ApiKeyAuth
// @Router /api/toa/{id} [DELETE]
// @Produce json
// @Param id path string true "TOA ID"
// @Success 200
func (h *TOAHandler) delete(c echo.Context) error {
	id := c.Param("id")

	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to delete toa: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid toa ID")
	}

	toa, _ := h.toaRepo.FindOneByID(oId)
	if toa == nil {
		return echo.NewHTTPError(http.StatusNotFound, "TOA not found")
	}

	err = h.toaRepo.DeleteOneByID(oId)
	if err != nil {
		log.Errorf("Failed to delete toa: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, "TOA deleted")
}
