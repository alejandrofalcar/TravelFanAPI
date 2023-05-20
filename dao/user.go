package dao

import (
	"errors"
	"fmt"
	"travelfanapi/domain"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func GetUserByID(db *gorm.DB, userID uint) (*domain.User, error) {
	var user domain.User

	result := db.First(&user, userID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("User not found for ID %v", userID)
		}
		logrus.Errorf("Error in dao.GetUserByID -> error: %s", result.Error)
		return nil, result.Error
	}

	return &user, nil
}
