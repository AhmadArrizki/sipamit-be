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

type komputerPH2Form struct {
	Nama     string `form:"nama" json:"nama"`
	Merk     string `form:"merk" json:"merk"`
	PC       string `form:"pc" json:"pc"`
	Monitor  string `form:"monitor" json:"monitor"`
	CPU      string `form:"cpu" json:"cpu"`
	RAM      string `form:"ram" json:"ram"`
	Internal string `form:"internal" json:"internal"`
	Lokasi   string `form:"lokasi" json:"lokasi"`
}

func newKomputerPH2Form(c echo.Context) (*komputerPH2Form, error) {
	f := new(komputerPH2Form)
	if err := c.Bind(f); err != nil {
		log.Errorf("Failed to bind komputer ph2 form: %v", err)
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	if f.Nama == "" && f.Merk == "" && f.PC == "" && f.Monitor == "" && f.CPU == "" && f.RAM == "" && f.Internal == "" && f.Lokasi == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Please fill provided field")
	}

	return f, nil
}

type KomputerPH2Handler struct {
	kph2Repo *repo.KomputerPH2CollRepository
}

func NewKomputerPH2APIHandler(e *echo.Echo, db *mongo.Database) *KomputerPH2Handler {
	h := &KomputerPH2Handler{
		kph2Repo: repo.NewKomputerPH2Repository(db),
	}

	group := e.Group("/api", context.Handler)

	group.GET("/komputer-ph2s", h.findAll)
	group.GET("/komputer-ph2/:id", h.findOne)

	group.POST("/komputer-ph2", h.create)

	group.PUT("/komputer-ph2/:id", h.update)

	group.DELETE("/komputer-ph2/:id", h.delete)

	return h
}

// findAll
// @Tags Device KomputerPH2
// @Summary Get all komputer-ph2s
// @ID get-all-komputer-ph2s
// @Security ApiKeyAuth
// @Router /api/komputer-ph2s [GET]
// @Produce json
// @Success 200
func (h *KomputerPH2Handler) findAll(c echo.Context) error {
	komputerPH2s, err := h.kph2Repo.FindAll()
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get komputerPH2s: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "KomputerPH2s not found")
	}
	return c.JSON(http.StatusOK, komputerPH2s)
}

// findOne
// @Tags Device KomputerPH2
// @Summary Get komputer-ph2 by id
// @ID get-komputer-ph2-by-id
// @Security ApiKeyAuth
// @Router /api/komputer-ph2/{id} [GET]
// @Produce json
// @Param id path string true "KomputerPH2 ID"
// @Success 200
func (h *KomputerPH2Handler) findOne(c echo.Context) error {
	id := c.Param("id")

	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to get komputerPH2: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid komputerPH2 ID")
	}

	komputerPH2, err := h.kph2Repo.FindOneByID(oId)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get komputerPH2: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "KomputerPH2 not found")
	}
	return c.JSON(http.StatusOK, komputerPH2)
}

// create
// @Tags Device KomputerPH2
// @Summary Create new komputer-ph2
// @ID create-komputer-ph2
// @Security ApiKeyAuth
// @Router /api/komputer-ph2 [POST]
// @Produce json
// @Param body body komputerPH2Form true "KomputerPH2 Form"
// @Success 200
func (h *KomputerPH2Handler) create(c echo.Context) error {
	nc := c.(*context.Context)

	f, err := newKomputerPH2Form(nc)
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

	kph2 := &repo.KomputerPH2{
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

	err = h.kph2Repo.InsertOne(kph2)
	if err != nil {
		log.Errorf("Failed to create komputerPH2: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, kph2)
}

// update
// @Tags Device KomputerPH2
// @Summary Update komputer-ph2 by id
// @ID update-komputer-ph2-by-id
// @Security ApiKeyAuth
// @Router /api/komputer-ph2/{id} [PUT]
// @Produce json
// @Param id path string true "KomputerPH2 ID"
// @Param body body komputerPH2Form true "KomputerPH2 Form"
// @Success 200
func (h *KomputerPH2Handler) update(c echo.Context) error {
	nc := c.(*context.Context)

	id := c.Param("id")
	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to update komputerPH2: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid komputer ph2 ID")
	}

	f, err := newKomputerPH2Form(nc)
	if err != nil {
		return err
	}

	kph2, err := h.kph2Repo.FindOneByID(oId)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to update komputerPH2: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "KomputerPH2 not found")
	}

	if f.Nama != "" {
		kph2.Nama = f.Nama
	}
	if f.Merk != "" {
		kph2.Merk = f.Merk
	}
	if f.PC != "" {
		kph2.PC = f.PC
	}
	if f.Monitor != "" {
		kph2.Monitor = f.Monitor
	}
	if f.CPU != "" {
		kph2.CPU = f.CPU
	}
	if f.RAM != "" {
		kph2.RAM = f.RAM
	}
	if f.Internal != "" {
		kph2.Internal = f.Internal
	}
	if f.Lokasi != "" {
		kph2.Lokasi = f.Lokasi
	}

	kph2.Updated = nc.Claims.ByAtPtr()
	err = h.kph2Repo.UpdateOneByID(oId, kph2)
	if err != nil {
		log.Errorf("Failed to update komputerPH2: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, kph2)
}

// delete
// @Tags Device KomputerPH2
// @Summary Delete komputer-ph2 by id
// @ID delete-komputer-ph2-by-id
// @Security ApiKeyAuth
// @Router /api/komputer-ph2/{id} [DELETE]
// @Produce json
// @Param id path string true "KomputerPH2 ID"
// @Success 200
func (h *KomputerPH2Handler) delete(c echo.Context) error {
	id := c.Param("id")

	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to delete komputerPH2: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid komputerPH2 ID")
	}

	kph2, _ := h.kph2Repo.FindOneByID(oId)
	if kph2 == nil {
		return echo.NewHTTPError(http.StatusNotFound, "KomputerPH2 not found")
	}

	err = h.kph2Repo.DeleteOneByID(oId)
	if err != nil {
		log.Errorf("Failed to delete komputerPH2: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, "KomputerPH2 deleted")
}
