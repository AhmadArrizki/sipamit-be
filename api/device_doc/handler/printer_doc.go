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

type PrinterDocHandler struct {
	printerRepo    *repo2.PrinterCollRepository
	printerDocRepo *repo.PrinterDocCollRepository
}

func NewPrinterDocAPIHandler(e *echo.Echo, db *mongo.Database) *PrinterDocHandler {
	h := &PrinterDocHandler{
		printerRepo:    repo2.NewPrinterRepository(db),
		printerDocRepo: repo.NewPrinterDocRepository(db),
	}

	group := e.Group("/api", context.Handler)

	group.GET("/doc/printers", h.findAll)
	group.GET("/doc/printer/:id", h.findByID)

	group.POST("/doc/printer", h.create)

	group.PUT("/doc/printer/:id", h.update)

	group.DELETE("/doc/printer/:id", h.delete)

	return h
}

// findAll
// @Tags Doc Printer
// @Summary Get all printer documents
// @ID get-all-printer-documents
// @Security ApiKeyAuth
// @Param q query string false "Search by nama"
// @Param page query int false "Page number pagination" default(1)
// @Param limit query int false "Limit pagination" default(10)
// @Param sort query string false "Sort" enums(asc,desc)
// @Router /api/doc/printers [GET]
// @Produce json
// @Success 200
func (h *PrinterDocHandler) findAll(c echo.Context) error {
	cq := util.NewCommonQuery(c)

	printerDocs, err := h.printerDocRepo.FindAll(cq)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get printerDocs: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "Printer Docs not found")
	}

	totalPrinterDocs, err := h.printerDocRepo.CountQuery(cq)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to count printerDocs: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "Printer Docs not found")
	}

	result := util.MakeResult(printerDocs, totalPrinterDocs, cq.Page, cq.Limit)
	return c.JSON(http.StatusOK, result)
}

// findByID
// @Tags Doc Printer
// @Summary Get printer document by ID
// @ID get-printer-document-by-id
// @Security ApiKeyAuth
// @Router /api/doc/printer/{id} [GET]
// @Produce json
// @Param id path string true "Printer Document ID"
// @Success 200
func (h *PrinterDocHandler) findByID(c echo.Context) error {
	id := c.Param("id")

	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to convert id to ObjectID: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	printerDoc, err := h.printerDocRepo.FindOneByID(oId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusNotFound, "Printer Doc not found")
		}
		log.Errorf("Failed to get printerDoc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, printerDoc)
}

// create
// @Tags Doc Printer
// @Summary Create printer document
// @ID create-printer-document
// @Security ApiKeyAuth
// @Router /api/doc/printer [POST]
// @Produce json
// @Param body body doc.DeviceDocForm true "Printer Document Form"
// @Success 200
func (h *PrinterDocHandler) create(c echo.Context) error {
	nc := c.(*context.Context)

	f, err := doc.NewDeviceDocForm(c)
	if err != nil {
		return err
	}

	printer, err := h.printerRepo.FindOneByID(f.DeviceOID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusNotFound, "Printer not found")
		}
		log.Errorf("Failed to get printer: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	printerDoc := &repo.PrinterDoc{
		ID:          bson.NewObjectID(),
		Nama:        printer.Nama,
		Departemen:  printer.Departemen,
		TipePrinter: printer.TipePrinter,
		NoSeri:      printer.NoSeri,
		Checkpoint:  f.Checkpoint,
		Inserted:    nc.Claims.ByAt(),
		IsDeleted:   false,
	}

	err = h.printerDocRepo.InsertOne(printerDoc)
	if err != nil {
		log.Errorf("Failed to create printerDoc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, printerDoc)
}

// update
// @Tags Doc Printer
// @Summary Update printer document by ID
// @ID update-printer-document-by-id
// @Security ApiKeyAuth
// @Router /api/doc/printer/{id} [PUT]
// @Produce json
// @Param id path string true "Printer Document ID"
// @Param body body doc.UpdateDeviceDocForm true "Printer Document Form"
// @Success 200
func (h *PrinterDocHandler) update(c echo.Context) error {
	nc := c.(*context.Context)

	id := c.Param("id")
	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to update printer: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid printer ID")
	}

	f, err := doc.NewUpdateDeviceDocForm(c)
	if err != nil {
		return err
	}

	printerDoc, err := h.printerDocRepo.FindOneByID(oId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusNotFound, "Printer Doc not found")
		}
		log.Errorf("Failed to update printerDoc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}

	printerDoc.Checkpoint = f.Checkpoint
	printerDoc.Updated = nc.Claims.ByAtPtr()
	err = h.printerDocRepo.UpdateOneByID(oId, printerDoc)
	if err != nil {
		log.Errorf("Failed to update printerDoc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, printerDoc)
}

// delete
// @Tags Doc Printer
// @Summary Delete printer document by ID
// @ID delete-printer-document-by-id
// @Security ApiKeyAuth
// @Router /api/doc/printer/{id} [DELETE]
// @Param id path string true "Printer Document ID"
// @Success 200
func (h *PrinterDocHandler) delete(c echo.Context) error {
	id := c.Param("id")

	oId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Errorf("Failed to convert id to ObjectID: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	err = h.printerDocRepo.DeleteOneByID(oId)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return echo.NewHTTPError(http.StatusNotFound, "Printer Doc not found")
		}
		log.Errorf("Failed to delete printerDoc: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, "Printer Doc deleted")
}
