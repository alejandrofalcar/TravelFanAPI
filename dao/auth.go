package dao

import (
	"travelfanapi/domain"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func GetUserByLogin(db *gorm.DB, login string) (*domain.User, error) {
	var res domain.User

	result := db.Where("email = ?", login).First(&res)
	if result.Error != nil {
		if result.Error != gorm.ErrRecordNotFound {
			logrus.Errorf("Error in dao.GetUserByLogin -> error: %s", result.Error)
		}
		return nil, result.Error
	}

	return &res, nil
}
