package controllers

import (
	"context"
	"errors"
	"github.com/NeedMoreVolume/FireTracker/models"
	"github.com/NeedMoreVolume/FireTracker/services"
	"github.com/rs/zerolog"
	"time"

	"github.com/NeedMoreVolume/FireTracker/gen/fire"
)

// fire service example implementation.
// The example methods log the requests and return zero values.
type fireController struct {
	logger         zerolog.Logger
	fireService    *services.FireService
	logService     *services.LogService
	weatherService *services.WeatherService
}

// NewFire returns the fire service implementation.
func NewFire(logger zerolog.Logger, fs *services.FireService, ls *services.LogService, ws *services.WeatherService) fire.Service {
	return &fireController{
		logger.With().Str("controller", "fire").Logger(),
		fs,
		ls,
		ws,
	}
}

// Create a fire and optional payloads
func (s *fireController) Create(_ context.Context, p *fire.Fire) (res *fire.Fire, err error) {
	res = &fire.Fire{}
	s.logger.Debug().Msg("fire.create")

	newFire := models.Fire{
		Start: time.Now(),
	}
	if p.Name != nil {
		newFire.Name = *p.Name
	}
	if p.Description != nil {
		newFire.Description = *p.Description
	}
	if p.Start != nil {
		newFire.Start, err = time.Parse(time.RFC3339, *p.Start)
		if err != nil {
			s.logger.Printf("failed to parse start time: %s", *p.Start)
			err = fire.MakeBadRequest(err)
			return
		}
	}

	newFire, err = s.fireService.Create(newFire)
	if err != nil {
		return
	}

	res = fireToTransport(newFire)

	return
}

// Get fire and data friends
func (s *fireController) Get(_ context.Context, p *fire.GetPayload) (res *fire.Fire, err error) {
	res = &fire.Fire{}
	s.logger.Debug().Msg("fire.get")

	if p.ID == nil {
		err = fire.MakeBadRequest(errors.New("id is required"))
		return
	}

	thisFire, err := s.fireService.Get(uint(*p.ID))
	if err != nil {
		return
	}

	res = fireToTransport(thisFire)

	return
}

// Update something about a fire specifically
func (s *fireController) Update(_ context.Context, p *fire.Fire) (res *fire.Fire, err error) {
	s.logger.Debug().Msg("fire.update")
	res = &fire.Fire{}

	if p.ID == nil {
		err = fire.MakeBadRequest(errors.New("id is required"))
		return
	}

	var fireUpdate models.Fire
	fireUpdate, err = s.fireService.Get(uint(*p.ID))
	if err != nil {
		return
	}

	if p.Name != nil {
		fireUpdate.Name = *p.Name
	}
	if p.Description != nil {
		fireUpdate.Description = *p.Description
	}
	if p.Start != nil {
		fireUpdate.Start, err = time.Parse(time.RFC3339, *p.Start)
		if err != nil {
			s.logger.Printf("failed to parse start time: %s", *p.Start)
			err = fire.MakeBadRequest(err)
			return
		}
	}
	if p.End != nil {
		var endTime time.Time
		endTime, err = time.Parse(time.RFC3339, *p.End)
		if err != nil {
			s.logger.Printf("failed to parse start time: %s", *p.Start)
			err = fire.MakeBadRequest(err)
			return
		}
		fireUpdate.End = &endTime
	}

	updatedFire, err := s.fireService.Update(fireUpdate)
	if err != nil {
		return
	}

	res = fireToTransport(updatedFire)

	return
}

// Update something about a fire specifically
func (s *fireController) Delete(_ context.Context, p *fire.DeletePayload) (err error) {
	s.logger.Debug().Msg("fire.delete")

	if p.ID == nil || *p.ID == 0 {
		err = fire.MakeBadRequest(errors.New("must provide a valid id to delete by"))
		return
	}

	err = s.fireService.Delete(uint(*p.ID))

	return
}

// List fires
func (s *fireController) List(_ context.Context, p *fire.FireListPayload) (res *fire.FireList, err error) {
	res = &fire.FireList{}
	s.logger.Debug().Msg("fire.list")

	fires, err := s.fireService.List(p)
	if err != nil {
		return
	}

	for _, thisFire := range fires {
		res.Fires = append(res.Fires, fireToTransport(thisFire))
	}

	return
}

// Gets a list of weather for a fire
func (s *fireController) GetWeatherForFire(_ context.Context, p *fire.GetWeatherForFirePayload) (res *fire.WeatherList, err error) {
	s.logger.Debug().Msg("fire.getWeatherForFire")

	if p.ID == nil {
		err = fire.MakeBadRequest(errors.New("fire ID is required"))
		return
	}

	weathers, count, err := s.weatherService.ListByFire(uint(*p.ID))
	if err != nil {
		return
	}

	total := int(count)
	res = &fire.WeatherList{
		Weathers: []*fire.Weather{},
		Pagination: &fire.Pagination{
			Total: &total,
		},
	}
	for _, thisWeather := range weathers {
		res.Weathers = append(res.Weathers, fireWeatherToTransport(thisWeather))
	}

	return
}

// Gets a list of logs for a fire
func (s *fireController) GetLogsForFire(_ context.Context, p *fire.GetLogsForFirePayload) (res *fire.LogList, err error) {
	s.logger.Debug().Msg("fire.getLogsForFire")

	if p.ID == nil {
		err = fire.MakeBadRequest(errors.New("fire ID is required"))
		return
	}

	logs, count, err := s.logService.ListByFire(uint(*p.ID))
	if err != nil {
		return
	}

	total := int(count)
	res = &fire.LogList{
		Logs: []*fire.Log{},
		Pagination: &fire.Pagination{
			Total: &total,
		},
	}
	for _, thisLog := range logs {
		res.Logs = append(res.Logs, fireLogToTransport(thisLog))
	}

	return
}
