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

type KomputerPH1DocHandler struct {
	kph1Repo    *repo2.KomputerPH1CollRepository
	kph1DocRepo *repo.KomputerPH1DocCollRepository
}

func NewKomputerPH1DocAPIHandler(e *echo.Echo, db *mongo.Database) *KomputerPH1DocHandler {
	h := &KomputerPH1DocHandler{
		kph1Repo:    repo2.NewKomputerPH1Repository(db),
		kph1DocRepo: repo.NewKomputerPH1DocRepository(db),
	}

	group := e.Group("/api", context.Handler)

	group.GET("/doc/komputer-ph1s", h.findAll)
	group.GET("/doc/komputer-ph1/:id", h.findByID)

	group.POST("/doc/komputer-ph1", h.create)

	group.PUT("/doc/komputer-ph1/:id", h.update)

	group.DELETE("/doc/komputer-ph1/:id", h.delete)

	return h
}

// findAll
// @Tags Doc Komputer PH1
// @Summary Get all komputer ph1 documents
// @ID get-all-komputer-ph1-documents
// @Security ApiKeyAuth
// @Param q query string false "Search by nama"
// @Param page query int false "Page number pagination" default(1)
// @Param limit query int false "Limit pagination" default(10)
// @Param sort query string false "Sort" enums(asc,desc)
// @Router /api/doc/komputer-ph1s [GET]
// @Produce json
// @Success 200
func (h *KomputerPH1DocHandler) findAll(c echo.Context) error {
	cq := util.NewCommonQuery(c)

	kph1Doc, err := h.kph1DocRepo.FindAll(cq)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get kph1Doc: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "Komputer PH1 Docs not found")
	}

	totalKph1Docs, err := h.kph1DocRepo.CountQuery(cq)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to count kph1Doc: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "Komputer PH1 Docs not found")
	}

	result := util.MakeResult(kph1Doc, totalKph1Docs, cq.Page, cq.Limit)
	return c.JSON(http.StatusOK, result)
}

// findByID
// @Tags Doc Komputer PH1
// @Summary Get komputer ph1 document by ID
// @ID get-komputer-ph1-document-by-id
// @Security ApiKeyAuth
// @Router /api/doc/komputer-ph1/{id} [GET]
// @Produce json
// @Param id path string true "Komputer PH1 Document ID"
// @Success 200
func (h *KomputerPH1DocHandler) findByID(c echo.Context) error {
	id := c.Param("id")

	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to convert id to ObjectID: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	kph1Doc, err := h.kph1DocRepo.FindOneByID(oId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusNotFound, "Komputer PH1 Doc not found")
		}
		log.Errorf("Failed to get kph1Doc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, kph1Doc)
}

// create
// @Tags Doc Komputer PH1
// @Summary Create new komputer ph1 document
// @ID create-new-komputer-ph1-document
// @Security ApiKeyAuth
// @Router /api/doc/komputer-ph1 [POST]
// @Produce json
// @Param body body doc.DeviceDocForm true "Komputer PH1 Document Form"
// @Success 200
func (h *KomputerPH1DocHandler) create(c echo.Context) error {
	nc := c.(*context.Context)

	f, err := doc.NewDeviceDocForm(c)
	if err != nil {
		return err
	}

	kph1, err := h.kph1Repo.FindOneByID(f.DeviceOID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusNotFound, "Komputer PH1 not found")
		}
		log.Errorf("Failed to get kph1: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	kph1Doc := &repo.KomputerPH1Doc{
		ID:         bson.NewObjectID(),
		Nama:       kph1.Nama,
		Merk:       kph1.Merk,
		PC:         kph1.PC,
		Monitor:    kph1.Monitor,
		CPU:        kph1.CPU,
		RAM:        kph1.RAM,
		Internal:   kph1.Internal,
		Lokasi:     kph1.Lokasi,
		Checkpoint: f.Checkpoint,
		Inserted:   nc.Claims.ByAt(),
		IsDeleted:  false,
	}

	err = h.kph1DocRepo.InsertOne(kph1Doc)
	if err != nil {
		log.Errorf("Failed to create kph1Doc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, kph1Doc)
}

// update
// @Tags Doc Komputer PH1
// @Summary Update komputer ph1 document by ID
// @ID update-komputer-ph1-document-by-id
// @Security ApiKeyAuth
// @Router /api/doc/komputer-ph1/{id} [PUT]
// @Produce json
// @Param id path string true "Komputer PH1 Document ID"
// @Param body body doc.UpdateDeviceDocForm true "Komputer PH1 Document Form"
// @Success 200
func (h *KomputerPH1DocHandler) update(c echo.Context) error {
	nc := c.(*context.Context)

	id := c.Param("id")
	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to update komputer ph1: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid komputer ph1 ID")
	}

	f, err := doc.NewUpdateDeviceDocForm(c)
	if err != nil {
		return err
	}

	kph1Doc, err := h.kph1DocRepo.FindOneByID(oId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusNotFound, "Komputer PH1 Doc not found")
		}
		log.Errorf("Failed to update kph1Doc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	kph1Doc.Checkpoint = f.Checkpoint
	kph1Doc.Updated = nc.Claims.ByAtPtr()
	err = h.kph1DocRepo.UpdateOneByID(oId, kph1Doc)
	if err != nil {
		log.Errorf("Failed to update kph1Doc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, kph1Doc)
}

// delete
// @Tags Doc Komputer PH1
// @Summary Delete komputer ph1 document by ID
// @ID delete-komputer-ph1-document-by-id
// @Security ApiKeyAuth
// @Router /api/doc/komputer-ph1/{id} [DELETE]
// @Produce json
// @Param id path string true "Komputer PH1 Document ID"
// @Success 200
func (h *KomputerPH1DocHandler) delete(c echo.Context) error {
	id := c.Param("id")

	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to convert id to ObjectID: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	err = h.kph1DocRepo.DeleteOneByID(oId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusNotFound, "Komputer PH1 Doc not found")
		}
		log.Errorf("Failed to delete kph1Doc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, "Komputer PH1 Doc deleted")
}
