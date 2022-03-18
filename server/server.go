package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	myjwt "github.com/devbyP/untitled/pkg/jwt"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var port string

func tokenMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		TokenLookup:   "header:Authorization",
		AuthScheme:    "JWT",
		SigningMethod: middleware.AlgorithmHS256,
		ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
			token, err := jwt.ParseWithClaims(auth, &jwt.StandardClaims{}, myjwt.MyKeyFunc)
			if err != nil {
				return nil, err
			}
			if !token.Valid {
				return nil, fmt.Errorf("invalid token")
			}
			return token, nil
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
	e.Use(tokenMiddleware())
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
