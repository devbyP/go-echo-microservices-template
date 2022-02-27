package server

import (
	"github.com/devbyP/untitled/server/handlers"
	"github.com/labstack/echo/v4"
)

func uriMapping(e *echo.Echo) {
	e.GET("/", handlers.HelloHandler)
	e.GET("user", handlers.FetchAllUsersHandler)
	e.GET("/user/:id", handlers.FetchUserHandler)
	e.POST("/sign-up", handlers.SignUpHandler)
	e.POST("/sign-in", handlers.SignInHandler)
}
