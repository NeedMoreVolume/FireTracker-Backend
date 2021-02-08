// Code generated by goa v3.2.5, DO NOT EDIT.
//
// log HTTP server types
//
// Command:
// $ goa gen github.com/NeedMoreVolume/FireTracker/design

package server

import (
	log "github.com/NeedMoreVolume/FireTracker/gen/log"
	goa "goa.design/goa/v3/pkg"
)

// CreateRequestBody is the type of the "log" service "create" endpoint HTTP
// request body.
type CreateRequestBody struct {
	// id
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// name
	CreatedAt *string `form:"createdAt,omitempty" json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// name
	UpdatedAt *string `form:"updatedAt,omitempty" json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// size of log
	Size *string `form:"size,omitempty" json:"size,omitempty" xml:"size,omitempty"`
	// Fire identifier log belongs to
	FireID *int `form:"fireID,omitempty" json:"fireID,omitempty" xml:"fireID,omitempty"`
	// time log was added to fire
	AddedAt *string `form:"addedAt,omitempty" json:"addedAt,omitempty" xml:"addedAt,omitempty"`
	// weather data at time log was added to fire
	Weather *WeatherRequestBody `form:"weather,omitempty" json:"weather,omitempty" xml:"weather,omitempty"`
}

// GetRequestBody is the type of the "log" service "get" endpoint HTTP request
// body.
type GetRequestBody struct {
	// name
	CreatedAt *string `form:"createdAt,omitempty" json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// name
	UpdatedAt *string `form:"updatedAt,omitempty" json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// size of log
	Size *string `form:"size,omitempty" json:"size,omitempty" xml:"size,omitempty"`
	// Fire identifier log belongs to
	FireID *int `form:"fireID,omitempty" json:"fireID,omitempty" xml:"fireID,omitempty"`
	// time log was added to fire
	AddedAt *string `form:"addedAt,omitempty" json:"addedAt,omitempty" xml:"addedAt,omitempty"`
	// weather data at time log was added to fire
	Weather *WeatherRequestBody `form:"weather,omitempty" json:"weather,omitempty" xml:"weather,omitempty"`
}

// UpdateRequestBody is the type of the "log" service "update" endpoint HTTP
// request body.
type UpdateRequestBody struct {
	// name
	CreatedAt *string `form:"createdAt,omitempty" json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// name
	UpdatedAt *string `form:"updatedAt,omitempty" json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// size of log
	Size *string `form:"size,omitempty" json:"size,omitempty" xml:"size,omitempty"`
	// Fire identifier log belongs to
	FireID *int `form:"fireID,omitempty" json:"fireID,omitempty" xml:"fireID,omitempty"`
	// time log was added to fire
	AddedAt *string `form:"addedAt,omitempty" json:"addedAt,omitempty" xml:"addedAt,omitempty"`
	// weather data at time log was added to fire
	Weather *WeatherRequestBody `form:"weather,omitempty" json:"weather,omitempty" xml:"weather,omitempty"`
}

// DeleteRequestBody is the type of the "log" service "delete" endpoint HTTP
// request body.
type DeleteRequestBody struct {
	// name
	CreatedAt *string `form:"createdAt,omitempty" json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// name
	UpdatedAt *string `form:"updatedAt,omitempty" json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// size of log
	Size *string `form:"size,omitempty" json:"size,omitempty" xml:"size,omitempty"`
	// Fire identifier log belongs to
	FireID *int `form:"fireID,omitempty" json:"fireID,omitempty" xml:"fireID,omitempty"`
	// time log was added to fire
	AddedAt *string `form:"addedAt,omitempty" json:"addedAt,omitempty" xml:"addedAt,omitempty"`
	// weather data at time log was added to fire
	Weather *WeatherRequestBody `form:"weather,omitempty" json:"weather,omitempty" xml:"weather,omitempty"`
}

// ListRequestBody is the type of the "log" service "list" endpoint HTTP
// request body.
type ListRequestBody struct {
	// product filters to apply
	Filters *LogFiltersRequestBody `form:"filters,omitempty" json:"filters,omitempty" xml:"filters,omitempty"`
	// product search to apply
	Search *LogSearchRequestBody `form:"search,omitempty" json:"search,omitempty" xml:"search,omitempty"`
	// product sort to apply
	Sort *LogSortsRequestBody `form:"sort,omitempty" json:"sort,omitempty" xml:"sort,omitempty"`
	// product pagination to apply
	Pagination *LogPaginationRequestBody `form:"pagination,omitempty" json:"pagination,omitempty" xml:"pagination,omitempty"`
}

