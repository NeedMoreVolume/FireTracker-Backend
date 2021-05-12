package controllers

import (
	"github.com/NeedMoreVolume/FireTracker/gen/fire"
	"github.com/NeedMoreVolume/FireTracker/gen/log"
	"github.com/NeedMoreVolume/FireTracker/gen/weather"
	"github.com/NeedMoreVolume/FireTracker/models"
	"time"
)

// fire transports
func fireToTransport(model models.Fire) (out *fire.Fire) {
	id := int(model.ID)
	createdAt := model.CreatedAt.Format(time.RFC3339)
	updatedAt := model.UpdatedAt.Format(time.RFC3339)
	start := model.Start.Format(time.RFC3339)

	out = &fire.Fire{
		ID:          &id,
		CreatedAt:   &createdAt,
		UpdatedAt:   &updatedAt,
		Name:        &model.Name,
		Description: &model.Description,
		Start:       &start,
	}

	if model.End != nil {
		if !model.End.IsZero() {
			end := model.End.Format(time.RFC3339)
			out.End = &end
		}
	}

	if model.DeletedAt.Valid {
		deletedAt := model.DeletedAt.Time.Format(time.RFC3339)
		out.DeletedAt = &deletedAt
	}

	for _, thisWeather := range model.Weather {
		out.Weathers = append(out.Weathers, fireWeatherToTransport(thisWeather))
	}

	for _, thisLog := range model.Logs {
		out.Logs = append(out.Logs, fireLogToTransport(thisLog))
	}

	return
}

func fireWeatherToTransport(model models.Weather) (out *fire.Weather) {
	id := int(model.ID)
	createdAt := model.CreatedAt.Format(time.RFC3339)
	temp := int32(model.Temp)
	humidity := int32(model.Humidity)
	dewPoint := int32(model.DewPoint)
	windSpeed := int32(model.WindSpeed)
	wt := model.WeatherTime.Format(time.RFC3339)

	// units
	windUnit := "MPH"
	if model.WindUnit != "" {
		windUnit = model.WindUnit
	}
	dewUnit := "K"
	if model.DewPointUnit != "" {
		dewUnit = model.DewPointUnit
	}
	tempUnit := "K"
	if model.TempUnit != "" {
		tempUnit = model.TempUnit
	}

	out = &fire.Weather{
		ID:        &id,
		CreatedAt: &createdAt,

		Temperature: &fire.Temperature{
			Value: &temp,
			Unit:  &tempUnit,
		},
		Humidity: &humidity,
		DewPoint: &fire.Temperature{
			Value: &dewPoint,
			Unit:  &dewUnit,
		},
		Wind: &fire.Wind{
			Speed:     &windSpeed,
			Direction: &model.WindDirection,
			Unit:      &windUnit,
		},
		WeatherType: &model.Type,
		WeatherTime: &wt,
	}

	if model.FireID != nil {
		fireID := int(*model.FireID)
		out.FireID = &fireID
	}

	if model.LogID != nil {
		logID := int(*model.LogID)
		out.LogID = &logID
	}

	return
}

func fireLogToTransport(model models.Log) (out *fire.Log) {
	id := int(model.ID)
	createdAt := model.CreatedAt.Format(time.RFC3339)
	updatedAt := model.UpdatedAt.Format(time.RFC3339)
	fireID := int(model.FireID)
	addedAt := model.AddedAt.Format(time.RFC3339)

	out = &fire.Log{
		ID:        &id,
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
		Name:      &model.Name,
		Size:      &model.Size,
		FireID:    &fireID,
		AddedAt:   &addedAt,
	}

	if model.WeatherID != nil {
		out.Weather = fireWeatherToTransport(model.Weather)
	}

	return
}

// weather transports
func weatherToTransport(model models.Weather) (out *weather.Weather) {
	id := int(model.ID)
	createdAt := model.CreatedAt.Format(time.RFC3339)
	temp := int32(model.Temp)
	humidity := int32(model.Humidity)
	dewPoint := int32(model.DewPoint)
	windSpeed := int32(model.WindSpeed)
	wt := model.WeatherTime.Format(time.RFC3339)

	out = &weather.Weather{
		ID:        &id,
		CreatedAt: &createdAt,

		Temperature: &weather.Temperature{
			Value: &temp,
		},
		Humidity: &humidity,
		DewPoint: &weather.Temperature{
			Value: &dewPoint,
		},
		Wind: &weather.Wind{
			Speed:     &windSpeed,
			Direction: &model.WindDirection,
		},
		WeatherType: &model.Type,
		WeatherTime: &wt,
	}

	if model.FireID != nil {
		fireID := int(*model.FireID)
		out.FireID = &fireID
	}

	if model.LogID != nil {
		logID := int(*model.LogID)
		out.LogID = &logID
	}

	return
}

// log transports
func logWeatherToTransport(model models.Weather) (out *log.Weather) {
	id := int(model.ID)
	createdAt := model.CreatedAt.Format(time.RFC3339)
	temp := int32(model.Temp)
	humidity := int32(model.Humidity)
	dewPoint := int32(model.DewPoint)
	windSpeed := int32(model.WindSpeed)
	weatherTime := model.WeatherTime.Format(time.RFC3339)

	out = &log.Weather{
		ID:        &id,
		CreatedAt: &createdAt,
		Temperature: &log.Temperature{
			Unit:  &model.TempUnit,
			Value: &temp,
		},
		Humidity: &humidity,
		DewPoint: &log.Temperature{
			Unit:  &model.DewPointUnit,
			Value: &dewPoint,
		},
		Wind: &log.Wind{
			Speed:     &windSpeed,
			Direction: &model.WindDirection,
			Unit:      &model.WindUnit,
		},
		WeatherType: &model.Type,
		WeatherTime: &weatherTime,
	}

	if model.LogID != nil {
		lid := int(*model.LogID)
		out.LogID = &lid
	}
	if model.FireID != nil {
		fid := int(*model.FireID)
		out.FireID = &fid
	}

	return
}

func logToTransport(model models.Log) (out *log.Log) {
	id := int(model.ID)
	createdAt := model.CreatedAt.Format(time.RFC3339)
	updatedAt := model.UpdatedAt.Format(time.RFC3339)
	fireID := int(model.FireID)
	addedAt := model.AddedAt.Format(time.RFC3339)

	out = &log.Log{
		ID:        &id,
		CreatedAt: &createdAt,
		UpdatedAt: &updatedAt,
		Name:      &model.Name,
		Size:      &model.Size,
		FireID:    &fireID,
		AddedAt:   &addedAt,
	}
	if model.WeatherID != nil {
		out.Weather = logWeatherToTransport(model.Weather)
	}

	return
}
