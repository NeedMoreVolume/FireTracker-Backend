// Code generated by goa v3.2.5, DO NOT EDIT.
//
// fire HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/NeedMoreVolume/FireTracker/design

package server

import (
	"context"
	"io"
	"net/http"
	"strconv"

	fire "github.com/NeedMoreVolume/FireTracker/gen/fire"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeCreateResponse returns an encoder for responses returned by the fire
// create endpoint.
func EncodeCreateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*fire.Fire)
		enc := encoder(ctx, w)
		body := NewCreateOKResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeCreateRequest returns a decoder for requests sent to the fire create
// endpoint.
func DecodeCreateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body CreateRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateCreateRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewCreateFire(&body)

		return payload, nil
	}
}

// EncodeCreateError returns an encoder for errors returned by the create fire
// endpoint.
func EncodeCreateError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "bad_request":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewCreateBadRequestResponseBody(res)
			}
			w.Header().Set("goa-error", "bad_request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeGetResponse returns an encoder for responses returned by the fire get
// endpoint.
func EncodeGetResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*fire.Fire)
		enc := encoder(ctx, w)
		body := NewGetOKResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetRequest returns a decoder for requests sent to the fire get
// endpoint.
func DecodeGetRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id  int
			err error

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseInt(idRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "integer"))
			}
			id = int(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewGetPayload(id)

		return payload, nil
	}
}

// EncodeGetError returns an encoder for errors returned by the get fire
// endpoint.
func EncodeGetError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewGetNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "not_found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeUpdateResponse returns an encoder for responses returned by the fire
// update endpoint.
func EncodeUpdateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*fire.Fire)
		enc := encoder(ctx, w)
		body := NewUpdateOKResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeUpdateRequest returns a decoder for requests sent to the fire update
// endpoint.
func DecodeUpdateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body UpdateRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateUpdateRequestBody(&body)
		if err != nil {
			return nil, err
		}

		var (
			id int

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseInt(idRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "integer"))
			}
			id = int(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewUpdateFire(&body, id)

		return payload, nil
	}
}

// EncodeUpdateError returns an encoder for errors returned by the update fire
// endpoint.
func EncodeUpdateError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewUpdateNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "not_found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeDeleteResponse returns an encoder for responses returned by the fire
// delete endpoint.
func EncodeDeleteResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// DecodeDeleteRequest returns a decoder for requests sent to the fire delete
// endpoint.
func DecodeDeleteRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id  int
			err error

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseInt(idRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "integer"))
			}
			id = int(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewDeletePayload(id)

		return payload, nil
	}
}

// EncodeDeleteError returns an encoder for errors returned by the delete fire
// endpoint.
func EncodeDeleteError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewDeleteNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "not_found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeListResponse returns an encoder for responses returned by the fire
// list endpoint.
func EncodeListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*fire.FireList)
		enc := encoder(ctx, w)
		body := NewListOKResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeListRequest returns a decoder for requests sent to the fire list
// endpoint.
func DecodeListRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body ListRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateListRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewListFireListPayload(&body)

		return payload, nil
	}
}

// EncodeGetWeatherForFireResponse returns an encoder for responses returned by
// the fire getWeatherForFire endpoint.
func EncodeGetWeatherForFireResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*fire.WeatherList)
		enc := encoder(ctx, w)
		body := NewGetWeatherForFireOKResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetWeatherForFireRequest returns a decoder for requests sent to the
// fire getWeatherForFire endpoint.
func DecodeGetWeatherForFireRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id  int
			err error

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseInt(idRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "integer"))
			}
			id = int(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewGetWeatherForFirePayload(id)

		return payload, nil
	}
}

// EncodeGetWeatherForFireError returns an encoder for errors returned by the
// getWeatherForFire fire endpoint.
func EncodeGetWeatherForFireError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewGetWeatherForFireNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "not_found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeGetLogsForFireResponse returns an encoder for responses returned by
// the fire getLogsForFire endpoint.
func EncodeGetLogsForFireResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*fire.LogList)
		enc := encoder(ctx, w)
		body := NewGetLogsForFireOKResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeGetLogsForFireRequest returns a decoder for requests sent to the fire
// getLogsForFire endpoint.
func DecodeGetLogsForFireRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id  int
			err error

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseInt(idRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "integer"))
			}
			id = int(v)
		}
		if err != nil {
			return nil, err
		}
		payload := NewGetLogsForFirePayload(id)

		return payload, nil
	}
}

// EncodeGetLogsForFireError returns an encoder for errors returned by the
// getLogsForFire fire endpoint.
func EncodeGetLogsForFireError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewGetLogsForFireNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "not_found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// unmarshalFireFiltersRequestBodyToFireFireFilters builds a value of type
// *fire.FireFilters from a value of type *FireFiltersRequestBody.
func unmarshalFireFiltersRequestBodyToFireFireFilters(v *FireFiltersRequestBody) *fire.FireFilters {
	res := &fire.FireFilters{}
	if v.Name != nil {
		res.Name = make([]*fire.StringFilter, len(v.Name))
		for i, val := range v.Name {
			res.Name[i] = unmarshalStringFilterRequestBodyToFireStringFilter(val)
		}
	}
	if v.Start != nil {
		res.Start = make([]*fire.TimeFilter, len(v.Start))
		for i, val := range v.Start {
			res.Start[i] = unmarshalTimeFilterRequestBodyToFireTimeFilter(val)
		}
	}
	if v.End != nil {
		res.End = make([]*fire.TimeFilter, len(v.End))
		for i, val := range v.End {
			res.End[i] = unmarshalTimeFilterRequestBodyToFireTimeFilter(val)
		}
	}

	return res
}