// CreateOKResponseBody is the type of the "log" service "create" endpoint HTTP
// response body.
type CreateOKResponseBody struct {
	// id
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// name
	CreatedAt *string `form:"createdAt,omitempty" json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// name
	UpdatedAt *string `form:"updatedAt,omitempty" json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// size of log
	Size *string `form:"size,omitempty" json:"size,omitempty" xml:"size,omitempty"`
	// Fire identifier log belongs to
	FireID *int `form:"fireID,omitempty" json:"fireID,omitempty" xml:"fireID,omitempty"`
	// time log was added to fire
	AddedAt *string `form:"addedAt,omitempty" json:"addedAt,omitempty" xml:"addedAt,omitempty"`
	// weather data at time log was added to fire
	Weather *WeatherResponseBody `form:"weather,omitempty" json:"weather,omitempty" xml:"weather,omitempty"`
}

// GetOKResponseBody is the type of the "log" service "get" endpoint HTTP
// response body.
type GetOKResponseBody struct {
	// id
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// name
	CreatedAt *string `form:"createdAt,omitempty" json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// name
	UpdatedAt *string `form:"updatedAt,omitempty" json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// size of log
	Size *string `form:"size,omitempty" json:"size,omitempty" xml:"size,omitempty"`
	// Fire identifier log belongs to
	FireID *int `form:"fireID,omitempty" json:"fireID,omitempty" xml:"fireID,omitempty"`
	// time log was added to fire
	AddedAt *string `form:"addedAt,omitempty" json:"addedAt,omitempty" xml:"addedAt,omitempty"`
	// weather data at time log was added to fire
	Weather *WeatherResponseBody `form:"weather,omitempty" json:"weather,omitempty" xml:"weather,omitempty"`
}

// UpdateOKResponseBody is the type of the "log" service "update" endpoint HTTP
// response body.
type UpdateOKResponseBody struct {
	// id
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// name
	CreatedAt *string `form:"createdAt,omitempty" json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// name
	UpdatedAt *string `form:"updatedAt,omitempty" json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// size of log
	Size *string `form:"size,omitempty" json:"size,omitempty" xml:"size,omitempty"`
	// Fire identifier log belongs to
	FireID *int `form:"fireID,omitempty" json:"fireID,omitempty" xml:"fireID,omitempty"`
	// time log was added to fire
	AddedAt *string `form:"addedAt,omitempty" json:"addedAt,omitempty" xml:"addedAt,omitempty"`
	// weather data at time log was added to fire
	Weather *WeatherResponseBody `form:"weather,omitempty" json:"weather,omitempty" xml:"weather,omitempty"`
}

// ListOKResponseBody is the type of the "log" service "list" endpoint HTTP
// response body.
type ListOKResponseBody struct {
	// logs
	Logs []*LogResponseBody `form:"logs,omitempty" json:"logs,omitempty" xml:"logs,omitempty"`
	// pagination results
	Pagination *PaginationResponseBody `form:"pagination,omitempty" json:"pagination,omitempty" xml:"pagination,omitempty"`
}

// CreateBadRequestResponseBody is the type of the "log" service "create"
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

// GetNotFoundResponseBody is the type of the "log" service "get" endpoint HTTP
// response body for the "not_found" error.
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

// UpdateNotFoundResponseBody is the type of the "log" service "update"
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

// DeleteNotFoundResponseBody is the type of the "log" service "delete"
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

