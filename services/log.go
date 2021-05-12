package services

import (
	"fmt"
	logs "github.com/NeedMoreVolume/FireTracker/gen/log"
	"github.com/NeedMoreVolume/FireTracker/models"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type LogService struct {
	logger zerolog.Logger
	db     *gorm.DB
}

// NewLog returns the log service implementation.
func NewLogService(logger zerolog.Logger, db *gorm.DB) *LogService {
	return &LogService{
		logger.With().Str("services", "log").Logger(),
		db,
	}
}

// Create a log
func (s *LogService) Create(log models.Log) (out models.Log, err error) {

	weather := log.Weather
	log.Weather = models.Weather{}
	// start transaction
	tx := s.db.Begin()
	// make log
	err = tx.Create(&log).Error
	if err != nil {
		s.logger.Error().Err(err).Msg("could not create log")
		// roll back transaction
		tx.Rollback()
		return
	}
	out = log

	// make weather if provided
	if weather != log.Weather {
		weather.LogID = &log.ID

		err = tx.Create(&weather).Error
		if err != nil {
			s.logger.Error().Err(err).Msg("could not create log")
			// roll back transaction
			tx.Rollback()
			return
		}
		out.Weather = weather
		out.WeatherID = &weather.ID
	}

	err = tx.Save(&out).Error
	if err != nil {
		s.logger.Error().Err(err).Msg("could not associate weather to log")
		// roll back transaction
		tx.Rollback()
		return
	}

	// commit transaction
	err = tx.Commit().Error
	return
}

// Get fire
func (s *LogService) Get(id uint) (out models.Log, err error) {

	err = s.db.Preload("Weather").First(&out, id).Error
	if err != nil {
		s.logger.Error().Err(err).Msgf("failed to find log %d", id)
	}

	return
}

// Update something about a fire specifically
func (s *LogService) Update(log models.Log) (out models.Log, err error) {

	err = s.db.Save(&log).Error
	if err != nil {
		s.logger.Err(err).Msg("failed to save log")
	}

	out = log

	return
}

// Update something about a fire specifically
func (s *LogService) Delete(id uint) (err error) {

	err = s.db.Delete(&models.Log{}, id).Error
	if err != nil {
		s.logger.Error().Err(err).Msgf("failed to delete log %d", id)
	}

	return
}

// List fires
func (s *LogService) List(payload *logs.LogListPayload) (out []models.Log, count int64, err error) {

	d := s.db.Preload("Weather")

	if payload != nil {
		if payload.Pagination != nil {
			// apply pagination
			d = d.Limit(payload.Pagination.Limit)
			offset := payload.Pagination.Limit * (payload.Pagination.Page - 1)
			d = d.Offset(offset)
		}
		if payload.Sort != nil {
			// apply sort
			if payload.Sort.ID != nil {
				query := fmt.Sprintf("id %s", *payload.Sort.ID)
				d = d.Order(query)
			}
			if payload.Sort.Start != nil {
				query := fmt.Sprintf("name %s", *payload.Sort.Start)
				d = d.Order(query)
			}
			if payload.Sort.End != nil {
				query := fmt.Sprintf("price %s", *payload.Sort.End)
				d = d.Order(query)
			}
		}
		if payload.Search != nil {
			// apply search
			if payload.Search.Name != nil {
				d = d.Where("name LIKE %?%", *payload.Search.Name)
			}
			if payload.Search.Description != nil {
				d = d.Where("description LIKE %?%", *payload.Search.Description)
			}
		}
		if payload.Filters != nil {
			// apply filters
			if payload.Filters.Name != nil {
				// apply name filters
				for _, filter := range payload.Filters.Name {
					query := fmt.Sprintf("%s %s ?", filter.Key, filter.Operator)
					d = d.Where(query, filter.Value)
				}
			}
			if payload.Filters.End != nil {
				// apply endtime filters
				for _, filter := range payload.Filters.Name {
					query := fmt.Sprintf("%s %s ?", filter.Key, filter.Operator)
					d = d.Where(query, filter.Value)
				}
			}
			if payload.Filters.Start != nil {
				// apply starttime filters
				for _, filter := range payload.Filters.Name {
					query := fmt.Sprintf("%s %s ?", filter.Key, filter.Operator)
					d = d.Where(query, filter.Value)
				}
			}
		}
	}

	err = d.Find(&out).Count(&count).Error
	if err != nil {
		s.logger.Error().Err(err).Msgf("failed to find logs")
		return
	}

	s.logger.Debug().Msgf("len of logs: %d", len(out))
	s.logger.Debug().Msgf("total logs for query: %d", count)

	return
}

func (s *LogService) ListByFire(fireID uint) (out []models.Log, count int64, err error) {

	err = s.db.Find(&out, "fire_id = ?", fireID).Count(&count).Error
	if err != nil {
		s.logger.Error().Err(err).Msgf("failed to find logs for fire %d", fireID)
	}

	return
}
