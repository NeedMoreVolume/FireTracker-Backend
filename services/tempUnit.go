package services

import (
	"github.com/NeedMoreVolume/FireTracker/models"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type TempUnitService struct {
	logger zerolog.Logger
	db     *gorm.DB
}

// NewWeather returns the weather service implementation.
func NewTempUnit(logger zerolog.Logger, db *gorm.DB) *TempUnitService {
	return &TempUnitService{
		logger.With().Str("services", "tempUnit").Logger(),
		db,
	}
}

func (s *TempUnitService) GetByUnit(unit string) (out models.TempUnit, err error) {

	err = s.db.First(&out, "unit = ?", unit).Error
	if err != nil {
		s.logger.Error().Err(err).Msg("failed to get temp unit")
	}

	return
}
