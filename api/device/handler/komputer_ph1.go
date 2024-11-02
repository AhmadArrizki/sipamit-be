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

type komputerPH1Form struct {
	Nama     string `form:"nama" json:"nama"`
	Merk     string `form:"merk" json:"merk"`
	PC       string `form:"pc" json:"pc"`
	Monitor  string `form:"monitor" json:"monitor"`
	CPU      string `form:"cpu" json:"cpu"`
	RAM      string `form:"ram" json:"ram"`
	Internal string `form:"internal" json:"internal"`
	Lokasi   string `form:"lokasi" json:"lokasi"`
}

func newKomputerPH1Form(c echo.Context) (*komputerPH1Form, error) {
	f := new(komputerPH1Form)
	if err := c.Bind(f); err != nil {
		log.Errorf("Failed to bind komputer ph1 form: %v", err)
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	if f.Nama == "" && f.Merk == "" && f.PC == "" && f.Monitor == "" && f.CPU == "" && f.RAM == "" && f.Internal == "" && f.Lokasi == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Please fill provided field")
	}

	return f, nil
}

type KomputerPH1Handler struct {
	kph1Repo *repo.KomputerPH1CollRepository
}

func NewKomputerPH1APIHandler(e *echo.Echo, db *mongo.Database) *KomputerPH1Handler {
	h := &KomputerPH1Handler{
		kph1Repo: repo.NewKomputerPH1Repository(db),
	}

	group := e.Group("/api", context.Handler)

	group.GET("/komputer-ph1s", h.findAll)
	group.GET("/komputer-ph1/:id", h.findOne)

	group.POST("/komputer-ph1", h.create)

	group.PUT("/komputer-ph1/:id", h.update)

	group.DELETE("/komputer-ph1/:id", h.delete)

	return h
}

// findAll
// @Tags Device KomputerPH1
// @Summary Get all komputer-ph1s
// @ID get-all-komputer-ph1s
// @Security ApiKeyAuth
// @Param q query string false "Search by nama"
// @Param page query int false "Page number pagination" default(1)
// @Param limit query int false "Limit pagination" default(10)
// @Param sort query string false "Sort" enums(asc,desc)
// @Router /api/komputer-ph1s [GET]
// @Produce json
// @Success 200
func (h *KomputerPH1Handler) findAll(c echo.Context) error {
	cq := util.NewCommonQuery(c)

	komputerPH1s, err := h.kph1Repo.FindAll(cq)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get komputerPH1s: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "KomputerPH1s not found")
	}

	totalKph1, err := h.kph1Repo.CountQuery(cq)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to count komputerPH1s: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "KomputerPH1s not found")
	}

	result := util.MakeResult(komputerPH1s, totalKph1, cq.Page, cq.Limit)
	return c.JSON(http.StatusOK, result)
}

// findOne
// @Tags Device KomputerPH1
// @Summary Get komputer-ph1 by id
// @ID get-komputer-ph1-by-id
// @Security ApiKeyAuth
// @Router /api/komputer-ph1/{id} [GET]
// @Produce json
// @Param id path string true "KomputerPH1 ID"
// @Success 200
func (h *KomputerPH1Handler) findOne(c echo.Context) error {
	id := c.Param("id")

	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to get komputer ph1: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid komputer ph1 ID")
	}

	komputerPH1, err := h.kph1Repo.FindOneByID(oId)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get komputer ph1: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "KomputerPH1 not found")
	}
	return c.JSON(http.StatusOK, komputerPH1)
}

// create
// @Tags Device KomputerPH1
// @Summary Create new komputer-ph1
// @ID create-komputer-ph1
// @Security ApiKeyAuth
// @Router /api/komputer-ph1 [POST]
// @Produce json
// @Param body body komputerPH1Form true "KomputerPH1 Form"
// @Success 200
func (h *KomputerPH1Handler) create(c echo.Context) error {
	nc := c.(*context.Context)

	f, err := newKomputerPH1Form(nc)
	if err != nil {
		return err
	}

	if f.Nama == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Nama is required")
	}
	if f.Merk == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Merk is required")
	}
	if f.PC == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "PC is required")
	}
	if f.Monitor == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Monitor is required")
	}
	if f.CPU == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "CPU is required")
	}
	if f.RAM == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "RAM is required")
	}
	if f.Internal == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Internal is required")
	}
	if f.Lokasi == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Lokasi is required")
	}

	kph1 := &repo.KomputerPH1{
		ID:        bson.NewObjectID(),
		Nama:      f.Nama,
		Merk:      f.Merk,
		PC:        f.PC,
		Monitor:   f.Monitor,
		CPU:       f.CPU,
		RAM:       f.RAM,
		Internal:  f.Internal,
		Lokasi:    f.Lokasi,
		Inserted:  nc.Claims.ByAt(),
		IsDeleted: false,
	}

	err = h.kph1Repo.InsertOne(kph1)
	if err != nil {
		log.Errorf("Failed to create komputerPH1: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, kph1)
}

// update
// @Tags Device KomputerPH1
// @Summary Update komputer-ph1 by id
// @ID update-komputer-ph1-by-id
// @Security ApiKeyAuth
// @Router /api/komputer-ph1/{id} [PUT]
// @Produce json
// @Param id path string true "KomputerPH1 ID"
// @Param body body komputerPH1Form true "KomputerPH1 Form"
// @Success 200
func (h *KomputerPH1Handler) update(c echo.Context) error {
	nc := c.(*context.Context)

	id := c.Param("id")
	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to update komputerPH1: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid komputer ph1 ID")
	}

	f, err := newKomputerPH1Form(nc)
	if err != nil {
		return err
	}

	kph1, err := h.kph1Repo.FindOneByID(oId)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to update komputerPH1: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "KomputerPH1 not found")
	}

	if f.Nama != "" {
		kph1.Nama = f.Nama
	}
	if f.Merk != "" {
		kph1.Merk = f.Merk
	}
	if f.PC != "" {
		kph1.PC = f.PC
	}
	if f.Monitor != "" {
		kph1.Monitor = f.Monitor
	}
	if f.CPU != "" {
		kph1.CPU = f.CPU
	}
	if f.RAM != "" {
		kph1.RAM = f.RAM
	}
	if f.Internal != "" {
		kph1.Internal = f.Internal
	}
	if f.Lokasi != "" {
		kph1.Lokasi = f.Lokasi
	}

	kph1.Updated = nc.Claims.ByAtPtr()
	err = h.kph1Repo.UpdateOneByID(oId, kph1)
	if err != nil {
		log.Errorf("Failed to update komputerPH1: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, kph1)
}

// delete
// @Tags Device KomputerPH1
// @Summary Delete komputer-ph1 by id
// @ID delete-komputer-ph1-by-id
// @Security ApiKeyAuth
// @Router /api/komputer-ph1/{id} [DELETE]
// @Produce json
// @Param id path string true "KomputerPH1 ID"
// @Success 200
func (h *KomputerPH1Handler) delete(c echo.Context) error {
	id := c.Param("id")

	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to delete komputerPH1: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid komputer ph1 ID")
	}

	kph1, _ := h.kph1Repo.FindOneByID(oId)
	if kph1 == nil {
		return echo.NewHTTPError(http.StatusNotFound, "KomputerPH1 not found")
	}

	err = h.kph1Repo.DeleteOneByID(oId)
	if err != nil {
		log.Errorf("Failed to delete komputerPH1: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, "KomputerPH1 deleted")
}
