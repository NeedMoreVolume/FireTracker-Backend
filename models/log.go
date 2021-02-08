package models

import (
	"gorm.io/gorm"
	"time"
)

type Log struct {
	gorm.Model
	FireID uint
	WeatherID *uint
	Name string
	Size string
	AddedAt time.Time
	Weather Weather // used if providing a weather observation with the log addition
}