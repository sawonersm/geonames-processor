package model

import (
	"database/sql"
	"time"
)

type Feature struct {
	Id          uint `gorm:"primaryKey;autoIncrement"`
	Code        string
	Class       string
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
}

func (Feature) TableName() string {
	return "feature"
}
