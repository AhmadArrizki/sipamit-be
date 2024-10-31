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

type printerForm struct {
	Nama        string `form:"nama" json:"nama"`
	Departemen  string `form:"departemen" json:"departemen"`
	TipePrinter string `form:"tipe_printer" json:"tipe_printer"`
	NoSeri      string `form:"no_seri" json:"no_seri"`
}

func newPrinterForm(c echo.Context) (*printerForm, error) {
	f := new(printerForm)
	if err := c.Bind(f); err != nil {
		log.Errorf("Failed to bind printer form: %v", err)
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	if f.Nama == "" && f.Departemen == "" && f.TipePrinter == "" && f.NoSeri == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Please fill provided field")
	}

	return f, nil
}

type PrinterHandler struct {
	printerRepo *repo.PrinterCollRepository
}

func NewPrinterAPIHandler(e *echo.Echo, db *mongo.Database) *PrinterHandler {
	h := &PrinterHandler{
		printerRepo: repo.NewPrinterRepository(db),
	}

	group := e.Group("/api", context.Handler)

	group.GET("/printers", h.findAll)
	group.GET("/printer/:id", h.findOne)

	group.POST("/printer", h.create)

	group.PUT("/printer/:id", h.update)

	group.DELETE("/printer/:id", h.delete)

	return h
}

// findAll
// @Tags Device Printer
// @Summary Get all printers
// @ID get-all-printers
// @Security ApiKeyAuth
// @Router /api/printers [GET]
// @Produce json
// @Success 200
func (h *PrinterHandler) findAll(c echo.Context) error {
	printers, err := h.printerRepo.FindAll()
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get printers: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "Printers not found")
	}
	return c.JSON(http.StatusOK, printers)
}

// findOne
// @Tags Device Printer
// @Summary Get printer by id
// @ID get-printer-by-id
// @Security ApiKeyAuth
// @Router /api/printer/{id} [GET]
// @Produce json
// @Param id path string true "Printer ID"
// @Success 200
func (h *PrinterHandler) findOne(c echo.Context) error {
	id := c.Param("id")

	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to get printer: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid printer ID")
	}

	printer, err := h.printerRepo.FindOneByID(oId)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get printer: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "Printer not found")
	}
	return c.JSON(http.StatusOK, printer)
}

// create
// @Tags Device Printer
// @Summary Create new printer
// @ID create-new-printer
// @Security ApiKeyAuth
// @Router /api/printer [POST]
// @Produce json
// @Param body body printerForm true "Printer Form"
// @Success 200
func (h *PrinterHandler) create(c echo.Context) error {
	nc := c.(*context.Context)

	f, err := newPrinterForm(nc)
	if err != nil {
		return err
	}

	if f.Nama == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Nama is required")
	}
	if f.Departemen == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Departemen is required")
	}
	if f.TipePrinter == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Tipe Printer is required")
	}
	if f.NoSeri == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "No Seri is required")
	}

	printer := &repo.Printer{
		ID:          bson.NewObjectID(),
		Nama:        f.Nama,
		Departemen:  f.Departemen,
		TipePrinter: f.TipePrinter,
		NoSeri:      f.NoSeri,
		Inserted:    nc.Claims.ByAt(),
		IsDeleted:   false,
	}

	err = h.printerRepo.InsertOne(printer)
	if err != nil {
		log.Errorf("Failed to create printer: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, printer)
}

// update
// @Tags Device Printer
// @Summary Update printer by id
// @ID update-printer-by-id
// @Security ApiKeyAuth
// @Router /api/printer/{id} [PUT]
// @Produce json
// @Param id path string true "Printer ID"
// @Param body body printerForm true "Printer Form"
// @Success 200
func (h *PrinterHandler) update(c echo.Context) error {
	nc := c.(*context.Context)

	id := c.Param("id")
	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to update printer: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid printer ID")
	}

	f, err := newPrinterForm(c)
	if err != nil {
		return err
	}

	printer, err := h.printerRepo.FindOneByID(oId)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to update printer: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "Printer not found")
	}

	if f.Nama != "" {
		printer.Nama = f.Nama
	}
	if f.Departemen != "" {
		printer.Departemen = f.Departemen
	}
	if f.TipePrinter != "" {
		printer.TipePrinter = f.TipePrinter
	}
	if f.NoSeri != "" {
		printer.NoSeri = f.NoSeri
	}

	printer.Updated = nc.Claims.ByAtPtr()
	err = h.printerRepo.UpdateOneByID(oId, printer)
	if err != nil {
		log.Errorf("Failed to update printer: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, printer)
}

// delete
// @Tags Device Printer
// @Summary Delete printer by id
// @ID delete-printer-by-id
// @Security ApiKeyAuth
// @Router /api/printer/{id} [DELETE]
// @Produce json
// @Param id path string true "Printer ID"
// @Success 200
func (h *PrinterHandler) delete(c echo.Context) error {
	id := c.Param("id")

	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to delete printer: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid printer ID")
	}

	printer, _ := h.printerRepo.FindOneByID(oId)
	if printer == nil {
		return echo.NewHTTPError(http.StatusNotFound, "Printer not found")
	}

	err = h.printerRepo.DeleteOneByID(oId)
	if err != nil {
		log.Errorf("Failed to delete printer: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, "Printer deleted")
}
