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

type TOADocHandler struct {
	toaRepo    *repo2.TOACollRepository
	toaDocRepo *repo.TOADocCollRepository
}

func NewTOADocAPIHandler(e *echo.Echo, db *mongo.Database) *TOADocHandler {
	h := &TOADocHandler{
		toaRepo:    repo2.NewTOARepository(db),
		toaDocRepo: repo.NewTOADocRepository(db),
	}

	group := e.Group("/api", context.Handler)

	group.GET("/doc/toas", h.findAll)
	group.GET("/doc/toa/:id", h.findByID)

	group.POST("/doc/toa", h.create)

	group.PUT("/doc/toa/:id", h.update)

	group.DELETE("/doc/toa/:id", h.delete)

	return h
}

// findAll
// @Tags Doc TOA
// @Summary Get all toas documents
// @ID get-all-toas-documents
// @Security ApiKeyAuth
// @Param q query string false "Search by nama"
// @Param page query int false "Page number pagination" default(1)
// @Param limit query int false "Limit pagination" default(10)
// @Param sort query string false "Sort" enums(asc,desc)
// @Router /api/doc/toas [GET]
// @Produce json
// @Success 200
func (h *TOADocHandler) findAll(c echo.Context) error {
	cq := util.NewCommonQuery(c)

	toaDocs, err := h.toaDocRepo.FindAll(cq)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get toaDocs: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "TOA Docs not found")
	}

	totalToaDocs, err := h.toaDocRepo.CountQuery(cq)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to count toaDocs: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "TOA Docs not found")
	}

	result := util.MakeResult(toaDocs, totalToaDocs, cq.Page, cq.Limit)
	return c.JSON(http.StatusOK, result)
}

// findByID
// @Tags Doc TOA
// @Summary Get toa document by ID
// @ID get-toa-document-by-id
// @Security ApiKeyAuth
// @Router /api/doc/toa/{id} [GET]
// @Produce json
// @Param id path string true "TOA Document ID"
// @Success 200
func (h *TOADocHandler) findByID(c echo.Context) error {
	id := c.Param("id")

	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to convert id to ObjectID: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	toaDoc, err := h.toaDocRepo.FindOneByID(oId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusNotFound, "TOA Doc not found")
		}
		log.Errorf("Failed to get toaDoc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, toaDoc)
}

// create
// @Tags Doc TOA
// @Summary Create toa document
// @ID create-toa-document
// @Security ApiKeyAuth
// @Router /api/doc/toa [POST]
// @Produce json
// @Param body body doc.DeviceDocForm true "TOA Document Form"
// @Success 200
func (h *TOADocHandler) create(c echo.Context) error {
	nc := c.(*context.Context)

	f, err := doc.NewDeviceDocForm(c)
	if err != nil {
		return err
	}

	toa, err := h.toaRepo.FindOneByID(f.DeviceOID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusNotFound, "TOA not found")
		}
		log.Errorf("Failed to get toa: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	toaDoc := &repo.TOADoc{
		ID:         bson.NewObjectID(),
		Nama:       toa.Nama,
		Lokasi:     toa.Lokasi,
		Kode:       toa.Kode,
		Posisi:     toa.Posisi,
		Checkpoint: f.Checkpoint,
		Inserted:   nc.Claims.ByAt(),
		IsDeleted:  false,
	}

	err = h.toaDocRepo.InsertOne(toaDoc)
	if err != nil {
		log.Errorf("Failed to create toaDoc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, toaDoc)
}

// update
// @Tags Doc TOA
// @Summary Update toa document by ID
// @ID update-toa-document-by-id
// @Security ApiKeyAuth
// @Router /api/doc/toa/{id} [PUT]
// @Produce json
// @Param id path string true "TOA Document ID"
// @Param body body doc.UpdateDeviceDocForm true "TOA Document Form"
// @Success 200
func (h *TOADocHandler) update(c echo.Context) error {
	nc := c.(*context.Context)

	id := c.Param("id")
	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to update toa: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid toa ID")
	}

	f, err := doc.NewUpdateDeviceDocForm(c)
	if err != nil {
		return err
	}

	toaDoc, err := h.toaDocRepo.FindOneByID(oId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusNotFound, "TOA Doc not found")
		}
		log.Errorf("Failed to update toaDoc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	toaDoc.Checkpoint = f.Checkpoint
	toaDoc.Updated = nc.Claims.ByAtPtr()
	err = h.toaDocRepo.UpdateOneByID(oId, toaDoc)
	if err != nil {
		log.Errorf("Failed to update toaDoc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, toaDoc)
}

// delete
// @Tags Doc TOA
// @Summary Delete toa document by ID
// @ID delete-toa-document-by-id
// @Security ApiKeyAuth
// @Router /api/doc/toa/{id} [DELETE]
// @Param id path string true "TOA Document ID"
// @Success 200
func (h *TOADocHandler) delete(c echo.Context) error {
	id := c.Param("id")

	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to convert id to ObjectID: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	err = h.toaDocRepo.DeleteOneByID(oId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusNotFound, "TOA Doc not found")
		}
		log.Errorf("Failed to delete toaDoc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, "TOA Doc deleted")
}
