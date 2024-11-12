package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoswagger "github.com/swaggo/echo-swagger"
	"github.com/unrolled/secure"
	"net/http"
	"sipamit-be/api"
	"sipamit-be/docs"
	"sipamit-be/internal/config"
	_db "sipamit-be/internal/db"
	"sipamit-be/internal/pkg/log"
	"sipamit-be/version"
	"strings"
)

// @title Sistem Pencatatan Maintenance IT Backend
// @description Sistem Pencatatan Maintenance IT Backend API
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

const AppName = "Sistem Pencatatan Maintenance IT Backend"

func main() {
	defer log.RecoverWithTrace()

	e := echo.New()
	log.SetLogger(e)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     strings.Split(config.App.CORSAllowOrigins, ","),
		AllowCredentials: true,
	}))
	e.Use(middleware.Gzip())

	secureMiddleware := secure.New(secure.Options{
		FrameDeny:               true,         // avoid clickjacking
		CustomFrameOptionsValue: "SAMEORIGIN", // only allow same origin iframe
		ContentTypeNosniff:      true,         // avoid MIME sniffing
		BrowserXssFilter:        true,         // avoid XSS attacks
	})
	e.Use(echo.WrapMiddleware(secureMiddleware.Handler))

	if config.App.Debug {
		e.GET("/", func(c echo.Context) error {
			return c.HTML(http.StatusOK, fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
    <head>
        <title>%v</title>
    </head>
    <body>
  		<h1>Welcome to %v</h1>
  		<p><a href="/api/version">version: %v</a></p>
  		<p><a href="/swagger/index.html#/">docs</a></p>
	</body>
</html>`, AppName, AppName, version.Version))
		})
		e.GET("/api/version", func(c echo.Context) error {
			return c.JSON(http.StatusOK, map[string]string{
				"name":    AppName,
				"version": version.Version,
			})
		})
		e.GET("/swagger/*", echoswagger.WrapHandler)
	}

	docs.SwaggerInfo.Version = version.Version
	docs.SwaggerInfo.Host = config.App.SwaggerHost

	api.NewInitHandler(e, _db.Client)

	log.Fatal(e.Start(fmt.Sprintf(`%v:%v`, config.App.Host, config.App.Port)))
}
