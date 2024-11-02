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

type cctvForm struct {
	Nama   string `form:"nama" json:"nama"`
	Lokasi string `form:"lokasi" json:"lokasi" `
	Kode   string `form:"kode" json:"kode"`
}

func newCCTVForm(c echo.Context) (*cctvForm, error) {
	f := new(cctvForm)
	if err := c.Bind(f); err != nil {
		log.Errorf("Failed to bind cctv form: %v", err)
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	if f.Nama == "" && f.Lokasi == "" && f.Kode == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Please fill provided field")
	}

	return f, nil
}

type CCTVHandler struct {
	cctvRepo *repo.CCTVCollRepository
}

func NewCCTVAPIHandler(e *echo.Echo, db *mongo.Database) *CCTVHandler {
	h := &CCTVHandler{
		cctvRepo: repo.NewCCTVRepository(db),
	}

	group := e.Group("/api", context.Handler)

	group.GET("/cctvs", h.findAll)
	group.GET("/cctv/:id", h.findOne)

	group.POST("/cctv", h.create)

	group.PUT("/cctv/:id", h.update)

	group.DELETE("/cctv/:id", h.delete)

	return h
}

// findAll
// @Tags Device CCTV
// @Summary Get all cctvs
// @ID get-all-cctvs
// @Security ApiKeyAuth
// @Param q query string false "Search by nama"
// @Param page query int false "Page number pagination" default(1)
// @Param limit query int false "Limit pagination" default(10)
// @Param sort query string false "Sort" enums(asc,desc)
// @Router /api/cctvs [GET]
// @Produce json
// @Success 200
func (h *CCTVHandler) findAll(c echo.Context) error {
	cq := util.NewCommonQuery(c)

	cctvs, err := h.cctvRepo.FindAll(cq)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get cctvs: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "CCTVs not found")
	}

	totalCctv, err := h.cctvRepo.CountQuery(cq)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to count cctvs: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "CCTVs not found")
	}

	result := util.MakeResult(cctvs, totalCctv, cq.Page, cq.Limit)
	return c.JSON(http.StatusOK, result)
}

// findOne
// @Tags Device CCTV
// @Summary Get cctv by id
// @ID get-cctv-by-id
// @Security ApiKeyAuth
// @Router /api/cctv/{id} [GET]
// @Produce json
// @Param id path string true "CCTV ID"
// @Success 200
func (h *CCTVHandler) findOne(c echo.Context) error {
	id := c.Param("id")

	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to get cctv: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid cctv ID")
	}

	cctv, err := h.cctvRepo.FindOneByID(oId)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get cctv: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "CCTV not found")
	}
	return c.JSON(http.StatusOK, cctv)
}

// create
// @Tags Device CCTV
// @Summary Create cctv
// @ID create-cctv
// @Security ApiKeyAuth
// @Router /api/cctv [POST]
// @Produce json
// @Param body body cctvForm true "CCTV Form"
// @Success 200
func (h *CCTVHandler) create(c echo.Context) error {
	nc := c.(*context.Context)

	f, err := newCCTVForm(c)
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

	cctv := &repo.CCTV{
		ID:        bson.NewObjectID(),
		Nama:      f.Nama,
		Lokasi:    f.Lokasi,
		Kode:      f.Kode,
		Inserted:  nc.Claims.ByAt(),
		IsDeleted: false,
	}

	err = h.cctvRepo.InsertOne(cctv)
	if err != nil {
		log.Errorf("Failed to create cctv: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, cctv)
}

// update
// @Tags Device CCTV
// @Summary Update cctv by id
// @ID update-cctv-by-id
// @Security ApiKeyAuth
// @Router /api/cctv/{id} [PUT]
// @Produce json
// @Param id path string true "CCTV ID"
// @Param body body cctvForm true "CCTV Form"
// @Success 200
func (h *CCTVHandler) update(c echo.Context) error {
	nc := c.(*context.Context)

	id := c.Param("id")
	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to update cctv: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid cctv ID")
	}

	f, err := newCCTVForm(c)
	if err != nil {
		return err
	}

	cctv, err := h.cctvRepo.FindOneByID(oId)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to update cctv: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "CCTV not found")
	}

	if f.Nama != "" {
		cctv.Nama = f.Nama
	}
	if f.Lokasi != "" {
		cctv.Lokasi = f.Lokasi
	}
	if f.Kode != "" {
		cctv.Kode = f.Kode
	}

	cctv.Updated = nc.Claims.ByAtPtr()
	err = h.cctvRepo.UpdateOneByID(oId, cctv)
	if err != nil {
		log.Errorf("Failed to update cctv: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, cctv)
}

// delete
// @Tags Device CCTV
// @Summary Delete cctv by id
// @ID delete-cctv-by-id
// @Security ApiKeyAuth
// @Router /api/cctv/{id} [DELETE]
// @Produce json
// @Param id path string true "CCTV ID"
// @Success 200
func (h *CCTVHandler) delete(c echo.Context) error {
	id := c.Param("id")

	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to delete cctv: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid cctv ID")
	}

	cctv, _ := h.cctvRepo.FindOneByID(oId)
	if cctv == nil {
		return echo.NewHTTPError(http.StatusNotFound, "CCTV not found")
	}

	err = h.cctvRepo.DeleteOneByID(oId)
	if err != nil {
		log.Errorf("Failed to delete cctv: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, "CCTV deleted")
}
