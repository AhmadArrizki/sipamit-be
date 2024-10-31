package api

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/mongo"
	appHandler "sipamit-be/api/app/handler"
	deviceHandler "sipamit-be/api/device/handler"
	checkpointHandler "sipamit-be/api/device_cp/handler"
	deviceDocHandler "sipamit-be/api/device_doc/handler"
)

func NewInitHandler(e *echo.Echo, db *mongo.Database) {
	appHandler.NewAuthHandler(e, db)
	appHandler.NewUserHandler(e, db)

	deviceHandler.NewCCTVAPIHandler(e, db)
	deviceHandler.NewFingerPrintAPIHandler(e, db)
	deviceHandler.NewKomputerPH1APIHandler(e, db)
	deviceHandler.NewKomputerPH2APIHandler(e, db)
	deviceHandler.NewPrinterAPIHandler(e, db)
	deviceHandler.NewTeleponAPIHandler(e, db)
	deviceHandler.NewTOAAPIHandler(e, db)
	deviceHandler.NewUPSAPIHandler(e, db)

	checkpointHandler.NewCheckpointAPIHandler(e, db)

	deviceDocHandler.NewCCTVDocAPIHandler(e, db)
	deviceDocHandler.NewFingerprintDocAPIHandler(e, db)
	deviceDocHandler.NewKomputerPH1DocAPIHandler(e, db)
	deviceDocHandler.NewKomputerPH2DocAPIHandler(e, db)
	deviceDocHandler.NewPrinterDocAPIHandler(e, db)
	deviceDocHandler.NewTeleponDocAPIHandler(e, db)
	deviceDocHandler.NewTOADocAPIHandler(e, db)
	deviceDocHandler.NewUPSDocAPIHandler(e, db)
}
