package server

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var port string

func customLogger() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rtc3339_nano}] - method:${method}, uri:${uri}, status:${status}\n",
	})
}

func SetPort(p string) {
	port = p
}

func PrintPort() {
	log.Println("server port: " + port)
}

func serverMiddlewares(e *echo.Echo) {
	e.Use(customLogger())
	e.Use(middleware.Recover())
}

func StartServer() {
	if port == "" {
		port = os.Getenv("PORT")
	}
	e := echo.New()

	serverMiddlewares(e)
	uriMapping(e)

	e.Logger.Fatal(e.Start(":" + port))
}
