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

type TeleponDocHandler struct {
	teleponDoc     *repo2.TeleponCollRepository
	teleponDocRepo *repo.TeleponDocCollRepository
}

func NewTeleponDocAPIHandler(e *echo.Echo, db *mongo.Database) *TeleponDocHandler {
	h := &TeleponDocHandler{
		teleponDoc:     repo2.NewTeleponRepository(db),
		teleponDocRepo: repo.NewTeleponDocRepository(db),
	}

	group := e.Group("/api", context.Handler)

	group.GET("/doc/telepons", h.findAll)
	group.GET("/doc/telepon/:id", h.findByID)

	group.POST("/doc/telepon", h.create)

	group.PUT("/doc/telepon/:id", h.update)

	group.DELETE("/doc/telepon/:id", h.delete)

	return h
}

// findAll
// @Tags Doc Telepon
// @Summary Get all telepons documents
// @ID get-all-telepons-documents
// @Security ApiKeyAuth
// @Router /api/doc/telepons [GET]
// @Produce json
// @Success 200
func (h *TeleponDocHandler) findAll(c echo.Context) error {
	teleponDocs, err := h.teleponDocRepo.FindAll()
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get teleponDocs: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "Telepon Docs not found")
	}
	return c.JSON(http.StatusOK, teleponDocs)
}

// findByID
// @Tags Doc Telepon
// @Summary Get telepon document by ID
// @ID get-telepon-document-by-id
// @Security ApiKeyAuth
// @Router /api/doc/telepon/{id} [GET]
// @Produce json
// @Param id path string true "Telepon Document ID"
// @Success 200
func (h *TeleponDocHandler) findByID(c echo.Context) error {
	id := c.Param("id")

	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to convert id to ObjectID: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	teleponDoc, err := h.teleponDocRepo.FindOneByID(oId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusNotFound, "Telepon Doc not found")
		}
		log.Errorf("Failed to get teleponDoc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, teleponDoc)
}

// create
// @Tags Doc Telepon
// @Summary Create new telepon document
// @ID create-new-telepon-document
// @Security ApiKeyAuth
// @Router /api/doc/telepon [POST]
// @Accept json
// @Produce json
// @Param body body doc.DeviceDocForm true "Telepon Document Form"
// @Success 200
func (h *TeleponDocHandler) create(c echo.Context) error {
	nc := c.(*context.Context)

	f, err := doc.NewDeviceDocForm(c)
	if err != nil {
		return err
	}

	telepon, err := h.teleponDoc.FindOneByID(f.DeviceOID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusNotFound, "Telepon not found")
		}
		log.Errorf("Failed to get telepon: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	teleponDoc := &repo.TeleponDoc{
		ID:         bson.NewObjectID(),
		Lokasi:     telepon.Lokasi,
		Departemen: telepon.Departemen,
		User:       telepon.User,
		Ext:        telepon.Ext,
		Merk:       telepon.Merk,
		Tipe:       telepon.Tipe,
		Checkpoint: f.Checkpoint,
		Inserted:   nc.Claims.ByAt(),
		IsDeleted:  false,
	}

	err = h.teleponDocRepo.InsertOne(teleponDoc)
	if err != nil {
		log.Errorf("Failed to create teleponDoc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, teleponDoc)
}

// update
// @Tags Doc Telepon
// @Summary Update telepon document by ID
// @ID update-telepon-document-by-id
// @Security ApiKeyAuth
// @Router /api/doc/telepon/{id} [PUT]
// @Produce json
// @Param id path string true "Telepon Document ID"
// @Param body body doc.UpdateDeviceDocForm true "Telepon Document Form"
// @Success 200
func (h *TeleponDocHandler) update(c echo.Context) error {
	nc := c.(*context.Context)

	id := c.Param("id")
	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to update telepon: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid telepon ID")
	}

	f, err := doc.NewUpdateDeviceDocForm(c)
	if err != nil {
		return err
	}

	teleponDoc, err := h.teleponDocRepo.FindOneByID(oId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusNotFound, "Telepon Doc not found")
		}
		log.Errorf("Failed to update teleponDoc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	teleponDoc.Checkpoint = f.Checkpoint
	teleponDoc.Updated = nc.Claims.ByAtPtr()
	err = h.teleponDocRepo.UpdateOneByID(oId, teleponDoc)
	if err != nil {
		log.Errorf("Failed to update teleponDoc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, teleponDoc)
}

// delete
// @Tags Doc Telepon
// @Summary Delete telepon document by ID
// @ID delete-telepon-document-by-id
// @Security ApiKeyAuth
// @Router /api/doc/telepon/{id} [DELETE]
// @Param id path string true "Telepon Document ID"
// @Success 200
func (h *TeleponDocHandler) delete(c echo.Context) error {
	id := c.Param("id")

	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to convert id to ObjectID: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	err = h.teleponDocRepo.DeleteOneByID(oId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusNotFound, "Telepon Doc not found")
		}
		log.Errorf("Failed to delete teleponDoc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, "Telepon Doc deleted")
}
