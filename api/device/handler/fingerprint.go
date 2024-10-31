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

type fingerPrintForm struct {
	Nama   string `form:"nama" json:"nama"`
	Lokasi string `form:"lokasi" json:"lokasi" `
	Kode   string `form:"kode" json:"kode"`
}

func newFingerPrintForm(c echo.Context) (*fingerPrintForm, error) {
	f := new(fingerPrintForm)
	if err := c.Bind(f); err != nil {
		log.Errorf("Failed to bind fingerprint form: %v", err)
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	if f.Nama == "" && f.Lokasi == "" && f.Kode == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Please fill provided field")
	}

	return f, nil
}

type FingerPrintHandler struct {
	fpRepo *repo.FingerPrintCollRepository
}

func NewFingerPrintAPIHandler(e *echo.Echo, db *mongo.Database) *FingerPrintHandler {
	h := &FingerPrintHandler{
		fpRepo: repo.NewFingerPrintRepository(db),
	}

	group := e.Group("/api", context.Handler)

	group.GET("/fingerprints", h.findAll)
	group.GET("/fingerprint/:id", h.findOne)

	group.POST("/fingerprint", h.create)

	group.PUT("/fingerprint/:id", h.update)

	group.DELETE("/fingerprint/:id", h.delete)

	return h
}

// findAll
// @Tags Device Fingerprint
// @Summary Get all fingerprints
// @ID get-all-fingerprints
// @Security ApiKeyAuth
// @Router /api/fingerprints [GET]
// @Produce json
// @Success 200
func (h *FingerPrintHandler) findAll(c echo.Context) error {
	fingerprints, err := h.fpRepo.FindAll()
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get fingerprints: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "Fingerprints not found")
	}
	return c.JSON(http.StatusOK, fingerprints)
}

// findOne
// @Tags Device Fingerprint
// @Summary Get fingerprint by id
// @ID get-fingerprint-by-id
// @Security ApiKeyAuth
// @Router /api/fingerprint/{id} [GET]
// @Produce json
// @Param id path string true "Fingerprint ID"
// @Success 200
func (h *FingerPrintHandler) findOne(c echo.Context) error {
	id := c.Param("id")

	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to get fingerprint: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid fingerprint ID")
	}

	fp, err := h.fpRepo.FindOneByID(oId)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get fingerprint: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "Fingerprint not found")
	}
	return c.JSON(http.StatusOK, fp)
}

// create
// @Tags Device Fingerprint
// @Summary Create new fingerprint
// @ID create-fingerprint
// @Security ApiKeyAuth
// @Router /api/fingerprint [POST]
// @Produce json
// @Param body body fingerPrintForm true "Fingerprint Form"
// @Success 200
func (h *FingerPrintHandler) create(c echo.Context) error {
	nc := c.(*context.Context)

	f, err := newFingerPrintForm(c)
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

	fp := &repo.FingerPrint{
		ID:        bson.NewObjectID(),
		Nama:      f.Nama,
		Lokasi:    f.Lokasi,
		Kode:      f.Kode,
		Inserted:  nc.Claims.ByAt(),
		IsDeleted: false,
	}

	err = h.fpRepo.InsertOne(fp)
	if err != nil {
		log.Errorf("Failed to create fingerprint: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, fp)
}

// update
// @Tags Device Fingerprint
// @Summary Update fingerprint by id
// @ID update-fingerprint-by-id
// @Security ApiKeyAuth
// @Router /api/fingerprint/{id} [PUT]
// @Produce json
// @Param id path string true "Fingerprint ID"
// @Param body body fingerPrintForm true "Fingerprint Form"
// @Success 200
func (h *FingerPrintHandler) update(c echo.Context) error {
	nc := c.(*context.Context)

	id := c.Param("id")
	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to update fingerprint: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid fingerprint ID")
	}

	f, err := newFingerPrintForm(c)
	if err != nil {
		return err
	}

	fp, err := h.fpRepo.FindOneByID(oId)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to update fingerprint: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "Fingerprint not found")
	}

	if f.Nama != "" {
		fp.Nama = f.Nama
	}
	if f.Lokasi != "" {
		fp.Lokasi = f.Lokasi
	}
	if f.Kode != "" {
		fp.Kode = f.Kode
	}

	fp.Updated = nc.Claims.ByAtPtr()
	err = h.fpRepo.UpdateOneByID(oId, fp)
	if err != nil {
		log.Errorf("Failed to update fingerprint: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, fp)
}

// delete
// @Tags Device Fingerprint
// @Summary Delete fingerprint by id
// @ID delete-fingerprint-by-id
// @Security ApiKeyAuth
// @Router /api/fingerprint/{id} [DELETE]
// @Produce json
// @Param id path string true "Fingerprint ID"
// @Success 200
func (h *FingerPrintHandler) delete(c echo.Context) error {
	id := c.Param("id")

	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to delete fingerprint: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid fingerprint ID")
	}

	fp, _ := h.fpRepo.FindOneByID(oId)
	if fp == nil {
		return echo.NewHTTPError(http.StatusNotFound, "Fingerprint not found")
	}

	err = h.fpRepo.DeleteOneByID(oId)
	if err != nil {
		log.Errorf("Failed to delete fingerprint: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, "Fingerprint deleted")
}
