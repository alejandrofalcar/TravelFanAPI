package domain

import (
	"time"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Trip struct {
	ID         uint           `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	UserID     uint           `gorm:"column:user_id;not null"`
	User       User           `gorm:"foreignKey:UserID"`
	StartDate  string         `gorm:"column:start_date;not null"`
	EndDate    string         `gorm:"column:end_date;not null"`
	Location   string         `gorm:"column:location;not null"`
	Image      string         `gorm:"column:image;not null"`
	Activities []Activity     `gorm:"foreignKey:TripID;onDelete:CASCADE"`
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoCreateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (d *Trip) Modify(mod Trip) {
	if mod.ID != d.ID {
		d.ID = mod.ID
	}
	if mod.StartDate != d.StartDate {
		d.StartDate = mod.StartDate
	}
	if mod.EndDate != d.EndDate {
		d.EndDate = mod.EndDate
	}
	if mod.Location != d.Location {
		d.Location = mod.Location
	}
	if mod.Image != d.Image {
		d.Image = mod.Image
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

func (d *Trip) Save(db *gorm.DB) error {
	result := db.Save(d)
	if result.Error != nil {
		logrus.Errorf("Error in domain.Trip-> Error while saving the driver: %s", result.Error)
		return result.Error
	}
	return nil
}
