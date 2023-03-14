package dao

import (
	"travelfanapi/domain"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func CreateTrip(db *gorm.DB, trip domain.Trip) (*domain.Trip, error) {
	result := db.Create(&trip)
	if result.Error != nil {
		logrus.Errorf("Error in dao.CreateTrip -> error: %s", result.Error)
		return nil, result.Error
	}

	return &trip, nil
}

func GetAllTrips(db *gorm.DB) ([]domain.Trip, error) {
	res := []domain.Trip{}

	result := db.Preload("Activities.Trip").Find(&res)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
			logrus.Errorf("Error in dao.GetAllTrips -> error: %s", result.Error)
			return res, result.Error
	}

	return res, nil
}


func GetTripByID(db *gorm.DB, tripID uint) (*domain.Trip, error) {
	var res domain.Trip

	result := db.Preload("Activities.Trip").First(&res, tripID)
	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			logrus.Errorf("Error in dao.GetTripById -> error: %s", result.Error)
		}
		return nil, result.Error
	}

	return &res, nil
}


func DeleteTripByID(db *gorm.DB, tripID uint) error {
	result := db.Where("ID = ?", tripID).Delete(&domain.Trip{})
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		logrus.Errorf("Error in dao.DeteleTripById -> error: %s", result.Error)
		return result.Error
	}
	return nil
}
