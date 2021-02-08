package models

import "gorm.io/gorm"

type SpeedUnit struct {
	gorm.Model
	Unit string
}
