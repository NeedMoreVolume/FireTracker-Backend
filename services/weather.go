package services

import (
	"fmt"
	"github.com/NeedMoreVolume/FireTracker/gen/weather"
	"github.com/NeedMoreVolume/FireTracker/models"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type WeatherService struct {
	logger zerolog.Logger
	db     *gorm.DB
}

// NewFire returns the fire service implementation.
func NewWeatherService(logger zerolog.Logger, db *gorm.DB) *WeatherService {
	return &WeatherService{
		logger.With().Str("services", "weather").Logger(),
		db,
	}
}

// Create a fire and optional payloads
func (s *WeatherService) Create(weather models.Weather) (out models.Weather, err error) {

	err = s.db.Create(&weather).Error
	if err != nil {
		s.logger.Error().Err(err).Msg("could not create weather")
	}

	out = weather

	return
}

// Get fire and data friends
func (s *WeatherService) Get(id uint) (out models.Weather, err error) {

	err = s.db.First(&out, id).Error
	if err != nil {
		s.logger.Error().Err(err).Msgf("failed to find fire %d", id)
	}

	return
}

// Update something about a fire specifically
func (s *WeatherService) Update(weather models.Weather) (out models.Weather, err error) {

	err = s.db.Save(&weather).Error
	if err != nil {
		s.logger.Error().Err(err).Msg("failed to save fire")
	}

	out = weather

	return
}

// Update something about a fire specifically
func (s *WeatherService) Delete(id uint) (err error) {

	err = s.db.Delete(&models.Weather{}, id).Error
	if err != nil {
		s.logger.Error().Err(err).Msgf("failed to delete fire %d", id)
	}

	return
}

// List fires
func (s *WeatherService) List(payload *weather.WeatherListPayload) (out []models.Weather, count int64, err error) {

	if payload != nil {
		if payload.Pagination != nil {
			// apply pagination
			s.db = s.db.Limit(payload.Pagination.Limit)
			offset := payload.Pagination.Limit * (payload.Pagination.Page - 1)
			s.db = s.db.Offset(offset)
		}
		if payload.Sort != nil {
			// apply sort
			if payload.Sort.ID != nil {
				query := fmt.Sprintf("id %s", *payload.Sort.ID)
				s.db = s.db.Order(query)
			}
		}
		if payload.Search != nil {
			// apply search
			if payload.Search.Name != nil {
				s.db = s.db.Where("name LIKE %?%", *payload.Search.Name)
			}
			if payload.Search.Description != nil {
				s.db = s.db.Where("description LIKE %?%", *payload.Search.Description)
			}
		}
		if payload.Filters != nil {
			// apply filters
			if payload.Filters.Time != nil {
				// apply time filters
				for _, filter := range payload.Filters.Time {
					query := fmt.Sprintf("%s %s ?", filter.Key, filter.Operator)
					s.db = s.db.Where(query, filter.Value)
				}
			}
			if payload.Filters.Humidity != nil {
				// apply hummidity filters
				for _, filter := range payload.Filters.Humidity {
					query := fmt.Sprintf("%s %s ?", filter.Key, filter.Operator)
					s.db = s.db.Where(query, filter.Value)
				}
			}
			if payload.Filters.Temperature != nil {
				// apply temp filters
				for _, filter := range payload.Filters.Temperature {
					query := fmt.Sprintf("%s %s ?", filter.Key, filter.Operator)
					s.db = s.db.Where(query, filter.Value)
				}
			}
		}
	}

	err = s.db.Find(&out).Count(&count).Error
	if err != nil {
		s.logger.Error().Err(err).Msgf("failed to find weathers")
		return
	}

	s.logger.Debug().Msgf("len of weathers: %d", len(out))
	s.logger.Debug().Msgf("total weathers for query: %d", count)

	return
}

func (s *WeatherService) ListByFire(fireID uint) (out []models.Weather, count int64, err error) {

	err = s.db.Find(&out, "fire_id = ?", fireID).Count(&count).Error
	if err != nil {
		s.logger.Error().Err(err).Msgf("failed to find weather for fire %d", fireID)
	}

	return
}
