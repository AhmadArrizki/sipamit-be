package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"net/http"
	"sipamit-be/api/device_cp/repo"
	"sipamit-be/internal/pkg/const"
	"sipamit-be/internal/pkg/context"
	"sipamit-be/internal/pkg/log"
)

type updateCheckpointForm struct {
	Checkpoint []string `json:"checkpoint" form:"checkpoint"`
}

func newUpdateCheckpointForm(c echo.Context) (*updateCheckpointForm, error) {
	f := new(updateCheckpointForm)
	if err := c.Bind(f); err != nil {
		log.Errorf("Failed to bind update checkpoint form: %v", err)
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	if len(f.Checkpoint) <= 0 {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Checkpoint is required")
	}
	return f, nil
}

type CheckpointHandler struct {
	cpRepo *repo.CheckpointCollRepository
}

func NewCheckpointAPIHandler(e *echo.Echo, db *mongo.Database) *CheckpointHandler {
	h := &CheckpointHandler{
		cpRepo: repo.NewCheckpointRepository(db),
	}

	group := e.Group("/api", context.Handler)

	group.GET("/checkpoint/cctv", h.cctv)
	group.GET("/checkpoint/fingerprint", h.fingerprint)
	group.GET("/checkpoint/komputer-ph1", h.komputerPh1)
	group.GET("/checkpoint/komputer-ph2", h.komputerPh2)
	group.GET("/checkpoint/printer", h.printer)
	group.GET("/checkpoint/telepon", h.telepon)
	group.GET("/checkpoint/toa", h.toa)
	group.GET("/checkpoint/ups", h.ups)

	group.PUT("/checkpoint/cctv", h.updateCCTV)
	group.PUT("/checkpoint/fingerprint", h.updateFingerprint)
	group.PUT("/checkpoint/komputer-ph1", h.updateKomputerPh1)
	group.PUT("/checkpoint/komputer-ph2", h.updateKomputerPh2)
	group.PUT("/checkpoint/printer", h.updatePrinter)
	group.PUT("/checkpoint/telepon", h.updateTelepon)
	group.PUT("/checkpoint/toa", h.updateToa)
	group.PUT("/checkpoint/ups", h.updateUps)

	return h
}

// cctv
// @Tags Checkpoint
// @Summary Get cctv checkpoint
// @ID get-cctv-checkpoint
// @Security ApiKeyAuth
// @Router /api/checkpoint/cctv [GET]
// @Produce json
// @Success 200
func (h *CheckpointHandler) cctv(c echo.Context) error {
	cctv, err := h.cpRepo.FindByDevice(_const.CCTV)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get cctv checkpoint: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "CCTV checkpoint not found")
	}
	return c.JSON(http.StatusOK, cctv)
}

// fingerprint
// @Tags Checkpoint
// @Summary Get fingerprint checkpoint
// @ID get-fingerprint-checkpoint
// @Security ApiKeyAuth
// @Router /api/checkpoint/fingerprint [GET]
// @Produce json
// @Success 200
func (h *CheckpointHandler) fingerprint(c echo.Context) error {
	fingerprint, err := h.cpRepo.FindByDevice(_const.Fingerprint)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get fingerprint checkpoint: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "Fingerprint checkpoint not found")
	}
	return c.JSON(http.StatusOK, fingerprint)
}

// komputerPh1
// @Tags Checkpoint
// @Summary Get komputer-ph1 checkpoint
// @ID get-komputer-ph1-checkpoint
// @Security ApiKeyAuth
// @Router /api/checkpoint/komputer-ph1 [GET]
// @Produce json
// @Success 200
func (h *CheckpointHandler) komputerPh1(c echo.Context) error {
	komputerPh1, err := h.cpRepo.FindByDevice(_const.KomputerPH1)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get komputer-ph1 checkpoint: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "Komputer-ph1 checkpoint not found")
	}
	return c.JSON(http.StatusOK, komputerPh1)
}

// komputerPh2
// @Tags Checkpoint
// @Summary Get komputer-ph2 checkpoint
// @ID get-komputer-ph2-checkpoint
// @Security ApiKeyAuth
// @Router /api/checkpoint/komputer-ph2 [GET]
// @Produce json
// @Success 200
func (h *CheckpointHandler) komputerPh2(c echo.Context) error {
	komputerPh2, err := h.cpRepo.FindByDevice(_const.KomputerPH2)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get komputer-ph2 checkpoint: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "Komputer-ph2 checkpoint not found")
	}
	return c.JSON(http.StatusOK, komputerPh2)
}

// printer
// @Tags Checkpoint
// @Summary Get printer checkpoint
// @ID get-printer-checkpoint
// @Security ApiKeyAuth
// @Router /api/checkpoint/printer [GET]
// @Produce json
// @Success 200
func (h *CheckpointHandler) printer(c echo.Context) error {
	printer, err := h.cpRepo.FindByDevice(_const.Printer)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get printer checkpoint: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "Printer checkpoint not found")
	}
	return c.JSON(http.StatusOK, printer)
}

