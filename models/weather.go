package models

import (
	"gorm.io/gorm"
	"time"
)

type Weather struct {
	gorm.Model
	FireID        *uint
	LogID         *uint
	Temp          uint
	TempUnitID    uint
	TempUnit      TempUnit
	Humidity      uint // humidity is stored as a % - 10.00 => 1000
	DewPoint      uint
	DewPointUnitID uint
	DewPointUnit  TempUnit
	WindSpeed     uint
	WindUnitID    uint
	WindUnit      SpeedUnit
	WindDirection string
	Type          string    // type of weather outside
	WeatherTime   time.Time // timestamp for the weather observation
}
