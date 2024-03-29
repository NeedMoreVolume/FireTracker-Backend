// Code generated by goa v3.3.1, DO NOT EDIT.
//
// weather service
//
// Command:
// $ goa gen github.com/NeedMoreVolume/FireTracker/design

package weather

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// The fire service creates new weather datas for fires, and gets/lists weather
// data
type Service interface {
	// Create a weather observation and optional payloads
	Create(context.Context, *Weather) (res *Weather, err error)
	// Get a specific piece of weather data
	Get(context.Context, *Weather) (res *Weather, err error)
	// Update something about a weather observation specifically
	Update(context.Context, *Weather) (res *Weather, err error)
	// Delete a weather observation
	Delete(context.Context, *Weather) (err error)
	// List weather
	List(context.Context, *WeatherListPayload) (res *WeatherList, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "weather"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [5]string{"create", "get", "update", "delete", "list"}

// Weather is the payload type of the weather service create method.
type Weather struct {
	// id
	ID *int
	// name
	CreatedAt *string
	// ID of the fire
	FireID *int
	// ID of the log
	LogID *int
	// temperature
	Temperature *Temperature
	// humidity level
	Humidity *int32
	// dew point
	DewPoint *Temperature
	// wind data
	Wind *Wind
	// type of weather
	WeatherType *string
	WeatherTime *string
}

// WeatherListPayload is the payload type of the weather service list method.
type WeatherListPayload struct {
	// product filters to apply
	Filters *WeatherFilters
	// product search to apply
	Search *WeatherSearch
	// product sort to apply
	Sort *WeatherSorts
	// product pagination to apply
	Pagination *WeatherPagination
}

// WeatherList is the result type of the weather service list method.
type WeatherList struct {
	// weather results
	Weathers []*Weather
	// pagination info
	Pagination *Pagination
}

type Temperature struct {
	// measurement unit
	Unit *string
	// temperature value
	Value *int32
}

type Wind struct {
	// wind speed
	Speed *int32
	// wind direction
	Direction *string
	// measurement unit
	Unit *string
}

// list weather payload for weather
type WeatherFilters struct {
	Time        []*TimeFilter
	WindSpeed   []*NumberFilter
	Temperature []*NumberFilter
	Humidity    []*NumberFilter
}

type TimeFilter struct {
	// filter key
	Key string
	// operator value
	Operator string
	// filter value
	Value string
}

type NumberFilter struct {
	// filter key
	Key string
	// operator value
	Operator string
	// filter value
	Value int
}

// list search payload for products, with each field representing a valid
// search key
type WeatherSearch struct {
	Name        *string
	Description *string
}

// list sort payload for products, with each field representing a valid sort key
type WeatherSorts struct {
	ID *string
}

// list pagination for products
type WeatherPagination struct {
	Page  int
	Limit int
}

type Pagination struct {
	// count of the things
	Total *int
	// page number
	Page *int
	// max number of things
	Limit *int
}

// MakeBadRequest builds a goa.ServiceError from an error.
func MakeBadRequest(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "bad_request",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeNotFound builds a goa.ServiceError from an error.
func MakeNotFound(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "not_found",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeUnauthorized builds a goa.ServiceError from an error.
func MakeUnauthorized(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "unauthorized",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeForbidden builds a goa.ServiceError from an error.
func MakeForbidden(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "forbidden",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}
