package app

import (
	"log"
	"travelfanapi/config"
	"travelfanapi/domain"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitConnection(e echo.Echo) *echo.Echo {
	dsn := config.GetDBConnection()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}
	db.AutoMigrate(&domain.Activity{}, domain.Trip{})

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("DB", db)
			return next(c)
		}
	})

	return &e
}
