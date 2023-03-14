package routes

import (

	"github.com/labstack/echo/v4"
)

func InitializeRoutes(e *echo.Echo) {
	appendTripRoutes(e.Group("/trip"))
}
