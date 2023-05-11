package routes

import (
	"travelfanapi/api"
	"github.com/labstack/echo/v4"
)

func appendUserRoutes(e *echo.Group) {
	e.POST("/register", api.Register)
	e.POST("/login", api.Authenticate)
	e.GET("/:userID", api.GetTripByUserID)
}