// telepon
// @Tags Checkpoint
// @Summary Get telepon checkpoint
// @ID get-telepon-checkpoint
// @Security ApiKeyAuth
// @Router /api/checkpoint/telepon [GET]
// @Produce json
// @Success 200
func (h *CheckpointHandler) telepon(c echo.Context) error {
	telepon, err := h.cpRepo.FindByDevice(_const.Telepon)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get telepon checkpoint: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "Telepon checkpoint not found")
	}
	return c.JSON(http.StatusOK, telepon)
}

// toa
// @Tags Checkpoint
// @Summary Get toa checkpoint
// @ID get-toa-checkpoint
// @Security ApiKeyAuth
// @Router /api/checkpoint/toa [GET]
// @Produce json
// @Success 200
func (h *CheckpointHandler) toa(c echo.Context) error {
	toa, err := h.cpRepo.FindByDevice("toa")
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get toa checkpoint: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "Toa checkpoint not found")
	}
	return c.JSON(http.StatusOK, toa)
}

// ups
// @Tags Checkpoint
// @Summary Get ups checkpoint
// @ID get-ups-checkpoint
// @Security ApiKeyAuth
// @Router /api/checkpoint/ups [GET]
// @Produce json
// @Success 200
func (h *CheckpointHandler) ups(c echo.Context) error {
	ups, err := h.cpRepo.FindByDevice(_const.Ups)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get ups checkpoint: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "Ups checkpoint not found")
	}
	return c.JSON(http.StatusOK, ups)
}