// unmarshalStringFilterRequestBodyToFireStringFilter builds a value of type
// *fire.StringFilter from a value of type *StringFilterRequestBody.
func unmarshalStringFilterRequestBodyToFireStringFilter(v *StringFilterRequestBody) *fire.StringFilter {
	if v == nil {
		return nil
	}
	res := &fire.StringFilter{
		Key:      *v.Key,
		Operator: *v.Operator,
		Value:    *v.Value,
	}

	return res
}

// unmarshalTimeFilterRequestBodyToFireTimeFilter builds a value of type
// *fire.TimeFilter from a value of type *TimeFilterRequestBody.
func unmarshalTimeFilterRequestBodyToFireTimeFilter(v *TimeFilterRequestBody) *fire.TimeFilter {
	if v == nil {
		return nil
	}
	res := &fire.TimeFilter{
		Key:      *v.Key,
		Operator: *v.Operator,
		Value:    *v.Value,
	}

	return res
}

// unmarshalFireSearchRequestBodyToFireFireSearch builds a value of type
// *fire.FireSearch from a value of type *FireSearchRequestBody.
func unmarshalFireSearchRequestBodyToFireFireSearch(v *FireSearchRequestBody) *fire.FireSearch {
	res := &fire.FireSearch{
		Name:        v.Name,
		Description: v.Description,
	}

	return res
}

// unmarshalFireSortsRequestBodyToFireFireSorts builds a value of type
// *fire.FireSorts from a value of type *FireSortsRequestBody.
func unmarshalFireSortsRequestBodyToFireFireSorts(v *FireSortsRequestBody) *fire.FireSorts {
	res := &fire.FireSorts{
		ID:    v.ID,
		Start: v.Start,
		End:   v.End,
	}

	return res
}

// unmarshalFirePaginationRequestBodyToFireFirePagination builds a value of
// type *fire.FirePagination from a value of type *FirePaginationRequestBody.
func unmarshalFirePaginationRequestBodyToFireFirePagination(v *FirePaginationRequestBody) *fire.FirePagination {
	res := &fire.FirePagination{
		Page:  *v.Page,
		Limit: *v.Limit,
	}

	return res
}

// marshalFireFireToFireResponseBody builds a value of type *FireResponseBody
// from a value of type *fire.Fire.
func marshalFireFireToFireResponseBody(v *fire.Fire) *FireResponseBody {
	if v == nil {
		return nil
	}
	res := &FireResponseBody{
		ID:          v.ID,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
		DeletedAt:   v.DeletedAt,
		Name:        v.Name,
		Description: v.Description,
		Start:       v.Start,
		End:         v.End,
	}

	return res
}

// marshalFirePaginationToPaginationResponseBody builds a value of type
// *PaginationResponseBody from a value of type *fire.Pagination.
func marshalFirePaginationToPaginationResponseBody(v *fire.Pagination) *PaginationResponseBody {
	if v == nil {
		return nil
	}
	res := &PaginationResponseBody{
		Total: v.Total,
		Page:  v.Page,
		Limit: v.Limit,
	}

	return res
}

// marshalFireWeatherToWeatherResponseBody builds a value of type
// *WeatherResponseBody from a value of type *fire.Weather.
func marshalFireWeatherToWeatherResponseBody(v *fire.Weather) *WeatherResponseBody {
	if v == nil {
		return nil
	}
	res := &WeatherResponseBody{
		ID:          v.ID,
		CreatedAt:   v.CreatedAt,
		FireID:      v.FireID,
		LogID:       v.LogID,
		Humidity:    v.Humidity,
		WeatherType: v.WeatherType,
		WeatherTime: v.WeatherTime,
	}
	if v.Temperature != nil {
		res.Temperature = marshalFireTemperatureToTemperatureResponseBody(v.Temperature)
	}
	if v.DewPoint != nil {
		res.DewPoint = marshalFireTemperatureToTemperatureResponseBody(v.DewPoint)
	}
	if v.Wind != nil {
		res.Wind = marshalFireWindToWindResponseBody(v.Wind)
	}

	return res
}

// marshalFireTemperatureToTemperatureResponseBody builds a value of type
// *TemperatureResponseBody from a value of type *fire.Temperature.
func marshalFireTemperatureToTemperatureResponseBody(v *fire.Temperature) *TemperatureResponseBody {
	if v == nil {
		return nil
	}
	res := &TemperatureResponseBody{
		Unit:  v.Unit,
		Value: v.Value,
	}

	return res
}

// marshalFireWindToWindResponseBody builds a value of type *WindResponseBody
// from a value of type *fire.Wind.
func marshalFireWindToWindResponseBody(v *fire.Wind) *WindResponseBody {
	if v == nil {
		return nil
	}
	res := &WindResponseBody{
		Speed:     v.Speed,
		Direction: v.Direction,
		Unit:      v.Unit,
	}

	return res
}

// marshalFireLogToLogResponseBody builds a value of type *LogResponseBody from
// a value of type *fire.Log.
func marshalFireLogToLogResponseBody(v *fire.Log) *LogResponseBody {
	if v == nil {
		return nil
	}
	res := &LogResponseBody{
		ID:        v.ID,
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
		Name:      v.Name,
		Size:      v.Size,
		FireID:    v.FireID,
		AddedAt:   v.AddedAt,
	}
	if v.Weather != nil {
		res.Weather = marshalFireWeatherToWeatherResponseBody(v.Weather)
	}

	return res
}