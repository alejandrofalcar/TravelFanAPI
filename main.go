package main

import (
	"travelfanapi/app"
	"travelfanapi/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e = app.InitConnection(*e)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.InitializeRoutes(e)

	e.Logger.Fatal(e.Start(":8081"))
}
