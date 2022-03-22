package server

import (
	"log"
	"net/http"
	"os"
	"strings"

	myjwt "github.com/devbyP/untitled/pkg/jwt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var port string

func tokenHeaderMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		Skipper: func(c echo.Context) bool {
			if c.Request().Header.Get("Authorization") == "" {
				return true
			}
			return false
		},
		TokenLookup:   "header:Authorization",
		AuthScheme:    "JWT",
		SigningMethod: middleware.AlgorithmHS256,
		ParseTokenFunc: func(auth string, c echo.Context) (any, error) {
			return myjwt.ValidateToken(auth)
		},
		ErrorHandlerWithContext: func(err error, c echo.Context) error {
			if strings.Contains(err.Error(), "token is expired") {
				return c.JSON(http.StatusUnauthorized, map[string]string{"message": "token is expired"})
			}
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "not authorize"})
		},
	})
}

func tokenCookieMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
    Skipper: func(c echo.Context) bool {
			if c.Request().Header.Get("Authorization") != "" {
        return true
      }
      return false
    },
		TokenLookup:   "cookie:aToken",
		AuthScheme:    "",
		SigningMethod: middleware.AlgorithmHS256,
		ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
			return myjwt.ValidateToken(auth)
		},
		ErrorHandlerWithContext: func(err error, c echo.Context) error {
			if strings.Contains(err.Error(), "token is expired") {
				return c.JSON(http.StatusUnauthorized, map[string]string{"message": "token is expired"})
			}
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "not authorize"})
		},
	})
}

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
	e.Use(tokenHeaderMiddleware())
	e.Use(tokenCookieMiddleware())
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
