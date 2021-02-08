package models

import "gorm.io/gorm"

type TempUnit struct {
	gorm.Model
	Unit string
}
