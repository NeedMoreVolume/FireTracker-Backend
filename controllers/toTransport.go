package controllers

import (
	"github.com/NeedMoreVolume/FireTracker/gen/fire"
	"github.com/NeedMoreVolume/FireTracker/gen/log"
	"github.com/NeedMoreVolume/FireTracker/gen/weather"
	"github.com/NeedMoreVolume/FireTracker/models"
	"time"
)

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
	if model.WindUnit.Unit != "" {
		windUnit = model.WindUnit.Unit
	}
	dewUnit := "K"
	if model.DewPointUnit.Unit != "" {
		dewUnit = model.DewPointUnit.Unit
	}
	tempUnit := "K"
	if model.TempUnit.Unit != "" {
		tempUnit = model.TempUnit.Unit
	}

	out = &fire.Weather{
		ID:          &id,
		CreatedAt:   &createdAt,

		Temperature: &fire.Temperature{
			Value: &temp,
			Unit:  &tempUnit,
		},
		Humidity:    &humidity,
		DewPoint:    &fire.Temperature{
			Value: &dewPoint,
			Unit:  &dewUnit,
		},
		Wind:        &fire.Wind{
			Speed:     &windSpeed,
			Direction: &model.WindDirection,
			Unit: 	   &windUnit,
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

func weatherToTransport(model models.Weather) (out *weather.Weather) {
	id := int(model.ID)
	createdAt := model.CreatedAt.Format(time.RFC3339)
	temp := int32(model.Temp)
	humidity := int32(model.Humidity)
	dewPoint := int32(model.DewPoint)
	windSpeed := int32(model.WindSpeed)
	wt := model.WeatherTime.Format(time.RFC3339)

	out = &weather.Weather{
		ID:          &id,
		CreatedAt:   &createdAt,

		Temperature: &weather.Temperature{
			Value: &temp,
		},
		Humidity:    &humidity,
		DewPoint:    &weather.Temperature{
			Value: &dewPoint,
		},
		Wind:        &weather.Wind{
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
	return
}
