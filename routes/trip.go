package routes

import (
	"travelfanapi/api"
	"github.com/labstack/echo/v4"
)

func appendTripRoutes(e *echo.Group) {
	e.GET("", api.GetAllTrips)
	e.POST("/new", api.CreateTrip)
	e.GET("/:tripID", api.GetTripByID)
	e.GET("/:location", api.GetTripsByLocation)
	e.DELETE("/:tripID", api.DeleteTripByID)
}
