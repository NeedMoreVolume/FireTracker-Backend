package controllers

import (
	"errors"
	"github.com/NeedMoreVolume/FireTracker/gen/fire"
	"github.com/NeedMoreVolume/FireTracker/gen/log"
	"github.com/NeedMoreVolume/FireTracker/gen/weather"
	"github.com/NeedMoreVolume/FireTracker/models"
	"time"
)

// fire model
func fireToModel(payload *fire.Fire) (out models.Fire, err error) {
	if payload.ID != nil {
		out.ID = uint(*payload.ID)
	}
	if payload.CreatedAt != nil {
		out.CreatedAt, err = time.Parse(time.RFC3339, *payload.CreatedAt)
		if err != nil {
			err = errors.New("invalid time format for created_at")
			return
		}
	}
	if payload.UpdatedAt != nil {
		out.UpdatedAt, err = time.Parse(time.RFC3339, *payload.UpdatedAt)
		if err != nil {
			err = errors.New("invalid time format for updated_at")
			return
		}
	}
	if payload.Name != nil {
		out.Name = *payload.Name
	}
	if payload.Description != nil {
		out.Description = *payload.Description
	}
	if payload.Start != nil {
		out.Start, err = time.Parse(time.RFC3339, *payload.Start)
		if err != nil {
			err = errors.New("invalid time format for start")
			return
		}
	}

	if payload.End != nil {
		var end time.Time
		end, err = time.Parse(time.RFC3339, *payload.End)
		if err != nil {
			err = errors.New("invalid time format for end")
			return
		}
		out.End = &end
	}

	return
}

func fireWeatherToModel(payload *fire.Weather) (out models.Weather, err error) {
	if payload.ID != nil {
		out.ID = uint(*payload.ID)
	}
	if payload.CreatedAt != nil {
		out.CreatedAt, err = time.Parse(time.RFC3339, *payload.CreatedAt)
		if err != nil {
			err = errors.New("invalid time format for created_at")
			return
		}
	}
	if payload.Temperature != nil {
		if payload.Temperature.Unit != nil {
			out.TempUnit = *payload.Temperature.Unit
		}
		if payload.Temperature.Value != nil {
			out.Temp = int64(*payload.Temperature.Value)
		}
	}
	if payload.Humidity != nil {
		out.Humidity = uint(*payload.Humidity)
	}
	if payload.DewPoint != nil {
		if payload.DewPoint.Value != nil {
			out.DewPoint = int64(*payload.DewPoint.Value)
		}
		if payload.DewPoint.Unit != nil {
			out.DewPointUnit = *payload.DewPoint.Unit
		}
	}
	if payload.Wind != nil {
		if payload.Wind.Unit != nil {
			out.WindUnit = *payload.Wind.Unit
		}
		if payload.Wind.Speed != nil {
			out.WindSpeed = uint(*payload.Wind.Speed)
		}
		if payload.Wind.Direction != nil {
			out.WindDirection = *payload.Wind.Direction
		}
	}
	if payload.WeatherTime != nil {
		out.WeatherTime, err = time.Parse(time.RFC3339, *payload.WeatherTime)
		if err != nil {
			err = errors.New("invalid time format for weather_time")
			return
		}
	}
	if payload.FireID != nil {
		fireID := uint(*payload.FireID)
		out.FireID = &fireID
	}
	if payload.LogID != nil {
		logID := uint(*payload.LogID)
		out.LogID = &logID
	}

	return
}

func fireLogToModel(payload *fire.Log) (out models.Log, err error) {
	if payload.ID != nil {
		out.ID = uint(*payload.ID)
	}
	if payload.CreatedAt != nil {
		out.CreatedAt, err = time.Parse(time.RFC3339, *payload.CreatedAt)
		if err != nil {
			err = errors.New("invalid time format for created_at")
			return
		}
	}
	if payload.UpdatedAt != nil {
		out.UpdatedAt, err = time.Parse(time.RFC3339, *payload.UpdatedAt)
		if err != nil {
			err = errors.New("invalid time format for updated_at")
			return
		}
	}
	if payload.FireID != nil {
		out.FireID = uint(*payload.FireID)
	}
	if payload.AddedAt != nil {
		out.AddedAt, err = time.Parse(time.RFC3339, *payload.AddedAt)
		if err != nil {
			err = errors.New("invalid time format for added_at")
			return
		}
	}
	if payload.Name != nil {
		out.Name = *payload.Name
	}
	if payload.Size != nil {
		out.Size = *payload.Size
	}
	if payload.Weather != nil {
		out.Weather, err = fireWeatherToModel(payload.Weather)
		if err != nil {
			return
		}
		out.WeatherID = &out.Weather.ID
	}

	return
}

