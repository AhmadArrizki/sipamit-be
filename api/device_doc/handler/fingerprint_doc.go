package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
	repo2 "sipamit-be/api/device/repo"
	"sipamit-be/api/device_doc/repo"
	"sipamit-be/internal/pkg/context"
	"sipamit-be/internal/pkg/doc"
	"sipamit-be/internal/pkg/log"
	"sipamit-be/internal/pkg/util"
)

type FingerprintDocHandler struct {
	fpRepo    *repo2.FingerPrintCollRepository
	fpDocRepo *repo.FingerprintDocCollRepository
}

func NewFingerprintDocAPIHandler(e *echo.Echo, db *mongo.Database) *FingerprintDocHandler {
	h := &FingerprintDocHandler{
		fpRepo:    repo2.NewFingerPrintRepository(db),
		fpDocRepo: repo.NewFingerprintDocRepository(db),
	}

	group := e.Group("/api", context.Handler)

	group.GET("/doc/fingerprints", h.findAll)
	group.GET("/doc/fingerprint/:id", h.findByID)

	group.POST("/doc/fingerprint", h.create)

	group.PUT("/doc/fingerprint/:id", h.update)

	group.DELETE("/doc/fingerprint/:id", h.delete)

	return h
}

// findAll
// @Tags Doc Fingerprint
// @Summary Get all fingerprint documents
// @ID get-all-fingerprint-documents
// @Security ApiKeyAuth
// @Param q query string false "Search by nama"
// @Param page query int false "Page number pagination" default(1)
// @Param limit query int false "Limit pagination" default(10)
// @Param sort query string false "Sort" enums(asc,desc)
// @Router /api/doc/fingerprints [GET]
// @Produce json
// @Success 200
func (h *FingerprintDocHandler) findAll(c echo.Context) error {
	cq := util.NewCommonQuery(c)

	fpDocs, err := h.fpDocRepo.FindAll(cq)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get fingerprintDocs: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "Fingerprint Docs not found")
	}

	totalFpDocs, err := h.fpDocRepo.CountQuery(cq)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to count fingerprintDocs: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "Fingerprint Docs not found")
	}

	result := util.MakeResult(fpDocs, totalFpDocs, cq.Page, cq.Limit)
	return c.JSON(http.StatusOK, result)
}

// findByID
// @Tags Doc Fingerprint
// @Summary Get fingerprint document by ID
// @ID get-fingerprint-document-by-id
// @Security ApiKeyAuth
// @Router /api/doc/fingerprint/{id} [GET]
// @Produce json
// @Param id path string true "Fingerprint Document ID"
// @Success 200
func (h *FingerprintDocHandler) findByID(c echo.Context) error {
	id := c.Param("id")

	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to convert id to ObjectID: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	fpDoc, err := h.fpDocRepo.FindOneByID(oId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusNotFound, "Fingerprint Doc not found")
		}
		log.Errorf("Failed to get fingerprintDoc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, fpDoc)
}

// create
// @Tags Doc Fingerprint
// @Summary Create fingerprint document
// @ID create-fingerprint-document
// @Security ApiKeyAuth
// @Router /api/doc/fingerprint [POST]
// @Produce json
// @Param body body doc.DeviceDocForm true "Fingerprint Document Form"
// @Success 200
func (h *FingerprintDocHandler) create(c echo.Context) error {
	nc := c.(*context.Context)

	f, err := doc.NewDeviceDocForm(c)
	if err != nil {
		return err
	}

	fp, err := h.fpRepo.FindOneByID(f.DeviceOID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusNotFound, "Fingerprint not found")
		}
		log.Errorf("Failed to get fingerprint: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	fpDoc := &repo.FingerprintDoc{
		ID:         bson.NewObjectID(),
		Nama:       fp.Nama,
		Lokasi:     fp.Lokasi,
		Kode:       fp.Kode,
		Checkpoint: f.Checkpoint,
		Inserted:   nc.Claims.ByAt(),
		IsDeleted:  false,
	}

	err = h.fpDocRepo.InsertOne(fpDoc)
	if err != nil {
		log.Errorf("Failed to create fingerprintDoc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, "Fingerprint Doc created")
}

// update
// @Tags Doc Fingerprint
// @Summary Update fingerprint document by ID
// @ID update-fingerprint-document-by-id
// @Security ApiKeyAuth
// @Router /api/doc/fingerprint/{id} [PUT]
// @Produce json
// @Param id path string true "Fingerprint Document ID"
// @Param body body doc.UpdateDeviceDocForm true "Fingerprint Document Form"
// @Success 200
func (h *FingerprintDocHandler) update(c echo.Context) error {
	nc := c.(*context.Context)

	id := c.Param("id")
	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to update fingerprint: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid fingerprint ID")
	}

	f, err := doc.NewUpdateDeviceDocForm(c)
	if err != nil {
		return err
	}

	fpDoc, err := h.fpDocRepo.FindOneByID(oId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusNotFound, "Fingerprint Doc not found")
		}
		log.Errorf("Failed to update fpDoc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	fpDoc.Checkpoint = f.Checkpoint
	fpDoc.Updated = nc.Claims.ByAtPtr()
	err = h.fpDocRepo.UpdateOneByID(oId, fpDoc)
	if err != nil {
		log.Errorf("Failed to update fpDoc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, fpDoc)
}

// delete
// @Tags Doc Fingerprint
// @Summary Delete fingerprint document by ID
// @ID delete-fingerprint-document-by-id
// @Security ApiKeyAuth
// @Router /api/doc/fingerprint/{id} [DELETE]
// @Produce json
// @Param id path string true "Fingerprint Document ID"
// @Success 200
func (h *FingerprintDocHandler) delete(c echo.Context) error {
	id := c.Param("id")

	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to convert id to ObjectID: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	err = h.fpDocRepo.DeleteOneByID(oId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusNotFound, "Fingerprint Doc not found")
		}
		log.Errorf("Failed to delete fingerprintDoc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, "Fingerprint Doc deleted")
}
