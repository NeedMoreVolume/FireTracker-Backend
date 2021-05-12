package controllers

import (
	"context"
	"errors"
	"github.com/NeedMoreVolume/FireTracker/gen/weather"
	"github.com/NeedMoreVolume/FireTracker/services"
	"github.com/rs/zerolog"
	"time"
)

// weather service example implementation.
// The example methods log the requests and return zero values.
type weatherController struct {
	logger         zerolog.Logger
	weatherService *services.WeatherService
}

// NewWeather returns the weather service implementation.
func NewWeather(logger zerolog.Logger, ws *services.WeatherService) weather.Service {
	return &weatherController{
		logger.With().Str("controller", "weather").Logger(),
		ws,
	}
}

// Create a weather observation and optional payloads
func (s *weatherController) Create(_ context.Context, p *weather.Weather) (res *weather.Weather, err error) {
	res = &weather.Weather{}
	s.logger.Debug().Msg("weather.create")

	newWeather, err := weatherToModel(p)
	if err != nil {
		err = weather.MakeBadRequest(err)
		return
	}

	newWeather, err = s.weatherService.Create(newWeather)
	if err != nil {
		return
	}

	res = weatherToTransport(newWeather)
	return
}

// Get a specific piece of weather data
func (s *weatherController) Get(_ context.Context, p *weather.Weather) (res *weather.Weather, err error) {
	res = &weather.Weather{}
	s.logger.Debug().Msg("weather.get")

	if p.ID == nil {
		err = weather.MakeBadRequest(errors.New("id is required"))
	}

	thisWeather, err := s.weatherService.Get(uint(*p.ID))
	if err != nil {
		return
	}

	res = weatherToTransport(thisWeather)

	return
}

// Update something about a weather observation specifically
func (s *weatherController) Update(_ context.Context, p *weather.Weather) (res *weather.Weather, err error) {
	res = &weather.Weather{}
	s.logger.Debug().Msg("weather.update")

	if p.ID == nil {
		err = weather.MakeBadRequest(errors.New("id is required"))
		return
	}

	tempWeather, err := s.weatherService.Get(uint(*p.ID))
	if err != nil {
		return
	}
	if p.WeatherTime != nil {
		tempWeather.WeatherTime, err = time.Parse(time.RFC3339, *p.WeatherTime)
		if err != nil {
			err = weather.MakeBadRequest(errors.New("weatherTime must be un RFC3339 format"))
		}
	}
	if p.Temperature != nil {
		if p.Temperature.Unit != nil {
			tempWeather.TempUnit = *p.Temperature.Unit
		}
		if p.Temperature.Value != nil {
			tempWeather.Temp = int64(*p.Temperature.Value)
		}
	}
	if p.FireID != nil {
		fireID := uint(*p.FireID)
		tempWeather.FireID = &fireID
	}
	if p.LogID != nil {
		logID := uint(*p.LogID)
		tempWeather.LogID = &logID
	}
	if p.Humidity != nil {
		tempWeather.Humidity = uint(*p.Humidity)
	}
	if p.DewPoint != nil {
		if p.DewPoint.Unit != nil {
			tempWeather.DewPointUnit = *p.DewPoint.Unit
		}
		if p.DewPoint.Value != nil {
			tempWeather.DewPoint = int64(*p.DewPoint.Value)
		}
	}
	if p.Wind != nil {
		if p.Wind.Unit != nil {
			tempWeather.WindUnit = *p.Wind.Unit
		}
		if p.Wind.Speed != nil {
			tempWeather.WindSpeed = uint(*p.Wind.Speed)
		}
		if p.Wind.Direction != nil {
			tempWeather.WindDirection = *p.Wind.Direction
		}
	}
	if p.WeatherType != nil {
		tempWeather.Type = *p.WeatherType
	}

	updatedWeather, err := s.weatherService.Update(tempWeather)
	if err != nil {
		return
	}

	res = weatherToTransport(updatedWeather)

	return
}

// Delete a weather observation
func (s *weatherController) Delete(_ context.Context, p *weather.Weather) (err error) {
	s.logger.Debug().Msg("weather.delete")

	if p.ID == nil || *p.ID == 0 {
		err = weather.MakeBadRequest(errors.New("must provide a valid id"))
		return
	}

	err = s.weatherService.Delete(uint(*p.ID))

	return
}

// List weather
func (s *weatherController) List(_ context.Context, p *weather.WeatherListPayload) (res *weather.WeatherList, err error) {
	s.logger.Debug().Msg("weather.list")

	weathers, count, err := s.weatherService.List(p)
	if err != nil {
		return
	}

	total := int(count)
	res = &weather.WeatherList{
		Weathers: []*weather.Weather{},
		Pagination: &weather.Pagination{
			Total: &total,
		},
	}

	for _, thisWeather := range weathers {
		res.Weathers = append(res.Weathers, weatherToTransport(thisWeather))
	}

	return
}
