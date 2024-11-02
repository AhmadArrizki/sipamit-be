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

type KomputerPH2DocHandler struct {
	kph2Repo    *repo2.KomputerPH2CollRepository
	kph2DocRepo *repo.KomputerPH2DocCollRepository
}

func NewKomputerPH2DocAPIHandler(e *echo.Echo, db *mongo.Database) *KomputerPH2DocHandler {
	h := &KomputerPH2DocHandler{
		kph2Repo:    repo2.NewKomputerPH2Repository(db),
		kph2DocRepo: repo.NewKomputerPH2DocRepository(db),
	}

	group := e.Group("/api", context.Handler)

	group.GET("/doc/komputer-ph2s", h.findAll)
	group.GET("/doc/komputer-ph2/:id", h.findByID)

	group.POST("/doc/komputer-ph2", h.create)

	group.PUT("/doc/komputer-ph2/:id", h.update)

	group.DELETE("/doc/komputer-ph2/:id", h.delete)

	return h
}

// findAll
// @Tags Doc Komputer PH2
// @Summary Get all komputer ph2 documents
// @ID get-all-komputer-ph2-documents
// @Security ApiKeyAuth
// @Param q query string false "Search by nama"
// @Param page query int false "Page number pagination" default(1)
// @Param limit query int false "Limit pagination" default(10)
// @Param sort query string false "Sort" enums(asc,desc)
// @Router /api/doc/komputer-ph2s [GET]
// @Produce json
// @Success 200
func (h *KomputerPH2DocHandler) findAll(c echo.Context) error {
	cq := util.NewCommonQuery(c)

	kph2Doc, err := h.kph2DocRepo.FindAll(cq)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get kph2Doc: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "Komputer PH2 Docs not found")
	}

	totalKph2Docs, err := h.kph2DocRepo.CountQuery(cq)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to count kph2Doc: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "Komputer PH2 Docs not found")
	}

	result := util.MakeResult(kph2Doc, totalKph2Docs, cq.Page, cq.Limit)
	return c.JSON(http.StatusOK, result)
}

// findByID
// @Tags Doc Komputer PH2
// @Summary Get komputer ph2 document by ID
// @ID get-komputer-ph2-document-by-id
// @Security ApiKeyAuth
// @Router /api/doc/komputer-ph2/{id} [GET]
// @Produce json
// @Param id path string true "Komputer PH2 Document ID"
// @Success 200
func (h *KomputerPH2DocHandler) findByID(c echo.Context) error {
	id := c.Param("id")

	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to convert id to ObjectID: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	kph2Doc, err := h.kph2DocRepo.FindOneByID(oId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusNotFound, "Komputer PH2 Doc not found")
		}
		log.Errorf("Failed to get kph2Doc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, kph2Doc)
}

// create
// @Tags Doc Komputer PH2
// @Summary Create new komputer ph2 document
// @ID create-new-komputer-ph2-document
// @Security ApiKeyAuth
// @Router /api/doc/komputer-ph2 [POST]
// @Produce json
// @Param body body doc.DeviceDocForm true "Komputer PH2 Document Form"
// @Success 200
func (h *KomputerPH2DocHandler) create(c echo.Context) error {
	nc := c.(*context.Context)

	f, err := doc.NewDeviceDocForm(c)
	if err != nil {
		return err
	}

	kph2, err := h.kph2Repo.FindOneByID(f.DeviceOID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusNotFound, "Komputer PH2 not found")
		}
		log.Errorf("Failed to get kph2: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	kph2Doc := &repo.KomputerPH2Doc{
		ID:         bson.NewObjectID(),
		Nama:       kph2.Nama,
		Merk:       kph2.Merk,
		PC:         kph2.PC,
		Monitor:    kph2.Monitor,
		CPU:        kph2.CPU,
		RAM:        kph2.RAM,
		Internal:   kph2.Internal,
		Lokasi:     kph2.Lokasi,
		Checkpoint: f.Checkpoint,
		Inserted:   nc.Claims.ByAt(),
		IsDeleted:  false,
	}

	err = h.kph2DocRepo.InsertOne(kph2Doc)
	if err != nil {
		log.Errorf("Failed to create kph2Doc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, kph2Doc)
}

// update
// @Tags Doc Komputer PH2
// @Summary Update komputer ph2 document by ID
// @ID update-komputer-ph2-document-by-id
// @Security ApiKeyAuth
// @Router /api/doc/komputer-ph2/{id} [PUT]
// @Produce json
// @Param id path string true "Komputer PH2 Document ID"
// @Param body body doc.UpdateDeviceDocForm true "Komputer PH2 Document Form"
// @Success 200
func (h *KomputerPH2DocHandler) update(c echo.Context) error {
	nc := c.(*context.Context)

	id := c.Param("id")
	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to update komputer ph2: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid komputer ph2 ID")
	}

	f, err := doc.NewUpdateDeviceDocForm(c)
	if err != nil {
		return err
	}

	kph2Doc, err := h.kph2DocRepo.FindOneByID(oId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusNotFound, "Komputer PH2 Doc not found")
		}
		log.Errorf("Failed to update kph2Doc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	kph2Doc.Checkpoint = f.Checkpoint
	kph2Doc.Updated = nc.Claims.ByAtPtr()
	err = h.kph2DocRepo.UpdateOneByID(oId, kph2Doc)
	if err != nil {
		log.Errorf("Failed to update kph2Doc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, kph2Doc)
}

// delete
// @Tags Doc Komputer PH2
// @Summary Delete komputer ph2 document by ID
// @ID delete-komputer-ph2-document-by-id
// @Security ApiKeyAuth
// @Router /api/doc/komputer-ph2/{id} [DELETE]
// @Produce json
// @Param id path string true "Komputer PH2 Document ID"
// @Success 200
func (h *KomputerPH2DocHandler) delete(c echo.Context) error {
	id := c.Param("id")

	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to convert id to ObjectID: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	err = h.kph2DocRepo.DeleteOneByID(oId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusNotFound, "Komputer PH2 Doc not found")
		}
		log.Errorf("Failed to delete kph2Doc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, "Komputer PH2 Doc deleted")
}
