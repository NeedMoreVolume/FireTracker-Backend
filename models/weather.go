package models

import (
	"gorm.io/gorm"
	"time"
)

type Weather struct {
	gorm.Model
	FireID        *uint
	LogID         *uint
	Temp          int64
	TempUnit      string
	Humidity      uint // humidity is stored as a % - 10.00 => 1000
	DewPoint      int64
	DewPointUnit  string
	WindSpeed     uint
	WindUnit      string
	WindDirection string
	Type          string    // type of weather outside
	WeatherTime   time.Time // timestamp for the weather observation
}
