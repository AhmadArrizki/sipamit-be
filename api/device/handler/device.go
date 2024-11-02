package handler

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
	"sipamit-be/api/device/repo"
	"sipamit-be/internal/pkg/context"
	"sipamit-be/internal/pkg/doc"
)

type DeviceHandler struct {
	cctvRepo    *repo.CCTVCollRepository
	fpRepo      *repo.FingerPrintCollRepository
	kph1Repo    *repo.KomputerPH1CollRepository
	kph2Repo    *repo.KomputerPH2CollRepository
	printerRepo *repo.PrinterCollRepository
	teleponRepo *repo.TeleponCollRepository
	toaRepo     *repo.TOACollRepository
	upsRepo     *repo.UPSCollRepository
}

func NewDeviceAPIHandler(e *echo.Echo, db *mongo.Database) *DeviceHandler {
	h := &DeviceHandler{
		cctvRepo:    repo.NewCCTVRepository(db),
		fpRepo:      repo.NewFingerPrintRepository(db),
		kph1Repo:    repo.NewKomputerPH1Repository(db),
		kph2Repo:    repo.NewKomputerPH2Repository(db),
		printerRepo: repo.NewPrinterRepository(db),
		teleponRepo: repo.NewTeleponRepository(db),
		toaRepo:     repo.NewTOARepository(db),
		upsRepo:     repo.NewUPSRepository(db),
	}

	group := e.Group("/api", context.Handler)

	group.GET("/device/count", h.count)

	return h
}

// count
// @Tags Device
// @Summary Count all devices
// @ID device-count
// @Security ApiKeyAuth
// @Param device query string false "Device type" enums(cctv, fingerprint, komputer_ph1, komputer_ph2, printer, telepon, toa, ups)
// @Router /api/device/count [GET]
// @Produce json
// @Success 200
func (h *DeviceHandler) count(c echo.Context) error {
	var total int64 = 0
	param := c.QueryParam("device")

	switch param {
	case doc.CCTV:
		cctvs, _ := h.cctvRepo.Count()
		total += cctvs
	case doc.Fingerprint:
		fps, _ := h.fpRepo.Count()
		total += fps
	case doc.KomputerPH1:
		kph1s, _ := h.kph1Repo.Count()
		total += kph1s
	case doc.KomputerPH2:
		kph2s, _ := h.kph2Repo.Count()
		total += kph2s
	case doc.Printer:
		printers, _ := h.printerRepo.Count()
		total += printers
	case doc.Telepon:
		telepons, _ := h.teleponRepo.Count()
		total += telepons
	case doc.Toa:
		toas, _ := h.toaRepo.Count()
		total += toas
	case doc.Ups:
		upss, _ := h.upsRepo.Count()
		total += upss
	default:
		cctvs, _ := h.cctvRepo.Count()
		fps, _ := h.fpRepo.Count()
		kph1s, _ := h.kph1Repo.Count()
		kph2s, _ := h.kph2Repo.Count()
		printers, _ := h.printerRepo.Count()
		telepons, _ := h.teleponRepo.Count()
		toas, _ := h.toaRepo.Count()
		upss, _ := h.upsRepo.Count()

		total = cctvs + fps + kph1s + kph2s + printers + telepons + toas + upss
	}
	return c.JSON(http.StatusOK, map[string]int64{"total": total})
}
