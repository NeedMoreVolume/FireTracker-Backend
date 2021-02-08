package services

import (
	"github.com/NeedMoreVolume/FireTracker/models"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type SpeedUnitService struct {
	logger zerolog.Logger
	db *gorm.DB
}

func NewSpeedUnit(logger zerolog.Logger, db *gorm.DB) *SpeedUnitService {
	return &SpeedUnitService{
		logger.With().Str("services", "speedUnit").Logger(),
		db,
	}
}

func (s *SpeedUnitService) GetByUnit(unit string) (out models.SpeedUnit, err error) {

	err = s.db.Find(&out, "unit = ?", unit).Error
	if err != nil {
		s.logger.Error().Err(err).Msg("failed to find speed unit")
	}

	return
}