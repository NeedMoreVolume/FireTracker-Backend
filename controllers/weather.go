package controllers

import (
	"context"
	"errors"
	"github.com/NeedMoreVolume/FireTracker/gen/weather"
	"github.com/NeedMoreVolume/FireTracker/models"
	"github.com/NeedMoreVolume/FireTracker/services"
	"github.com/rs/zerolog"
	"time"
)

// weather service example implementation.
// The example methods log the requests and return zero values.
type weatherController struct {
	logger zerolog.Logger
	weatherService *services.WeatherService
	tempUnitService *services.TempUnitService
	speedUnitService *services.SpeedUnitService
}

// NewWeather returns the weather service implementation.
func NewWeather(logger zerolog.Logger, ws *services.WeatherService, tus *services.TempUnitService, sus *services.SpeedUnitService) weather.Service {
	return &weatherController{
		logger.With().Str("controller", "weather").Logger(),
		ws,
		tus,
		sus,
	}
}

// Create a weather observation and optional payloads
func (s *weatherController) Create(_ context.Context, p *weather.Weather) (res *weather.Weather, err error) {
	res = &weather.Weather{}
	s.logger.Debug().Msg("weather.create")

	newWeather := models.Weather{}
	if p.WeatherTime != nil {
		newWeather.WeatherTime, err = time.Parse(time.RFC3339, *p.WeatherTime)
		if err != nil {
			err = weather.MakeBadRequest(errors.New("weatherTime must be un RFC3339 format"))
		}
	}
	if p.Temperature != nil {
		if p.Temperature.Unit != nil && p.Temperature.Value != nil {
			switch *p.Temperature.Unit {
			case "C", "F","K":
				// ok, get the appropriate tempUnit model
				var tempUnit models.TempUnit
				tempUnit, err = s.tempUnitService.GetByUnit(*p.Temperature.Unit)
				if err != nil {
					return
				}
				newWeather.TempUnit = tempUnit
				newWeather.TempUnitID = tempUnit.ID
			default:
				err = weather.MakeBadRequest(errors.New("invalid temperature type provided"))
				return
			}
			newWeather.Temp = uint(*p.Temperature.Value)
		}
	}
	if p.FireID != nil {
		fireID := uint(*p.FireID)
		newWeather.FireID = &fireID
	}
	if p.LogID != nil {
		logID := uint(*p.LogID)
		newWeather.LogID = &logID
	}
	if p.Humidity != nil {
		newWeather.Humidity = uint(*p.Humidity)
	}
	if p.DewPoint != nil {
		if p.DewPoint.Unit != nil && p.DewPoint.Value != nil {
			switch *p.DewPoint.Unit {
			case "C", "F", "K":
				// ok, get the appropriate tempUnit model
				var tempUnit models.TempUnit
				tempUnit, err = s.tempUnitService.GetByUnit(*p.DewPoint.Unit)
				if err != nil {
					return
				}
				newWeather.DewPointUnit = tempUnit
				newWeather.DewPointUnitID = tempUnit.ID
			default:
				err = weather.MakeBadRequest(errors.New("invalid temperature type provided"))
				return
			}
			newWeather.DewPoint = uint(*p.DewPoint.Value)
		}
	}
	if p.Wind != nil {
		if p.Wind.Unit != nil && p.Wind.Speed != nil && p.Wind.Direction != nil {
			switch *p.Wind.Unit {
			case "KPH", "MPH":
				// ok, get the appropriate speedUnit model
				var speedUnit models.SpeedUnit
				speedUnit, err = s.speedUnitService.GetByUnit(*p.Wind.Unit)
				if err != nil {
					return
				}
				newWeather.WindUnit = speedUnit
				newWeather.WindUnitID = speedUnit.ID
			default:
				err = weather.MakeBadRequest(errors.New("invalid speed type provided"))
				return
			}
			newWeather.WindSpeed = uint(*p.Wind.Speed)
			newWeather.WindDirection = *p.Wind.Direction
		}
	}
	if p.WeatherType != nil {
		newWeather.Type = *p.WeatherType
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
		if p.Temperature.Unit != nil && p.Temperature.Value != nil {
			switch *p.Temperature.Unit {
			case "C", "F","K":
				// ok, get the appropriate tempUnit model
				var tempUnit models.TempUnit
				tempUnit, err = s.tempUnitService.GetByUnit(*p.Temperature.Unit)
				if err != nil {
					return
				}
				tempWeather.TempUnit = tempUnit
				tempWeather.TempUnitID = tempUnit.ID
			default:
				err = weather.MakeBadRequest(errors.New("invalid temperature type provided"))
				return
			}
			tempWeather.Temp = uint(*p.Temperature.Value)
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
		if p.DewPoint.Unit != nil && p.DewPoint.Value != nil {
			switch *p.DewPoint.Unit {
			case "C", "F", "K":
				// ok, get the appropriate tempUnit model
				var tempUnit models.TempUnit
				tempUnit, err = s.tempUnitService.GetByUnit(*p.DewPoint.Unit)
				if err != nil {
					return
				}
				tempWeather.DewPointUnit = tempUnit
				tempWeather.DewPointUnitID = tempUnit.ID
			default:
				err = weather.MakeBadRequest(errors.New("invalid temperature type provided"))
				return
			}
			tempWeather.DewPoint = uint(*p.DewPoint.Value)
		}
	}
	if p.Wind != nil {
		if p.Wind.Unit != nil && p.Wind.Speed != nil && p.Wind.Direction != nil {
			switch *p.Wind.Unit {
			case "KPH", "MPH":
				// ok, get the appropriate speedUnit model
				var speedUnit models.SpeedUnit
				speedUnit, err = s.speedUnitService.GetByUnit(*p.Wind.Unit)
				if err != nil {
					return
				}
				tempWeather.WindUnit = speedUnit
				tempWeather.WindUnitID = speedUnit.ID
			default:
				err = weather.MakeBadRequest(errors.New("invalid speed type provided"))
				return
			}
			tempWeather.WindSpeed = uint(*p.Wind.Speed)
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
