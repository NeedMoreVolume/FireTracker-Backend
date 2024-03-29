// Code generated by goa v3.3.1, DO NOT EDIT.
//
// weather HTTP server types
//
// Command:
// $ goa gen github.com/NeedMoreVolume/FireTracker/design

package server

import (
	weather "github.com/NeedMoreVolume/FireTracker/gen/weather"
	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "weather" service "create" endpoint
// HTTP request body.
type CreateRequestBody struct {
	// id
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// name
	CreatedAt *string `form:"createdAt,omitempty" json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// ID of the fire
	FireID *int `form:"fireID,omitempty" json:"fireID,omitempty" xml:"fireID,omitempty"`
	// ID of the log
	LogID *int `form:"logID,omitempty" json:"logID,omitempty" xml:"logID,omitempty"`
	// temperature
	Temperature *TemperatureRequestBody `form:"temperature,omitempty" json:"temperature,omitempty" xml:"temperature,omitempty"`
	// humidity level
	Humidity *int32 `form:"humidity,omitempty" json:"humidity,omitempty" xml:"humidity,omitempty"`
	// dew point
	DewPoint *TemperatureRequestBody `form:"dewPoint,omitempty" json:"dewPoint,omitempty" xml:"dewPoint,omitempty"`
	// wind data
	Wind *WindRequestBody `form:"wind,omitempty" json:"wind,omitempty" xml:"wind,omitempty"`
	// type of weather
	WeatherType *string `form:"weatherType,omitempty" json:"weatherType,omitempty" xml:"weatherType,omitempty"`
	WeatherTime *string `form:"weatherTime,omitempty" json:"weatherTime,omitempty" xml:"weatherTime,omitempty"`
}

// GetRequestBody is the type of the "weather" service "get" endpoint HTTP
// request body.
type GetRequestBody struct {
	// name
	CreatedAt *string `form:"createdAt,omitempty" json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// ID of the fire
	FireID *int `form:"fireID,omitempty" json:"fireID,omitempty" xml:"fireID,omitempty"`
	// ID of the log
	LogID *int `form:"logID,omitempty" json:"logID,omitempty" xml:"logID,omitempty"`
	// temperature
	Temperature *TemperatureRequestBody `form:"temperature,omitempty" json:"temperature,omitempty" xml:"temperature,omitempty"`
	// humidity level
	Humidity *int32 `form:"humidity,omitempty" json:"humidity,omitempty" xml:"humidity,omitempty"`
	// dew point
	DewPoint *TemperatureRequestBody `form:"dewPoint,omitempty" json:"dewPoint,omitempty" xml:"dewPoint,omitempty"`
	// wind data
	Wind *WindRequestBody `form:"wind,omitempty" json:"wind,omitempty" xml:"wind,omitempty"`
	// type of weather
	WeatherType *string `form:"weatherType,omitempty" json:"weatherType,omitempty" xml:"weatherType,omitempty"`
	WeatherTime *string `form:"weatherTime,omitempty" json:"weatherTime,omitempty" xml:"weatherTime,omitempty"`
}

// UpdateRequestBody is the type of the "weather" service "update" endpoint
// HTTP request body.
type UpdateRequestBody struct {
	// name
	CreatedAt *string `form:"createdAt,omitempty" json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// ID of the fire
	FireID *int `form:"fireID,omitempty" json:"fireID,omitempty" xml:"fireID,omitempty"`
	// ID of the log
	LogID *int `form:"logID,omitempty" json:"logID,omitempty" xml:"logID,omitempty"`
	// temperature
	Temperature *TemperatureRequestBody `form:"temperature,omitempty" json:"temperature,omitempty" xml:"temperature,omitempty"`
	// humidity level
	Humidity *int32 `form:"humidity,omitempty" json:"humidity,omitempty" xml:"humidity,omitempty"`
	// dew point
	DewPoint *TemperatureRequestBody `form:"dewPoint,omitempty" json:"dewPoint,omitempty" xml:"dewPoint,omitempty"`
	// wind data
	Wind *WindRequestBody `form:"wind,omitempty" json:"wind,omitempty" xml:"wind,omitempty"`
	// type of weather
	WeatherType *string `form:"weatherType,omitempty" json:"weatherType,omitempty" xml:"weatherType,omitempty"`
	WeatherTime *string `form:"weatherTime,omitempty" json:"weatherTime,omitempty" xml:"weatherTime,omitempty"`
}

// DeleteRequestBody is the type of the "weather" service "delete" endpoint
// HTTP request body.
type DeleteRequestBody struct {
	// name
	CreatedAt *string `form:"createdAt,omitempty" json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// ID of the fire
	FireID *int `form:"fireID,omitempty" json:"fireID,omitempty" xml:"fireID,omitempty"`
	// ID of the log
	LogID *int `form:"logID,omitempty" json:"logID,omitempty" xml:"logID,omitempty"`
	// temperature
	Temperature *TemperatureRequestBody `form:"temperature,omitempty" json:"temperature,omitempty" xml:"temperature,omitempty"`
	// humidity level
	Humidity *int32 `form:"humidity,omitempty" json:"humidity,omitempty" xml:"humidity,omitempty"`
	// dew point
	DewPoint *TemperatureRequestBody `form:"dewPoint,omitempty" json:"dewPoint,omitempty" xml:"dewPoint,omitempty"`
	// wind data
	Wind *WindRequestBody `form:"wind,omitempty" json:"wind,omitempty" xml:"wind,omitempty"`
	// type of weather
	WeatherType *string `form:"weatherType,omitempty" json:"weatherType,omitempty" xml:"weatherType,omitempty"`
	WeatherTime *string `form:"weatherTime,omitempty" json:"weatherTime,omitempty" xml:"weatherTime,omitempty"`
}

// ListRequestBody is the type of the "weather" service "list" endpoint HTTP
// request body.
type ListRequestBody struct {
	// product filters to apply
	Filters *WeatherFiltersRequestBody `form:"filters,omitempty" json:"filters,omitempty" xml:"filters,omitempty"`
	// product search to apply
	Search *WeatherSearchRequestBody `form:"search,omitempty" json:"search,omitempty" xml:"search,omitempty"`
	// product sort to apply
	Sort *WeatherSortsRequestBody `form:"sort,omitempty" json:"sort,omitempty" xml:"sort,omitempty"`
	// product pagination to apply
	Pagination *WeatherPaginationRequestBody `form:"pagination,omitempty" json:"pagination,omitempty" xml:"pagination,omitempty"`
}

// CreateOKResponseBody is the type of the "weather" service "create" endpoint
// HTTP response body.
type CreateOKResponseBody struct {
	// id
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// name
	CreatedAt *string `form:"createdAt,omitempty" json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// ID of the fire
	FireID *int `form:"fireID,omitempty" json:"fireID,omitempty" xml:"fireID,omitempty"`
	// ID of the log
	LogID *int `form:"logID,omitempty" json:"logID,omitempty" xml:"logID,omitempty"`
	// temperature
	Temperature *TemperatureResponseBody `form:"temperature,omitempty" json:"temperature,omitempty" xml:"temperature,omitempty"`
	// humidity level
	Humidity *int32 `form:"humidity,omitempty" json:"humidity,omitempty" xml:"humidity,omitempty"`
	// dew point
	DewPoint *TemperatureResponseBody `form:"dewPoint,omitempty" json:"dewPoint,omitempty" xml:"dewPoint,omitempty"`
	// wind data
	Wind *WindResponseBody `form:"wind,omitempty" json:"wind,omitempty" xml:"wind,omitempty"`
	// type of weather
	WeatherType *string `form:"weatherType,omitempty" json:"weatherType,omitempty" xml:"weatherType,omitempty"`
	WeatherTime *string `form:"weatherTime,omitempty" json:"weatherTime,omitempty" xml:"weatherTime,omitempty"`
}

// GetOKResponseBody is the type of the "weather" service "get" endpoint HTTP
// response body.
type GetOKResponseBody struct {
	// id
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// name
	CreatedAt *string `form:"createdAt,omitempty" json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// ID of the fire
	FireID *int `form:"fireID,omitempty" json:"fireID,omitempty" xml:"fireID,omitempty"`
	// ID of the log
	LogID *int `form:"logID,omitempty" json:"logID,omitempty" xml:"logID,omitempty"`
	// temperature
	Temperature *TemperatureResponseBody `form:"temperature,omitempty" json:"temperature,omitempty" xml:"temperature,omitempty"`
	// humidity level
	Humidity *int32 `form:"humidity,omitempty" json:"humidity,omitempty" xml:"humidity,omitempty"`
	// dew point
	DewPoint *TemperatureResponseBody `form:"dewPoint,omitempty" json:"dewPoint,omitempty" xml:"dewPoint,omitempty"`
	// wind data
	Wind *WindResponseBody `form:"wind,omitempty" json:"wind,omitempty" xml:"wind,omitempty"`
	// type of weather
	WeatherType *string `form:"weatherType,omitempty" json:"weatherType,omitempty" xml:"weatherType,omitempty"`
	WeatherTime *string `form:"weatherTime,omitempty" json:"weatherTime,omitempty" xml:"weatherTime,omitempty"`
}

// UpdateOKResponseBody is the type of the "weather" service "update" endpoint
// HTTP response body.
type UpdateOKResponseBody struct {
	// id
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// name
	CreatedAt *string `form:"createdAt,omitempty" json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// ID of the fire
	FireID *int `form:"fireID,omitempty" json:"fireID,omitempty" xml:"fireID,omitempty"`
	// ID of the log
	LogID *int `form:"logID,omitempty" json:"logID,omitempty" xml:"logID,omitempty"`
	// temperature
	Temperature *TemperatureResponseBody `form:"temperature,omitempty" json:"temperature,omitempty" xml:"temperature,omitempty"`
	// humidity level
	Humidity *int32 `form:"humidity,omitempty" json:"humidity,omitempty" xml:"humidity,omitempty"`
	// dew point
	DewPoint *TemperatureResponseBody `form:"dewPoint,omitempty" json:"dewPoint,omitempty" xml:"dewPoint,omitempty"`
	// wind data
	Wind *WindResponseBody `form:"wind,omitempty" json:"wind,omitempty" xml:"wind,omitempty"`
	// type of weather
	WeatherType *string `form:"weatherType,omitempty" json:"weatherType,omitempty" xml:"weatherType,omitempty"`
	WeatherTime *string `form:"weatherTime,omitempty" json:"weatherTime,omitempty" xml:"weatherTime,omitempty"`
}

// ListOKResponseBody is the type of the "weather" service "list" endpoint HTTP
// response body.
type ListOKResponseBody struct {
	// weather results
	Weathers []*WeatherResponseBody `form:"weathers,omitempty" json:"weathers,omitempty" xml:"weathers,omitempty"`
	// pagination info
	Pagination *PaginationResponseBody `form:"pagination,omitempty" json:"pagination,omitempty" xml:"pagination,omitempty"`
}

// CreateBadRequestResponseBody is the type of the "weather" service "create"
// endpoint HTTP response body for the "bad_request" error.
type CreateBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// GetNotFoundResponseBody is the type of the "weather" service "get" endpoint
// HTTP response body for the "not_found" error.
type GetNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateNotFoundResponseBody is the type of the "weather" service "update"
// endpoint HTTP response body for the "not_found" error.
type UpdateNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// DeleteNotFoundResponseBody is the type of the "weather" service "delete"
// endpoint HTTP response body for the "not_found" error.
type DeleteNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// TemperatureResponseBody is used to define fields on response body types.
type TemperatureResponseBody struct {
	// measurement unit
	Unit *string `form:"unit,omitempty" json:"unit,omitempty" xml:"unit,omitempty"`
	// temperature value
	Value *int32 `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// WindResponseBody is used to define fields on response body types.
type WindResponseBody struct {
	// wind speed
	Speed *int32 `form:"speed,omitempty" json:"speed,omitempty" xml:"speed,omitempty"`
	// wind direction
	Direction *string `form:"direction,omitempty" json:"direction,omitempty" xml:"direction,omitempty"`
	// measurement unit
	Unit *string `form:"unit,omitempty" json:"unit,omitempty" xml:"unit,omitempty"`
}

// WeatherResponseBody is used to define fields on response body types.
type WeatherResponseBody struct {
	// id
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// name
	CreatedAt *string `form:"createdAt,omitempty" json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// ID of the fire
	FireID *int `form:"fireID,omitempty" json:"fireID,omitempty" xml:"fireID,omitempty"`
	// ID of the log
	LogID *int `form:"logID,omitempty" json:"logID,omitempty" xml:"logID,omitempty"`
	// temperature
	Temperature *TemperatureResponseBody `form:"temperature,omitempty" json:"temperature,omitempty" xml:"temperature,omitempty"`
	// humidity level
	Humidity *int32 `form:"humidity,omitempty" json:"humidity,omitempty" xml:"humidity,omitempty"`
	// dew point
	DewPoint *TemperatureResponseBody `form:"dewPoint,omitempty" json:"dewPoint,omitempty" xml:"dewPoint,omitempty"`
	// wind data
	Wind *WindResponseBody `form:"wind,omitempty" json:"wind,omitempty" xml:"wind,omitempty"`
	// type of weather
	WeatherType *string `form:"weatherType,omitempty" json:"weatherType,omitempty" xml:"weatherType,omitempty"`
	WeatherTime *string `form:"weatherTime,omitempty" json:"weatherTime,omitempty" xml:"weatherTime,omitempty"`
}

// PaginationResponseBody is used to define fields on response body types.
type PaginationResponseBody struct {
	// count of the things
	Total *int `form:"total,omitempty" json:"total,omitempty" xml:"total,omitempty"`
	// page number
	Page *int `form:"page,omitempty" json:"page,omitempty" xml:"page,omitempty"`
	// max number of things
	Limit *int `form:"limit,omitempty" json:"limit,omitempty" xml:"limit,omitempty"`
}

// TemperatureRequestBody is used to define fields on request body types.
type TemperatureRequestBody struct {
	// measurement unit
	Unit *string `form:"unit,omitempty" json:"unit,omitempty" xml:"unit,omitempty"`
	// temperature value
	Value *int32 `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// WindRequestBody is used to define fields on request body types.
type WindRequestBody struct {
	// wind speed
	Speed *int32 `form:"speed,omitempty" json:"speed,omitempty" xml:"speed,omitempty"`
	// wind direction
	Direction *string `form:"direction,omitempty" json:"direction,omitempty" xml:"direction,omitempty"`
	// measurement unit
	Unit *string `form:"unit,omitempty" json:"unit,omitempty" xml:"unit,omitempty"`
}

// WeatherFiltersRequestBody is used to define fields on request body types.
type WeatherFiltersRequestBody struct {
	Time        []*TimeFilterRequestBody   `form:"time,omitempty" json:"time,omitempty" xml:"time,omitempty"`
	WindSpeed   []*NumberFilterRequestBody `form:"windSpeed,omitempty" json:"windSpeed,omitempty" xml:"windSpeed,omitempty"`
	Temperature []*NumberFilterRequestBody `form:"temperature,omitempty" json:"temperature,omitempty" xml:"temperature,omitempty"`
	Humidity    []*NumberFilterRequestBody `form:"humidity,omitempty" json:"humidity,omitempty" xml:"humidity,omitempty"`
}

// TimeFilterRequestBody is used to define fields on request body types.
type TimeFilterRequestBody struct {
	// filter key
	Key *string `form:"key,omitempty" json:"key,omitempty" xml:"key,omitempty"`
	// operator value
	Operator *string `form:"operator,omitempty" json:"operator,omitempty" xml:"operator,omitempty"`
	// filter value
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// NumberFilterRequestBody is used to define fields on request body types.
type NumberFilterRequestBody struct {
	// filter key
	Key *string `form:"key,omitempty" json:"key,omitempty" xml:"key,omitempty"`
	// operator value
	Operator *string `form:"operator,omitempty" json:"operator,omitempty" xml:"operator,omitempty"`
	// filter value
	Value *int `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// WeatherSearchRequestBody is used to define fields on request body types.
type WeatherSearchRequestBody struct {
	Name        *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
}

// WeatherSortsRequestBody is used to define fields on request body types.
type WeatherSortsRequestBody struct {
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
}

// WeatherPaginationRequestBody is used to define fields on request body types.
type WeatherPaginationRequestBody struct {
	Page  *int `form:"page,omitempty" json:"page,omitempty" xml:"page,omitempty"`
	Limit *int `form:"limit,omitempty" json:"limit,omitempty" xml:"limit,omitempty"`
}

// NewCreateOKResponseBody builds the HTTP response body from the result of the
// "create" endpoint of the "weather" service.
func NewCreateOKResponseBody(res *weather.Weather) *CreateOKResponseBody {
	body := &CreateOKResponseBody{
		ID:          res.ID,
		CreatedAt:   res.CreatedAt,
		FireID:      res.FireID,
		LogID:       res.LogID,
		Humidity:    res.Humidity,
		WeatherType: res.WeatherType,
		WeatherTime: res.WeatherTime,
	}
	if res.Temperature != nil {
		body.Temperature = marshalWeatherTemperatureToTemperatureResponseBody(res.Temperature)
	}
	if res.DewPoint != nil {
		body.DewPoint = marshalWeatherTemperatureToTemperatureResponseBody(res.DewPoint)
	}
	if res.Wind != nil {
		body.Wind = marshalWeatherWindToWindResponseBody(res.Wind)
	}
	return body
}

// NewGetOKResponseBody builds the HTTP response body from the result of the
// "get" endpoint of the "weather" service.
func NewGetOKResponseBody(res *weather.Weather) *GetOKResponseBody {
	body := &GetOKResponseBody{
		ID:          res.ID,
		CreatedAt:   res.CreatedAt,
		FireID:      res.FireID,
		LogID:       res.LogID,
		Humidity:    res.Humidity,
		WeatherType: res.WeatherType,
		WeatherTime: res.WeatherTime,
	}
	if res.Temperature != nil {
		body.Temperature = marshalWeatherTemperatureToTemperatureResponseBody(res.Temperature)
	}
	if res.DewPoint != nil {
		body.DewPoint = marshalWeatherTemperatureToTemperatureResponseBody(res.DewPoint)
	}
	if res.Wind != nil {
		body.Wind = marshalWeatherWindToWindResponseBody(res.Wind)
	}
	return body
}

// NewUpdateOKResponseBody builds the HTTP response body from the result of the
// "update" endpoint of the "weather" service.
func NewUpdateOKResponseBody(res *weather.Weather) *UpdateOKResponseBody {
	body := &UpdateOKResponseBody{
		ID:          res.ID,
		CreatedAt:   res.CreatedAt,
		FireID:      res.FireID,
		LogID:       res.LogID,
		Humidity:    res.Humidity,
		WeatherType: res.WeatherType,
		WeatherTime: res.WeatherTime,
	}
	if res.Temperature != nil {
		body.Temperature = marshalWeatherTemperatureToTemperatureResponseBody(res.Temperature)
	}
	if res.DewPoint != nil {
		body.DewPoint = marshalWeatherTemperatureToTemperatureResponseBody(res.DewPoint)
	}
	if res.Wind != nil {
		body.Wind = marshalWeatherWindToWindResponseBody(res.Wind)
	}
	return body
}

// NewListOKResponseBody builds the HTTP response body from the result of the
// "list" endpoint of the "weather" service.
func NewListOKResponseBody(res *weather.WeatherList) *ListOKResponseBody {
	body := &ListOKResponseBody{}
	if res.Weathers != nil {
		body.Weathers = make([]*WeatherResponseBody, len(res.Weathers))
		for i, val := range res.Weathers {
			body.Weathers[i] = marshalWeatherWeatherToWeatherResponseBody(val)
		}
	}
	if res.Pagination != nil {
		body.Pagination = marshalWeatherPaginationToPaginationResponseBody(res.Pagination)
	}
	return body
}

// NewCreateBadRequestResponseBody builds the HTTP response body from the
// result of the "create" endpoint of the "weather" service.
func NewCreateBadRequestResponseBody(res *goa.ServiceError) *CreateBadRequestResponseBody {
	body := &CreateBadRequestResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetNotFoundResponseBody builds the HTTP response body from the result of
// the "get" endpoint of the "weather" service.
func NewGetNotFoundResponseBody(res *goa.ServiceError) *GetNotFoundResponseBody {
	body := &GetNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateNotFoundResponseBody builds the HTTP response body from the result
// of the "update" endpoint of the "weather" service.
func NewUpdateNotFoundResponseBody(res *goa.ServiceError) *UpdateNotFoundResponseBody {
	body := &UpdateNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewDeleteNotFoundResponseBody builds the HTTP response body from the result
// of the "delete" endpoint of the "weather" service.
func NewDeleteNotFoundResponseBody(res *goa.ServiceError) *DeleteNotFoundResponseBody {
	body := &DeleteNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateWeather builds a weather service create endpoint payload.
func NewCreateWeather(body *CreateRequestBody) *weather.Weather {
	v := &weather.Weather{
		ID:          body.ID,
		CreatedAt:   body.CreatedAt,
		FireID:      body.FireID,
		LogID:       body.LogID,
		Humidity:    body.Humidity,
		WeatherType: body.WeatherType,
		WeatherTime: body.WeatherTime,
	}
	if body.Temperature != nil {
		v.Temperature = unmarshalTemperatureRequestBodyToWeatherTemperature(body.Temperature)
	}
	if body.DewPoint != nil {
		v.DewPoint = unmarshalTemperatureRequestBodyToWeatherTemperature(body.DewPoint)
	}
	if body.Wind != nil {
		v.Wind = unmarshalWindRequestBodyToWeatherWind(body.Wind)
	}

	return v
}

// NewGetWeather builds a weather service get endpoint payload.
func NewGetWeather(body *GetRequestBody, id int) *weather.Weather {
	v := &weather.Weather{
		CreatedAt:   body.CreatedAt,
		FireID:      body.FireID,
		LogID:       body.LogID,
		Humidity:    body.Humidity,
		WeatherType: body.WeatherType,
		WeatherTime: body.WeatherTime,
	}
	if body.Temperature != nil {
		v.Temperature = unmarshalTemperatureRequestBodyToWeatherTemperature(body.Temperature)
	}
	if body.DewPoint != nil {
		v.DewPoint = unmarshalTemperatureRequestBodyToWeatherTemperature(body.DewPoint)
	}
	if body.Wind != nil {
		v.Wind = unmarshalWindRequestBodyToWeatherWind(body.Wind)
	}
	v.ID = &id

	return v
}

// NewUpdateWeather builds a weather service update endpoint payload.
func NewUpdateWeather(body *UpdateRequestBody, id int) *weather.Weather {
	v := &weather.Weather{
		CreatedAt:   body.CreatedAt,
		FireID:      body.FireID,
		LogID:       body.LogID,
		Humidity:    body.Humidity,
		WeatherType: body.WeatherType,
		WeatherTime: body.WeatherTime,
	}
	if body.Temperature != nil {
		v.Temperature = unmarshalTemperatureRequestBodyToWeatherTemperature(body.Temperature)
	}
	if body.DewPoint != nil {
		v.DewPoint = unmarshalTemperatureRequestBodyToWeatherTemperature(body.DewPoint)
	}
	if body.Wind != nil {
		v.Wind = unmarshalWindRequestBodyToWeatherWind(body.Wind)
	}
	v.ID = &id

	return v
}

// NewDeleteWeather builds a weather service delete endpoint payload.
func NewDeleteWeather(body *DeleteRequestBody, id int) *weather.Weather {
	v := &weather.Weather{
		CreatedAt:   body.CreatedAt,
		FireID:      body.FireID,
		LogID:       body.LogID,
		Humidity:    body.Humidity,
		WeatherType: body.WeatherType,
		WeatherTime: body.WeatherTime,
	}
	if body.Temperature != nil {
		v.Temperature = unmarshalTemperatureRequestBodyToWeatherTemperature(body.Temperature)
	}
	if body.DewPoint != nil {
		v.DewPoint = unmarshalTemperatureRequestBodyToWeatherTemperature(body.DewPoint)
	}
	if body.Wind != nil {
		v.Wind = unmarshalWindRequestBodyToWeatherWind(body.Wind)
	}
	v.ID = &id

	return v
}

// NewListWeatherListPayload builds a weather service list endpoint payload.
func NewListWeatherListPayload(body *ListRequestBody) *weather.WeatherListPayload {
	v := &weather.WeatherListPayload{}
	v.Filters = unmarshalWeatherFiltersRequestBodyToWeatherWeatherFilters(body.Filters)
	v.Search = unmarshalWeatherSearchRequestBodyToWeatherWeatherSearch(body.Search)
	v.Sort = unmarshalWeatherSortsRequestBodyToWeatherWeatherSorts(body.Sort)
	v.Pagination = unmarshalWeatherPaginationRequestBodyToWeatherWeatherPagination(body.Pagination)

	return v
}

// ValidateCreateRequestBody runs the validations defined on CreateRequestBody
func ValidateCreateRequestBody(body *CreateRequestBody) (err error) {
	if body.CreatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.createdAt", *body.CreatedAt, goa.FormatDateTime))
	}
	if body.Temperature != nil {
		if err2 := ValidateTemperatureRequestBody(body.Temperature); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.DewPoint != nil {
		if err2 := ValidateTemperatureRequestBody(body.DewPoint); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.Wind != nil {
		if err2 := ValidateWindRequestBody(body.Wind); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.WeatherType != nil {
		if !(*body.WeatherType == "Sunny" || *body.WeatherType == "Cloudy" || *body.WeatherType == "Raining" || *body.WeatherType == "Windy") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.weatherType", *body.WeatherType, []interface{}{"Sunny", "Cloudy", "Raining", "Windy"}))
		}
	}
	if body.WeatherTime != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.weatherTime", *body.WeatherTime, goa.FormatDateTime))
	}
	return
}

// ValidateGetRequestBody runs the validations defined on GetRequestBody
func ValidateGetRequestBody(body *GetRequestBody) (err error) {
	if body.CreatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.createdAt", *body.CreatedAt, goa.FormatDateTime))
	}
	if body.Temperature != nil {
		if err2 := ValidateTemperatureRequestBody(body.Temperature); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.DewPoint != nil {
		if err2 := ValidateTemperatureRequestBody(body.DewPoint); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.Wind != nil {
		if err2 := ValidateWindRequestBody(body.Wind); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.WeatherType != nil {
		if !(*body.WeatherType == "Sunny" || *body.WeatherType == "Cloudy" || *body.WeatherType == "Raining" || *body.WeatherType == "Windy") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.weatherType", *body.WeatherType, []interface{}{"Sunny", "Cloudy", "Raining", "Windy"}))
		}
	}
	if body.WeatherTime != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.weatherTime", *body.WeatherTime, goa.FormatDateTime))
	}
	return
}

// ValidateUpdateRequestBody runs the validations defined on UpdateRequestBody
func ValidateUpdateRequestBody(body *UpdateRequestBody) (err error) {
	if body.CreatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.createdAt", *body.CreatedAt, goa.FormatDateTime))
	}
	if body.Temperature != nil {
		if err2 := ValidateTemperatureRequestBody(body.Temperature); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.DewPoint != nil {
		if err2 := ValidateTemperatureRequestBody(body.DewPoint); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.Wind != nil {
		if err2 := ValidateWindRequestBody(body.Wind); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.WeatherType != nil {
		if !(*body.WeatherType == "Sunny" || *body.WeatherType == "Cloudy" || *body.WeatherType == "Raining" || *body.WeatherType == "Windy") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.weatherType", *body.WeatherType, []interface{}{"Sunny", "Cloudy", "Raining", "Windy"}))
		}
	}
	if body.WeatherTime != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.weatherTime", *body.WeatherTime, goa.FormatDateTime))
	}
	return
}

// ValidateDeleteRequestBody runs the validations defined on DeleteRequestBody
func ValidateDeleteRequestBody(body *DeleteRequestBody) (err error) {
	if body.CreatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.createdAt", *body.CreatedAt, goa.FormatDateTime))
	}
	if body.Temperature != nil {
		if err2 := ValidateTemperatureRequestBody(body.Temperature); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.DewPoint != nil {
		if err2 := ValidateTemperatureRequestBody(body.DewPoint); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.Wind != nil {
		if err2 := ValidateWindRequestBody(body.Wind); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.WeatherType != nil {
		if !(*body.WeatherType == "Sunny" || *body.WeatherType == "Cloudy" || *body.WeatherType == "Raining" || *body.WeatherType == "Windy") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.weatherType", *body.WeatherType, []interface{}{"Sunny", "Cloudy", "Raining", "Windy"}))
		}
	}
	if body.WeatherTime != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.weatherTime", *body.WeatherTime, goa.FormatDateTime))
	}
	return
}

// ValidateListRequestBody runs the validations defined on ListRequestBody
func ValidateListRequestBody(body *ListRequestBody) (err error) {
	if body.Filters == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("filters", "body"))
	}
	if body.Search == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("search", "body"))
	}
	if body.Sort == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("sort", "body"))
	}
	if body.Pagination == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("pagination", "body"))
	}
	if body.Filters != nil {
		if err2 := ValidateWeatherFiltersRequestBody(body.Filters); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.Sort != nil {
		if err2 := ValidateWeatherSortsRequestBody(body.Sort); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.Pagination != nil {
		if err2 := ValidateWeatherPaginationRequestBody(body.Pagination); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateTemperatureRequestBody runs the validations defined on
// TemperatureRequestBody
func ValidateTemperatureRequestBody(body *TemperatureRequestBody) (err error) {
	if body.Unit != nil {
		if !(*body.Unit == "K" || *body.Unit == "C" || *body.Unit == "F") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.unit", *body.Unit, []interface{}{"K", "C", "F"}))
		}
	}
	return
}

// ValidateWindRequestBody runs the validations defined on WindRequestBody
func ValidateWindRequestBody(body *WindRequestBody) (err error) {
	if body.Direction != nil {
		if !(*body.Direction == "S" || *body.Direction == "SE" || *body.Direction == "E" || *body.Direction == "NE" || *body.Direction == "N" || *body.Direction == "NW" || *body.Direction == "W" || *body.Direction == "SW") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.direction", *body.Direction, []interface{}{"S", "SE", "E", "NE", "N", "NW", "W", "SW"}))
		}
	}
	if body.Unit != nil {
		if !(*body.Unit == "KPH" || *body.Unit == "MPH") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.unit", *body.Unit, []interface{}{"KPH", "MPH"}))
		}
	}
	return
}

// ValidateWeatherFiltersRequestBody runs the validations defined on
// WeatherFiltersRequestBody
func ValidateWeatherFiltersRequestBody(body *WeatherFiltersRequestBody) (err error) {
	for _, e := range body.Time {
		if e != nil {
			if err2 := ValidateTimeFilterRequestBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	for _, e := range body.WindSpeed {
		if e != nil {
			if err2 := ValidateNumberFilterRequestBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	for _, e := range body.Temperature {
		if e != nil {
			if err2 := ValidateNumberFilterRequestBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	for _, e := range body.Humidity {
		if e != nil {
			if err2 := ValidateNumberFilterRequestBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateTimeFilterRequestBody runs the validations defined on
// TimeFilterRequestBody
func ValidateTimeFilterRequestBody(body *TimeFilterRequestBody) (err error) {
	if body.Key == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("key", "body"))
	}
	if body.Operator == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("operator", "body"))
	}
	if body.Value == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("value", "body"))
	}
	if body.Operator != nil {
		if !(*body.Operator == "=" || *body.Operator == ">" || *body.Operator == ">=" || *body.Operator == "<=" || *body.Operator == "<") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.operator", *body.Operator, []interface{}{"=", ">", ">=", "<=", "<"}))
		}
	}
	if body.Value != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.value", *body.Value, goa.FormatDateTime))
	}
	return
}

// ValidateNumberFilterRequestBody runs the validations defined on
// NumberFilterRequestBody
func ValidateNumberFilterRequestBody(body *NumberFilterRequestBody) (err error) {
	if body.Key == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("key", "body"))
	}
	if body.Operator == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("operator", "body"))
	}
	if body.Value == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("value", "body"))
	}
	if body.Operator != nil {
		if !(*body.Operator == "=" || *body.Operator == ">" || *body.Operator == ">=" || *body.Operator == "<=" || *body.Operator == "<") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.operator", *body.Operator, []interface{}{"=", ">", ">=", "<=", "<"}))
		}
	}
	return
}

// ValidateWeatherSortsRequestBody runs the validations defined on
// WeatherSortsRequestBody
func ValidateWeatherSortsRequestBody(body *WeatherSortsRequestBody) (err error) {
	if body.ID != nil {
		if !(*body.ID == "ASC, DESC") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.id", *body.ID, []interface{}{"ASC, DESC"}))
		}
	}
	return
}

// ValidateWeatherPaginationRequestBody runs the validations defined on
// WeatherPaginationRequestBody
func ValidateWeatherPaginationRequestBody(body *WeatherPaginationRequestBody) (err error) {
	if body.Page == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("page", "body"))
	}
	if body.Limit == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("limit", "body"))
	}
	return
}
