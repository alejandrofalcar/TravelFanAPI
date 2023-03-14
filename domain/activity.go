package domain

import (
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Activity struct {
	ID          uint   `gorm:"column:ID;primary_key;AUTO_INCREMENT;NOT NULL"`
	Date        time.Time `gorm:"column:date;not null"`
	Image       string `gorm:"column:image;not null"`
	Type        string `gorm:"column:type;not null"`
	Time        string `gorm:"column:time;not null"`
	Location    string `gorm:"column:location;not null"`
	Latitude    float64 `gorm:"column:latitude;not null"`
	Longitude   float64 `gorm:"column:longitude;not null"`
	Name        string `gorm:"column:name;not null"`
	Description string `gorm:"column:description;not null"`
	TripID      uint `gorm:"column:trip_id;not null"`
	Trip        Trip `gorm:"foreignKey:TripID"`
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoCreateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (d *Activity) Modify(mod Activity) {
	if mod.ID != d.ID {
		d.ID = mod.ID
	}
	if mod.Date != d.Date {
		d.Date = mod.Date
	}
	if mod.Image != d.Image {
		d.Image = mod.Image
	}
	if mod.Type != d.Type {
		d.Type = mod.Type
	}
	if mod.Location != d.Location {
		d.Location = mod.Location
	}
	if mod.Latitude != d.Latitude {
		d.Latitude = mod.Latitude
	}
	if mod.Longitude != d.Longitude {
		d.Longitude = mod.Longitude
	}
	if mod.Name != d.Name {
		d.Name = mod.Name
	}
	if mod.Description != d.Description {
		d.Description = mod.Description
	}
	if mod.TripID != d.TripID {
		d.TripID = mod.TripID
	}
	if mod.UpdatedAt != d.UpdatedAt {
		d.UpdatedAt = mod.UpdatedAt
	}
	if mod.DeletedAt != d.DeletedAt {
		d.DeletedAt = mod.DeletedAt
	}
}

func (d *Activity) Save(db *gorm.DB) error {
	result := db.Save(d)
	if result.Error != nil {
		logrus.Errorf("Error in domain.Activity-> Error while saving the driver: %s", result.Error)
		return result.Error
	}
	return nil
}