package model

import (
	"database/sql"
	"time"
)

type Location struct {
	GeonameId string `gorm:"primaryKey"`
	Name      string
	Feature   uint
	Latitude  float32
	Longitude float32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

func (Location) TableName() string {
	return "location"
}
