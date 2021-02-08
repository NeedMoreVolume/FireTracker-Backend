package models

import (
	"gorm.io/gorm"
	"time"
)

type Fire struct {
	gorm.Model
	Name        string
	Description string
	Start       time.Time  // start of the fire
	End         *time.Time // end of the fire
	Logs        []Log      // log additions to the fire
	Weather     []Weather  // weather observations for the fire
}
