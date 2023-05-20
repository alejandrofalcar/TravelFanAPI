package api

import (
	"fmt"
	"net/http"
	"strconv"
	"travelfanapi/dao"
	"travelfanapi/domain"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/url"

)

func GetTripsByLocation(c echo.Context) error {
	encodedLocation := c.Param("location")
	
	// Decodifica la ubicación
	decodedLocation, err := url.QueryUnescape(encodedLocation)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Error al decodificar la ubicación")
	}
	
	db := c.Get("DB").(*gorm.DB)
	trips, err := dao.GetTripsByLocation(db, decodedLocation)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	
	return c.JSON(http.StatusOK, trips)
}


func GetTripByID(c echo.Context) error {
	trip, isOK := c.Get("trip").(*domain.Trip)
	if !isOK {
		db := c.Get("DB").(*gorm.DB)
		tripID, err := strconv.ParseUint(c.Param("tripID"), 10, 0)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid trip ID %v", c.Param("tripID")))
		}
		trip, err = dao.GetTripByID(db, uint(tripID))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Trip not found for ID %v", uint(tripID)))
			}
			return err
		}
	}

	return c.JSON(http.StatusOK, trip)
}

func GetTripByUserID(c echo.Context) error {
	db := c.Get("DB").(*gorm.DB)
	userID, err := strconv.ParseUint(c.Param("userID"), 10, 0)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid user ID %v", c.Param("userID")))
	}
	trips, err := dao.GetTripsByUserId(db, uint(userID))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, trips)
}

func GetAllTrips(c echo.Context) error {
	var trip []domain.Trip

	var err error

	db := c.Get("DB").(*gorm.DB)

	trip, err = dao.GetAllTrips(db)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, trip)
}

func CreateTrip(c echo.Context) error {
	db := c.Get("DB").(*gorm.DB)
	var trip domain.Trip
	if err := c.Bind(&trip); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	result, err := dao.CreateTrip(db, trip)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteTripByID(c echo.Context) error {
	tripID, err := strconv.ParseUint(c.Param("tripID"), 10, 0)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid trip ID %s", c.Param("tripID")))
	}

	db := c.Get("DB").(*gorm.DB)

	if err := dao.DeleteTripByID(db, uint(tripID)); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func UpdateTripByID(c echo.Context) error {
	db := c.Get("DB").(*gorm.DB)
	var body domain.Trip
	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var isOK bool
	var tripToModify *domain.Trip
	tripToModify, isOK = c.Get("trip").(*domain.Trip)
	if !isOK {
		tripID, err := strconv.ParseUint(c.Param("tripID"), 10, 0)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid trip ID %v", c.Param("tripID")))
		}
		tripToModify, err = dao.GetTripByID(db, uint(tripID))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Trip not found for ID %v", uint(tripID)))
			}
			return err
		}
	}
	//The modify call domain not to DAO. In this case don't need
	tripToModify.Modify(body)
	if err := tripToModify.Save(db); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, tripToModify)
}