// LogResponseBody is used to define fields on response body types.
type LogResponseBody struct {
	// id
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// name
	CreatedAt *string `form:"createdAt,omitempty" json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// name
	UpdatedAt *string `form:"updatedAt,omitempty" json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// size of log
	Size *string `form:"size,omitempty" json:"size,omitempty" xml:"size,omitempty"`
	// Fire identifier log belongs to
	FireID *int `form:"fireID,omitempty" json:"fireID,omitempty" xml:"fireID,omitempty"`
	// time log was added to fire
	AddedAt *string `form:"addedAt,omitempty" json:"addedAt,omitempty" xml:"addedAt,omitempty"`
	// weather data at time log was added to fire
	Weather *WeatherResponseBody `form:"weather,omitempty" json:"weather,omitempty" xml:"weather,omitempty"`
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

// WeatherRequestBody is used to define fields on request body types.
type WeatherRequestBody struct {
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

// LogFiltersRequestBody is used to define fields on request body types.
type LogFiltersRequestBody struct {
	Name  []*StringFilterRequestBody `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Start []*TimeFilterRequestBody   `form:"start,omitempty" json:"start,omitempty" xml:"start,omitempty"`
	End   []*TimeFilterRequestBody   `form:"end,omitempty" json:"end,omitempty" xml:"end,omitempty"`
}

// StringFilterRequestBody is used to define fields on request body types.
type StringFilterRequestBody struct {
	// filter key
	Key *string `form:"key,omitempty" json:"key,omitempty" xml:"key,omitempty"`
	// operator value
	Operator *string `form:"operator,omitempty" json:"operator,omitempty" xml:"operator,omitempty"`
	// filter value
	Value *string `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// TimeFilterRequestBody is used to define fields on request body types.
type TimeFilterRequestBody struct {
	// filter key
	Key *string `form:"key,omitempty" json:"key,omitempty" xml:"key,omitempty"`
	// operator value
	Operator *string `form:"operator,omitempty" json:"operator,omitempty" xml:"operator,omitempty"`
	// filter value
	Value *int `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
}

// LogSearchRequestBody is used to define fields on request body types.
type LogSearchRequestBody struct {
	Name        *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
}

// LogSortsRequestBody is used to define fields on request body types.
type LogSortsRequestBody struct {
	ID    *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Start *string `form:"start,omitempty" json:"start,omitempty" xml:"start,omitempty"`
	End   *string `form:"end,omitempty" json:"end,omitempty" xml:"end,omitempty"`
}

// LogPaginationRequestBody is used to define fields on request body types.
type LogPaginationRequestBody struct {
	Page  *int `form:"page,omitempty" json:"page,omitempty" xml:"page,omitempty"`
	Limit *int `form:"limit,omitempty" json:"limit,omitempty" xml:"limit,omitempty"`
}

// NewCreateOKResponseBody builds the HTTP response body from the result of the
// "create" endpoint of the "log" service.
func NewCreateOKResponseBody(res *log.Log) *CreateOKResponseBody {
	body := &CreateOKResponseBody{
		ID:        res.ID,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
		Name:      res.Name,
		Size:      res.Size,
		FireID:    res.FireID,
		AddedAt:   res.AddedAt,
	}
	if res.Weather != nil {
		body.Weather = marshalLogWeatherToWeatherResponseBody(res.Weather)
	}
	return body
}

// NewGetOKResponseBody builds the HTTP response body from the result of the
// "get" endpoint of the "log" service.
func NewGetOKResponseBody(res *log.Log) *GetOKResponseBody {
	body := &GetOKResponseBody{
		ID:        res.ID,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
		Name:      res.Name,
		Size:      res.Size,
		FireID:    res.FireID,
		AddedAt:   res.AddedAt,
	}
	if res.Weather != nil {
		body.Weather = marshalLogWeatherToWeatherResponseBody(res.Weather)
	}
	return body
}

// NewUpdateOKResponseBody builds the HTTP response body from the result of the
// "update" endpoint of the "log" service.
func NewUpdateOKResponseBody(res *log.Log) *UpdateOKResponseBody {
	body := &UpdateOKResponseBody{
		ID:        res.ID,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
		Name:      res.Name,
		Size:      res.Size,
		FireID:    res.FireID,
		AddedAt:   res.AddedAt,
	}
	if res.Weather != nil {
		body.Weather = marshalLogWeatherToWeatherResponseBody(res.Weather)
	}
	return body
}

// NewListOKResponseBody builds the HTTP response body from the result of the
// "list" endpoint of the "log" service.
func NewListOKResponseBody(res *log.LogList) *ListOKResponseBody {
	body := &ListOKResponseBody{}
	if res.Logs != nil {
		body.Logs = make([]*LogResponseBody, len(res.Logs))
		for i, val := range res.Logs {
			body.Logs[i] = marshalLogLogToLogResponseBody(val)
		}
	}
	if res.Pagination != nil {
		body.Pagination = marshalLogPaginationToPaginationResponseBody(res.Pagination)
	}
	return body
}

// NewCreateBadRequestResponseBody builds the HTTP response body from the
// result of the "create" endpoint of the "log" service.
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
// the "get" endpoint of the "log" service.
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
// of the "update" endpoint of the "log" service.
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
// of the "delete" endpoint of the "log" service.
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

// NewCreateLog builds a log service create endpoint payload.
func NewCreateLog(body *CreateRequestBody) *log.Log {
	v := &log.Log{
		ID:        body.ID,
		CreatedAt: body.CreatedAt,
		UpdatedAt: body.UpdatedAt,
		Name:      body.Name,
		Size:      body.Size,
		FireID:    body.FireID,
		AddedAt:   body.AddedAt,
	}
	if body.Weather != nil {
		v.Weather = unmarshalWeatherRequestBodyToLogWeather(body.Weather)
	}

	return v
}

// NewGetLog builds a log service get endpoint payload.
func NewGetLog(body *GetRequestBody, id int) *log.Log {
	v := &log.Log{
		CreatedAt: body.CreatedAt,
		UpdatedAt: body.UpdatedAt,
		Name:      body.Name,
		Size:      body.Size,
		FireID:    body.FireID,
		AddedAt:   body.AddedAt,
	}
	if body.Weather != nil {
		v.Weather = unmarshalWeatherRequestBodyToLogWeather(body.Weather)
	}
	v.ID = &id

	return v
}

// NewUpdateLog builds a log service update endpoint payload.
func NewUpdateLog(body *UpdateRequestBody, id int) *log.Log {
	v := &log.Log{
		CreatedAt: body.CreatedAt,
		UpdatedAt: body.UpdatedAt,
		Name:      body.Name,
		Size:      body.Size,
		FireID:    body.FireID,
		AddedAt:   body.AddedAt,
	}
	if body.Weather != nil {
		v.Weather = unmarshalWeatherRequestBodyToLogWeather(body.Weather)
	}
	v.ID = &id

	return v
}

// NewDeleteLog builds a log service delete endpoint payload.
func NewDeleteLog(body *DeleteRequestBody, id int) *log.Log {
	v := &log.Log{
		CreatedAt: body.CreatedAt,
		UpdatedAt: body.UpdatedAt,
		Name:      body.Name,
		Size:      body.Size,
		FireID:    body.FireID,
		AddedAt:   body.AddedAt,
	}
	if body.Weather != nil {
		v.Weather = unmarshalWeatherRequestBodyToLogWeather(body.Weather)
	}
	v.ID = &id

	return v
}

// NewListLogListPayload builds a log service list endpoint payload.
func NewListLogListPayload(body *ListRequestBody) *log.LogListPayload {
	v := &log.LogListPayload{}
	v.Filters = unmarshalLogFiltersRequestBodyToLogLogFilters(body.Filters)
	v.Search = unmarshalLogSearchRequestBodyToLogLogSearch(body.Search)
	v.Sort = unmarshalLogSortsRequestBodyToLogLogSorts(body.Sort)
	v.Pagination = unmarshalLogPaginationRequestBodyToLogLogPagination(body.Pagination)

	return v
}

// ValidateCreateRequestBody runs the validations defined on CreateRequestBody
func ValidateCreateRequestBody(body *CreateRequestBody) (err error) {
	if body.CreatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.createdAt", *body.CreatedAt, goa.FormatDateTime))
	}
	if body.UpdatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.updatedAt", *body.UpdatedAt, goa.FormatDateTime))
	}
	if body.Size != nil {
		if !(*body.Size == "S" || *body.Size == "M" || *body.Size == "L" || *body.Size == "XL") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.size", *body.Size, []interface{}{"S", "M", "L", "XL"}))
		}
	}
	if body.AddedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.addedAt", *body.AddedAt, goa.FormatDateTime))
	}
	if body.Weather != nil {
		if err2 := ValidateWeatherRequestBody(body.Weather); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateGetRequestBody runs the validations defined on GetRequestBody
func ValidateGetRequestBody(body *GetRequestBody) (err error) {
	if body.CreatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.createdAt", *body.CreatedAt, goa.FormatDateTime))
	}
	if body.UpdatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.updatedAt", *body.UpdatedAt, goa.FormatDateTime))
	}
	if body.Size != nil {
		if !(*body.Size == "S" || *body.Size == "M" || *body.Size == "L" || *body.Size == "XL") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.size", *body.Size, []interface{}{"S", "M", "L", "XL"}))
		}
	}
	if body.AddedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.addedAt", *body.AddedAt, goa.FormatDateTime))
	}
	if body.Weather != nil {
		if err2 := ValidateWeatherRequestBody(body.Weather); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateUpdateRequestBody runs the validations defined on UpdateRequestBody
func ValidateUpdateRequestBody(body *UpdateRequestBody) (err error) {
	if body.CreatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.createdAt", *body.CreatedAt, goa.FormatDateTime))
	}
	if body.UpdatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.updatedAt", *body.UpdatedAt, goa.FormatDateTime))
	}
	if body.Size != nil {
		if !(*body.Size == "S" || *body.Size == "M" || *body.Size == "L" || *body.Size == "XL") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.size", *body.Size, []interface{}{"S", "M", "L", "XL"}))
		}
	}
	if body.AddedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.addedAt", *body.AddedAt, goa.FormatDateTime))
	}
	if body.Weather != nil {
		if err2 := ValidateWeatherRequestBody(body.Weather); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateDeleteRequestBody runs the validations defined on DeleteRequestBody
func ValidateDeleteRequestBody(body *DeleteRequestBody) (err error) {
	if body.CreatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.createdAt", *body.CreatedAt, goa.FormatDateTime))
	}
	if body.UpdatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.updatedAt", *body.UpdatedAt, goa.FormatDateTime))
	}
	if body.Size != nil {
		if !(*body.Size == "S" || *body.Size == "M" || *body.Size == "L" || *body.Size == "XL") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.size", *body.Size, []interface{}{"S", "M", "L", "XL"}))
		}
	}
	if body.AddedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.addedAt", *body.AddedAt, goa.FormatDateTime))
	}
	if body.Weather != nil {
		if err2 := ValidateWeatherRequestBody(body.Weather); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
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
		if err2 := ValidateLogFiltersRequestBody(body.Filters); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.Sort != nil {
		if err2 := ValidateLogSortsRequestBody(body.Sort); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if body.Pagination != nil {
		if err2 := ValidateLogPaginationRequestBody(body.Pagination); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateWeatherRequestBody runs the validations defined on WeatherRequestBody
func ValidateWeatherRequestBody(body *WeatherRequestBody) (err error) {
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

// ValidateLogFiltersRequestBody runs the validations defined on
// LogFiltersRequestBody
func ValidateLogFiltersRequestBody(body *LogFiltersRequestBody) (err error) {
	for _, e := range body.Name {
		if e != nil {
			if err2 := ValidateStringFilterRequestBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	for _, e := range body.Start {
		if e != nil {
			if err2 := ValidateTimeFilterRequestBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	for _, e := range body.End {
		if e != nil {
			if err2 := ValidateTimeFilterRequestBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateStringFilterRequestBody runs the validations defined on
// StringFilterRequestBody
func ValidateStringFilterRequestBody(body *StringFilterRequestBody) (err error) {
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
		if !(*body.Operator == "=" || *body.Operator == "!=") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.operator", *body.Operator, []interface{}{"=", "!="}))
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
	if body.Key != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.key", *body.Key, goa.FormatDateTime))
	}
	if body.Operator != nil {
		if !(*body.Operator == "=" || *body.Operator == ">" || *body.Operator == ">=" || *body.Operator == "<=" || *body.Operator == "<") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.operator", *body.Operator, []interface{}{"=", ">", ">=", "<=", "<"}))
		}
	}
	return
}

// ValidateLogSortsRequestBody runs the validations defined on
// LogSortsRequestBody
func ValidateLogSortsRequestBody(body *LogSortsRequestBody) (err error) {
	if body.ID != nil {
		if !(*body.ID == "ASC, DESC") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.id", *body.ID, []interface{}{"ASC, DESC"}))
		}
	}
	if body.Start != nil {
		if !(*body.Start == "ASC, DESC") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.start", *body.Start, []interface{}{"ASC, DESC"}))
		}
	}
	if body.End != nil {
		if !(*body.End == "ASC, DESC") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.end", *body.End, []interface{}{"ASC, DESC"}))
		}
	}
	return
}

// ValidateLogPaginationRequestBody runs the validations defined on
// LogPaginationRequestBody
func ValidateLogPaginationRequestBody(body *LogPaginationRequestBody) (err error) {
	if body.Page == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("page", "body"))
	}
	if body.Limit == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("limit", "body"))
	}
	return
}