// updateCCTV
// @Tags Checkpoint
// @Summary Update cctv checkpoint
// @ID update-cctv-checkpoint
// @Security ApiKeyAuth
// @Router /api/checkpoint/cctv [PUT]
// @Accept json
// @Produce json
// @Param checkpoint body updateCheckpointForm true "Checkpoint"
// @Success 200
func (h *CheckpointHandler) updateCCTV(c echo.Context) error {
	nc := c.(*context.Context)

	f, err := newUpdateCheckpointForm(c)
	if err != nil {
		return err
	}

	cctv, err := h.cpRepo.FindByDevice(_const.CCTV)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get cctv checkpoint: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "CCTV checkpoint not found")
	}

	cctv.Checkpoint = f.Checkpoint
	cctv.Updated = nc.Claims.ByAtPtr()

	err = h.cpRepo.UpdateByDevice(_const.CCTV, cctv)
	if err != nil {
		log.Errorf("Failed to update cctv checkpoint: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, "CCTV checkpoint updated")
}

// updateFingerprint
// @Tags Checkpoint
// @Summary Update fingerprint checkpoint
// @ID update-fingerprint-checkpoint
// @Security ApiKeyAuth
// @Router /api/checkpoint/fingerprint [PUT]
// @Accept json
// @Produce json
// @Param checkpoint body updateCheckpointForm true "Checkpoint"
// @Success 200
func (h *CheckpointHandler) updateFingerprint(c echo.Context) error {
	nc := c.(*context.Context)

	f, err := newUpdateCheckpointForm(c)
	if err != nil {
		return err
	}

	fingerprint, err := h.cpRepo.FindByDevice(_const.Fingerprint)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get fingerprint checkpoint: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "Fingerprint checkpoint not found")
	}

	fingerprint.Checkpoint = f.Checkpoint
	fingerprint.Updated = nc.Claims.ByAtPtr()

	err = h.cpRepo.UpdateByDevice(_const.Fingerprint, fingerprint)
	if err != nil {
		log.Errorf("Failed to update fingerprint checkpoint: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, "Fingerprint checkpoint updated")
}

// updateKomputerPh1
// @Tags Checkpoint
// @Summary Update komputer-ph1 checkpoint
// @ID update-komputer-ph1-checkpoint
// @Security ApiKeyAuth
// @Router /api/checkpoint/komputer-ph1 [PUT]
// @Accept json
// @Produce json
// @Param checkpoint body updateCheckpointForm true "Checkpoint"
// @Success 200
func (h *CheckpointHandler) updateKomputerPh1(c echo.Context) error {
	nc := c.(*context.Context)

	f, err := newUpdateCheckpointForm(c)
	if err != nil {
		return err
	}

	komputerPh1, err := h.cpRepo.FindByDevice(_const.KomputerPH1)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get komputer-ph1 checkpoint: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "Komputer-ph1 checkpoint not found")
	}

	komputerPh1.Checkpoint = f.Checkpoint
	komputerPh1.Updated = nc.Claims.ByAtPtr()

	err = h.cpRepo.UpdateByDevice(_const.KomputerPH1, komputerPh1)
	if err != nil {
		log.Errorf("Failed to update komputer-ph1 checkpoint: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, "Komputer-ph1 checkpoint updated")
}

// updateKomputerPh2
// @Tags Checkpoint
// @Summary Update komputer-ph2 checkpoint
// @ID update-komputer-ph2-checkpoint
// @Security ApiKeyAuth
// @Router /api/checkpoint/komputer-ph2 [PUT]
// @Accept json
// @Produce json
// @Param checkpoint body updateCheckpointForm true "Checkpoint"
// @Success 200
func (h *CheckpointHandler) updateKomputerPh2(c echo.Context) error {
	nc := c.(*context.Context)

	f, err := newUpdateCheckpointForm(c)
	if err != nil {
		return err
	}

	komputerPh2, err := h.cpRepo.FindByDevice(_const.KomputerPH2)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get komputer-ph2 checkpoint: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "Komputer-ph2 checkpoint not found")
	}

	komputerPh2.Checkpoint = f.Checkpoint
	komputerPh2.Updated = nc.Claims.ByAtPtr()

	err = h.cpRepo.UpdateByDevice(_const.KomputerPH2, komputerPh2)
	if err != nil {
		log.Errorf("Failed to update komputer-ph2 checkpoint: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, "Komputer-ph2 checkpoint updated")
}

// updatePrinter
// @Tags Checkpoint
// @Summary Update printer checkpoint
// @ID update-printer-checkpoint
// @Security ApiKeyAuth
// @Router /api/checkpoint/printer [PUT]
// @Accept json
// @Produce json
// @Param checkpoint body updateCheckpointForm true "Checkpoint"
// @Success 200
func (h *CheckpointHandler) updatePrinter(c echo.Context) error {
	nc := c.(*context.Context)

	f, err := newUpdateCheckpointForm(c)
	if err != nil {
		return err
	}

	printer, err := h.cpRepo.FindByDevice(_const.Printer)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get printer checkpoint: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "Printer checkpoint not found")
	}

	printer.Checkpoint = f.Checkpoint
	printer.Updated = nc.Claims.ByAtPtr()

	err = h.cpRepo.UpdateByDevice(_const.Printer, printer)
	if err != nil {
		log.Errorf("Failed to update printer checkpoint: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, "Printer checkpoint updated")
}

// updateTelepon
// @Tags Checkpoint
// @Summary Update telepon checkpoint
// @ID update-telepon-checkpoint
// @Security ApiKeyAuth
// @Router /api/checkpoint/telepon [PUT]
// @Accept json
// @Produce json
// @Param checkpoint body updateCheckpointForm true "Checkpoint"
// @Success 200
func (h *CheckpointHandler) updateTelepon(c echo.Context) error {
	nc := c.(*context.Context)

	f, err := newUpdateCheckpointForm(c)
	if err != nil {
		return err
	}

	telepon, err := h.cpRepo.FindByDevice(_const.Telepon)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get telepon checkpoint: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "Telepon checkpoint not found")
	}

	telepon.Checkpoint = f.Checkpoint
	telepon.Updated = nc.Claims.ByAtPtr()

	err = h.cpRepo.UpdateByDevice(_const.Telepon, telepon)
	if err != nil {
		log.Errorf("Failed to update telepon checkpoint: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, "Telepon checkpoint updated")
}

// updateToa
// @Tags Checkpoint
// @Summary Update toa checkpoint
// @ID update-toa-checkpoint
// @Security ApiKeyAuth
// @Router /api/checkpoint/toa [PUT]
// @Accept json
// @Produce json
// @Param checkpoint body updateCheckpointForm true "Checkpoint"
// @Success 200
func (h *CheckpointHandler) updateToa(c echo.Context) error {
	nc := c.(*context.Context)

	f, err := newUpdateCheckpointForm(c)
	if err != nil {
		return err
	}

	toa, err := h.cpRepo.FindByDevice(_const.Toa)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get toa checkpoint: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "Toa checkpoint not found")
	}

	toa.Checkpoint = f.Checkpoint
	toa.Updated = nc.Claims.ByAtPtr()

	err = h.cpRepo.UpdateByDevice(_const.Toa, toa)
	if err != nil {
		log.Errorf("Failed to update toa checkpoint: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, "Toa checkpoint updated")
}

// updateUps
// @Tags Checkpoint
// @Summary Update ups checkpoint
// @ID update-ups-checkpoint
// @Security ApiKeyAuth
// @Router /api/checkpoint/ups [PUT]
// @Accept json
// @Produce json
// @Param checkpoint body updateCheckpointForm true "Checkpoint"
// @Success 200
func (h *CheckpointHandler) updateUps(c echo.Context) error {
	nc := c.(*context.Context)

	f, err := newUpdateCheckpointForm(c)
	if err != nil {
		return err
	}

	ups, err := h.cpRepo.FindByDevice(_const.Ups)
	if err != nil {
		if !errors.Is(err, mongo.ErrNoDocuments) {
			log.Errorf("Failed to get ups checkpoint: %v", err)
			return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
		}
		return echo.NewHTTPError(http.StatusNotFound, "Ups checkpoint not found")
	}

	ups.Checkpoint = f.Checkpoint
	ups.Updated = nc.Claims.ByAtPtr()

	err = h.cpRepo.UpdateByDevice(_const.Ups, ups)
	if err != nil {
		log.Errorf("Failed to update ups checkpoint: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, "Ups checkpoint updated")
}
