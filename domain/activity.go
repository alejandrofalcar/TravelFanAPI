package domain

import (
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Activity struct {
	ID        uint           `gorm:"column:ID;primary_key;AUTO_INCREMENT;NOT NULL"`
	Date      string         `gorm:"column:date"`
	Image     string         `gorm:"column:image"`
	Type      string         `gorm:"column:type"`
	Time      string         `gorm:"column:time"`
	Location  string         `gorm:"column:location"`
	Latitude  float64        `gorm:"column:latitude"`
	Longitude float64        `gorm:"column:longitude"`
	Street    string         `gorm:"column:street"`
	File      string         `gorm:"column:file"`
	FileData  []byte         `gorm:"column:file_data"`
	TripID    uint           `gorm:"column:trip_id;not null"`
	Trip      Trip           `gorm:"foreignKey:TripID"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (d *Activity) Modify(mod Activity, file []byte) {
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
	if mod.Street != d.Street {
		d.Street = mod.Street
	}
	if mod.File != d.File {
		d.File = mod.File
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

func (d *Activity) Save(db *gorm.DB, file []byte) error {
	result := db.Save(d)
	if result.Error != nil {
		logrus.Errorf("Error in domain.Activity-> Error while saving the driver: %s", result.Error)
		return result.Error
	}
	return nil
}
