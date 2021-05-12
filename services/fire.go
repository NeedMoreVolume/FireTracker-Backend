package services

import (
	"fmt"
	"github.com/NeedMoreVolume/FireTracker/gen/fire"
	"github.com/NeedMoreVolume/FireTracker/models"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type FireService struct {
	logger zerolog.Logger
	db     *gorm.DB
}

// NewFire returns the fire service implementation.
func NewFireService(logger zerolog.Logger, db *gorm.DB) *FireService {
	return &FireService{
		logger.With().Str("services", "fire").Logger(),
		db,
	}
}

// Create a fire and optional payloads
func (s *FireService) Create(fire models.Fire) (out models.Fire, err error) {

	err = s.db.Create(&fire).Error
	if err != nil {
		s.logger.Error().Err(err).Msg("could not create fire")
	}

	out = fire

	return
}

// Get fire and data friends
func (s *FireService) Get(id uint) (out models.Fire, err error) {

	err = s.db.Preload("Weather").Preload("Logs").Preload("Logs.Weather").First(&out, id).Error
	if err != nil {
		s.logger.Error().Err(err).Msgf("failed to find fire %d", id)
	}

	s.logger.Debug().Msgf("len of logs: %d", len(out.Logs))
	s.logger.Debug().Msgf("len of weather: %d", len(out.Weather))

	return
}

// Update something about a fire specifically
func (s *FireService) Update(fire models.Fire) (out models.Fire, err error) {

	err = s.db.Save(&fire).Error
	if err != nil {
		s.logger.Error().Err(err).Msg("failed to save fire")
	}

	out = fire

	return
}

// Update something about a fire specifically
func (s *FireService) Delete(id uint) (err error) {

	err = s.db.Delete(&models.Fire{}, id).Error
	if err != nil {
		s.logger.Error().Err(err).Msgf("failed to delete fire %d", id)
	}

	return
}

// List fires
func (s *FireService) List(payload *fire.FireListPayload) (out []models.Fire, err error) {

	d := s.db.Preload("Weather").Preload("Logs").Preload("Logs.Weather")

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

	var total int64
	err = d.Find(&out).Count(&total).Error
	if err != nil {
		s.logger.Error().Err(err).Msg("failed to find fires")
		return
	}

	s.logger.Debug().Msgf("len of fires: %d", len(out))
	s.logger.Debug().Msgf("total fires for query: %d", int(total))

	return
}