// weather transports
func weatherToModel(payload *weather.Weather) (out models.Weather, err error) {
	if payload.ID != nil {
		out.ID = uint(*payload.ID)
	}
	if payload.CreatedAt != nil {
		out.CreatedAt, err = time.Parse(time.RFC3339, *payload.CreatedAt)
		if err != nil {
			err = errors.New("invalid time format for created_at")
			return
		}
	}
	if payload.Temperature != nil {
		if payload.Temperature.Unit != nil {
			out.TempUnit = *payload.Temperature.Unit
		}
		if payload.Temperature.Value != nil {
			out.Temp = int64(*payload.Temperature.Value)
		}
	}
	if payload.Humidity != nil {
		out.Humidity = uint(*payload.Humidity)
	}
	if payload.DewPoint != nil {
		if payload.DewPoint.Value != nil {
			out.DewPoint = int64(*payload.DewPoint.Value)
		}
		if payload.DewPoint.Unit != nil {
			out.DewPointUnit = *payload.DewPoint.Unit
		}
	}
	if payload.Wind != nil {
		if payload.Wind.Unit != nil {
			out.WindUnit = *payload.Wind.Unit
		}
		if payload.Wind.Direction != nil {
			out.WindDirection = *payload.Wind.Direction
		}
		if payload.Wind.Speed != nil {
			out.WindSpeed = uint(*payload.Wind.Speed)
		}
	}
	if payload.WeatherType != nil {
		out.Type = *payload.WeatherType
	}
	if payload.WeatherTime != nil {
		out.WeatherTime, err = time.Parse(time.RFC3339, *payload.WeatherTime)
		if err != nil {
			err = errors.New("invalid time format for weather_time")
			return
		}
	}
	if payload.FireID != nil {
		fireID := uint(*payload.FireID)
		out.FireID = &fireID
	}
	if payload.LogID != nil {
		logID := uint(*payload.LogID)
		out.LogID = &logID
	}

	return
}

// log transports
func logWeatherToModel(payload *log.Weather) (out models.Weather, err error) {
	if payload.ID != nil {
		out.ID = uint(*payload.ID)
	}
	if payload.LogID != nil {
		lid := uint(*payload.LogID)
		out.LogID = &lid
	}
	if payload.FireID != nil {
		fid := uint(*payload.FireID)
		out.FireID = &fid
	}
	if payload.CreatedAt != nil {
		out.CreatedAt, err = time.Parse(time.RFC3339, *payload.CreatedAt)
		if err != nil {
			err = errors.New("invalid time format for created_at")
			return
		}
	}
	if payload.Temperature != nil {
		if payload.Temperature.Unit != nil {
			out.TempUnit = *payload.Temperature.Unit
		}
		if payload.Temperature.Value != nil {
			out.Temp = int64(*payload.Temperature.Value)
		}
	}
	if payload.Humidity != nil {
		out.Humidity = uint(*payload.Humidity)
	}
	if payload.DewPoint != nil {
		if payload.DewPoint.Unit != nil {
			out.DewPointUnit = *payload.DewPoint.Unit
		}
		if payload.DewPoint.Value != nil {
			out.DewPoint = int64(*payload.DewPoint.Value)
		}
	}
	if payload.Wind != nil {
		if payload.Wind.Unit != nil {
			out.WindUnit = *payload.Wind.Unit
		}
		if payload.Wind.Speed != nil {
			out.WindSpeed = uint(*payload.Wind.Speed)
		}
		if payload.Wind.Direction != nil {
			out.WindDirection = *payload.Wind.Direction
		}
	}
	if payload.WeatherType != nil {
		out.Type = *payload.WeatherType
	}
	if payload.WeatherTime != nil {
		out.WeatherTime, err = time.Parse(time.RFC3339, *payload.WeatherTime)
		if err != nil {
			err = errors.New("invalid time format for weather_time")
			return
		}
	}

	return
}

func logToModel(payload *log.Log) (out models.Log, err error) {
	if payload.ID != nil {
		out.ID = uint(*payload.ID)
	}
	if payload.CreatedAt != nil {
		out.CreatedAt, err = time.Parse(time.RFC3339, *payload.CreatedAt)
		if err != nil {
			err = errors.New("invalid time format for created_at")
			return
		}
	}
	if payload.UpdatedAt != nil {
		out.UpdatedAt, err = time.Parse(time.RFC3339, *payload.UpdatedAt)
		if err != nil {
			err = errors.New("invalid time format for updated_at")
			return
		}
	}
	if payload.FireID != nil {
		out.FireID = uint(*payload.FireID)
	}
	if payload.AddedAt != nil {
		out.AddedAt, err = time.Parse(time.RFC3339, *payload.AddedAt)
		if err != nil {
			err = errors.New("invalid time format for added_at")
			return
		}
	}
	if payload.Name != nil {
		out.Name = *payload.Name
	}
	if payload.Size != nil {
		out.Size = *payload.Size
	}
	if payload.Weather != nil {
		out.Weather, err = logWeatherToModel(payload.Weather)
		if err != nil {
			return
		}
		out.WeatherID = &out.Weather.ID
	}

	return
}
