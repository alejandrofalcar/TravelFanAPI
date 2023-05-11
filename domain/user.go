package domain

import (
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Name      string         `gorm:"column:name;not null"`
	Email     string         `gorm:"column:email;not null;unique"`
	Password  string         `gorm:"column:password;not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (d *User) Modify(mod User) {
	if mod.ID != d.ID {
		d.ID = mod.ID
	}
	if mod.Name != d.Name {
		d.Name = mod.Name
	}
	if mod.Email != d.Email {
		d.Email = mod.Email
	}
	if mod.CreatedAt != d.CreatedAt {
		d.CreatedAt = mod.CreatedAt
	}
	if mod.UpdatedAt != d.UpdatedAt {
		d.UpdatedAt = mod.UpdatedAt
	}
	if mod.DeletedAt != d.DeletedAt {
		d.DeletedAt = mod.DeletedAt
	}
}

func (d *User) Save(db *gorm.DB) error {
	result := db.Save(d)
	if result.Error != nil {
		logrus.Errorf("Error in domain.User-> Error while saving the driver: %s", result.Error)
		return result.Error
	}
	return nil
}

func (ti *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(ti.Password), []byte(password))
	return err == nil
}