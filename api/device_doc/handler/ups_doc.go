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

type UPSDocHandler struct {
	upsRepo    *repo2.UPSCollRepository
	upsDocRepo *repo.UPSDocCollRepository
}

func NewUPSDocAPIHandler(e *echo.Echo, db *mongo.Database) *UPSDocHandler {
	h := &UPSDocHandler{
		upsRepo:    repo2.NewUPSRepository(db),
		upsDocRepo: repo.NewUPSDocRepository(db),
	}

	group := e.Group("/api", context.Handler)

	group.GET("/doc/ups", h.findAll)
	group.GET("/doc/ups/:id", h.findByID)

	group.POST("/doc/ups", h.create)

	group.PUT("/doc/ups/:id", h.update)

	group.DELETE("/doc/ups/:id", h.delete)

	return h
}

// findAll
// @Tags Doc UPS
// @Summary Get all ups documents
// @ID get-all-ups-documents
// @Security ApiKeyAuth
// @Param q query string false "Search by nama"
// @Param page query int false "Page number pagination" default(1)
// @Param limit query int false "Limit pagination" default(10)
// @Param sort query string false "Sort" enums(asc,desc)
// @Router /api/doc/ups [GET]
// @Produce json
// @Success 200
func (h *UPSDocHandler) findAll(c echo.Context) error {
	cq := util.NewCommonQuery(c)

	upsDocs, err := h.upsDocRepo.FindAll(cq)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get upsDocs: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "UPS Docs not found")
	}

	totalUpsDocs, err := h.upsDocRepo.CountQuery(cq)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to count upsDocs: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "UPS Docs not found")
	}

	result := util.MakeResult(upsDocs, totalUpsDocs, cq.Page, cq.Limit)
	return c.JSON(http.StatusOK, result)
}

// findByID
// @Tags Doc UPS
// @Summary Get ups document by ID
// @ID get-ups-document-by-id
// @Security ApiKeyAuth
// @Router /api/doc/ups/{id} [GET]
// @Produce json
// @Param id path string true "UPS Document ID"
// @Success 200
func (h *UPSDocHandler) findByID(c echo.Context) error {
	id := c.Param("id")

	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to convert id to ObjectID: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	upsDoc, err := h.upsDocRepo.FindOneByID(oId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusNotFound, "UPS Doc not found")
		}
		log.Errorf("Failed to get upsDoc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, upsDoc)
}

// create
// @Tags Doc UPS
// @Summary Create new ups document
// @ID create-new-ups-document
// @Security ApiKeyAuth
// @Router /api/doc/ups [POST]
// @Produce json
// @Param body body doc.DeviceDocForm true "UPS Document Form"
// @Success 200
func (h *UPSDocHandler) create(c echo.Context) error {
	nc := c.(*context.Context)

	f, err := doc.NewDeviceDocForm(c)
	if err != nil {
		return err
	}

	ups, err := h.upsRepo.FindOneByID(f.DeviceOID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusNotFound, "UPS not found")
		}
		log.Errorf("Failed to get ups: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	upsDoc := &repo.UPSDoc{
		ID:         bson.NewObjectID(),
		Nama:       ups.Nama,
		Departemen: ups.Departemen,
		Tipe:       ups.Tipe,
		NoSeri:     ups.NoSeri,
		Lokasi:     ups.Lokasi,
		Checkpoint: f.Checkpoint,
		Inserted:   nc.Claims.ByAt(),
		IsDeleted:  false,
	}

	err = h.upsDocRepo.InsertOne(upsDoc)
	if err != nil {
		log.Errorf("Failed to create upsDoc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, upsDoc)
}

// update
// @Tags Doc UPS
// @Summary Update ups document by ID
// @ID update-ups-document-by-id
// @Security ApiKeyAuth
// @Router /api/doc/ups/{id} [PUT]
// @Produce json
// @Param id path string true "UPS Document ID"
// @Param body body doc.UpdateDeviceDocForm true "UPS Document Form"
// @Success 200
func (h *UPSDocHandler) update(c echo.Context) error {
	nc := c.(*context.Context)

	id := c.Param("id")
	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to update ups: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ups ID")
	}

	f, err := doc.NewUpdateDeviceDocForm(c)
	if err != nil {
		return err
	}

	upsDoc, err := h.upsDocRepo.FindOneByID(oId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusNotFound, "UPS Doc not found")
		}
		log.Errorf("Failed to update upsDoc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	upsDoc.Checkpoint = f.Checkpoint
	upsDoc.Updated = nc.Claims.ByAtPtr()
	err = h.upsDocRepo.UpdateOneByID(oId, upsDoc)
	if err != nil {
		log.Errorf("Failed to update upsDoc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, upsDoc)
}

// delete
// @Tags Doc UPS
// @Summary Delete ups document by ID
// @ID delete-ups-document-by-id
// @Security ApiKeyAuth
// @Router /api/doc/ups/{id} [DELETE]
// @Param id path string true "UPS Document ID"
// @Success 200
func (h *UPSDocHandler) delete(c echo.Context) error {
	id := c.Param("id")

	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to convert id to ObjectID: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	err = h.upsDocRepo.DeleteOneByID(oId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusNotFound, "UPS Doc not found")
		}
		log.Errorf("Failed to delete upsDoc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, "UPS Doc deleted")
}
