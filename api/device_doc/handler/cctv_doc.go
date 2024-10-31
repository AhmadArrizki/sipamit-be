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
)

type CCTVDocHandler struct {
	cctvRepo    *repo2.CCTVCollRepository
	cctvDocRepo *repo.CCTVDocCollRepository
}

func NewCCTVDocAPIHandler(e *echo.Echo, db *mongo.Database) *CCTVDocHandler {
	h := &CCTVDocHandler{
		cctvRepo:    repo2.NewCCTVRepository(db),
		cctvDocRepo: repo.NewCCTVDocRepository(db),
	}

	group := e.Group("/api", context.Handler)

	group.GET("/doc/cctvs", h.findAll)
	group.GET("/doc/cctv/:id", h.findByID)

	group.POST("/doc/cctv", h.create)

	group.PUT("/doc/cctv/:id", h.update)

	group.DELETE("/doc/cctv/:id", h.delete)

	return h
}

// findAll
// @Tags Doc CCTV
// @Summary Get all cctv documents
// @ID get-all-cctv-documents
// @Security ApiKeyAuth
// @Router /api/doc/cctvs [GET]
// @Produce json
// @Success 200
func (h *CCTVDocHandler) findAll(c echo.Context) error {
	cctvDocs, err := h.cctvDocRepo.FindAll()
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get cctvDocs: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "CCTV Docs not found")
	}
	return c.JSON(http.StatusOK, cctvDocs)
}

// findByID
// @Tags Doc CCTV
// @Summary Get cctv document by ID
// @ID get-cctv-document-by-id
// @Security ApiKeyAuth
// @Router /api/doc/cctv/{id} [GET]
// @Produce json
// @Param id path string true "CCTV Document ID"
// @Success 200
func (h *CCTVDocHandler) findByID(c echo.Context) error {
	id := c.Param("id")

	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to convert id to ObjectID: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	cctvDoc, err := h.cctvDocRepo.FindOneByID(oId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusNotFound, "CCTV Doc not found")
		}
		log.Errorf("Failed to get cctvDoc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, cctvDoc)
}

// create
// @Tags Doc CCTV
// @Summary Create cctv document
// @ID create-cctv-document
// @Security ApiKeyAuth
// @Router /api/doc/cctv [POST]
// @Produce json
// @Param body body doc.DeviceDocForm true "CCTV Document Form"
// @Success 200
func (h *CCTVDocHandler) create(c echo.Context) error {
	nc := c.(*context.Context)

	f, err := doc.NewDeviceDocForm(c)
	if err != nil {
		return err
	}

	cctv, err := h.cctvRepo.FindOneByID(f.DeviceOID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusNotFound, "CCTV not found")
		}
		log.Errorf("Failed to get cctv: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	cctvDoc := &repo.CCTVDoc{
		ID:         bson.NewObjectID(),
		Nama:       cctv.Nama,
		Lokasi:     cctv.Lokasi,
		Kode:       cctv.Kode,
		Checkpoint: f.Checkpoint,
		Inserted:   nc.Claims.ByAt(),
		IsDeleted:  false,
	}

	err = h.cctvDocRepo.InsertOne(cctvDoc)
	if err != nil {
		log.Errorf("Failed to create cctvDoc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, cctvDoc)
}

// update
// @Tags Doc CCTV
// @Summary Update cctv document by ID
// @ID update-cctv-document-by-id
// @Security ApiKeyAuth
// @Router /api/doc/cctv/{id} [PUT]
// @Produce json
// @Param id path string true "CCTV Document ID"
// @Param body body doc.UpdateDeviceDocForm true "CCTV Document Form"
// @Success 200
func (h *CCTVDocHandler) update(c echo.Context) error {
	nc := c.(*context.Context)

	id := c.Param("id")
	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to update cctv: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid cctv ID")
	}

	f, err := doc.NewUpdateDeviceDocForm(c)
	if err != nil {
		return err
	}

	cctvDoc, err := h.cctvDocRepo.FindOneByID(oId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusNotFound, "CCTV Doc not found")
		}
		log.Errorf("Failed to update cctvDoc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	cctvDoc.Checkpoint = f.Checkpoint
	cctvDoc.Updated = nc.Claims.ByAtPtr()
	err = h.cctvDocRepo.UpdateOneByID(oId, cctvDoc)
	if err != nil {
		log.Errorf("Failed to update cctvDoc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, cctvDoc)
}

// delete
// @Tags Doc CCTV
// @Summary Delete cctv document by ID
// @ID delete-cctv-document-by-id
// @Security ApiKeyAuth
// @Router /api/doc/cctv/{id} [DELETE]
// @Produce json
// @Param id path string true "CCTV Document ID"
// @Success 200
func (h *CCTVDocHandler) delete(c echo.Context) error {
	id := c.Param("id")

	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to convert id to ObjectID: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	err = h.cctvDocRepo.DeleteOneByID(oId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusNotFound, "CCTV Doc not found")
		}
		log.Errorf("Failed to delete cctvDoc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, "CCTV Doc deleted")
}
