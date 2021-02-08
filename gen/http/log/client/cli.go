// Code generated by goa v3.2.5, DO NOT EDIT.
//
// log HTTP client CLI support package
//
// Command:
// $ goa gen github.com/NeedMoreVolume/FireTracker/design

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	log "github.com/NeedMoreVolume/FireTracker/gen/log"
	goa "goa.design/goa/v3/pkg"
)

// BuildCreatePayload builds the payload for the log create endpoint from CLI
// flags.
func BuildCreatePayload(logCreateBody string) (*log.Log, error) {
	var err error
	var body CreateRequestBody
	{
		err = json.Unmarshal([]byte(logCreateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"addedAt\": \"1981-02-18T04:47:49Z\",\n      \"createdAt\": \"1992-06-27T21:53:41Z\",\n      \"fireID\": 5692558848136296972,\n      \"id\": 6577929380304453415,\n      \"name\": \"Doloremque qui non.\",\n      \"size\": \"M\",\n      \"updatedAt\": \"1985-09-10T06:26:16Z\",\n      \"weather\": {\n         \"createdAt\": \"2002-05-27T23:56:45Z\",\n         \"dewPoint\": {\n            \"unit\": \"K\",\n            \"value\": 1984646352\n         },\n         \"fireID\": 4618423612491386150,\n         \"humidity\": 1322058308,\n         \"id\": 743463055768621296,\n         \"logID\": 3982859360461558219,\n         \"temperature\": {\n            \"unit\": \"K\",\n            \"value\": 1984646352\n         },\n         \"weatherTime\": \"1986-11-28T16:39:22Z\",\n         \"weatherType\": \"Windy\",\n         \"wind\": {\n            \"direction\": \"E\",\n            \"speed\": 762000299,\n            \"unit\": \"KPH\"\n         }\n      }\n   }'")
		}
	}
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
		v.Weather = marshalWeatherRequestBodyToLogWeather(body.Weather)
	}

	return v, nil
}

// BuildGetPayload builds the payload for the log get endpoint from CLI flags.
func BuildGetPayload(logGetBody string, logGetID string) (*log.Log, error) {
	var err error
	var body GetRequestBody
	{
		err = json.Unmarshal([]byte(logGetBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"addedAt\": \"2002-04-25T04:42:22Z\",\n      \"createdAt\": \"1998-12-11T12:17:29Z\",\n      \"fireID\": 6713603813374072266,\n      \"name\": \"Deserunt et et cumque nesciunt.\",\n      \"size\": \"XL\",\n      \"updatedAt\": \"1986-02-10T00:00:19Z\",\n      \"weather\": {\n         \"createdAt\": \"2002-05-27T23:56:45Z\",\n         \"dewPoint\": {\n            \"unit\": \"K\",\n            \"value\": 1984646352\n         },\n         \"fireID\": 4618423612491386150,\n         \"humidity\": 1322058308,\n         \"id\": 743463055768621296,\n         \"logID\": 3982859360461558219,\n         \"temperature\": {\n            \"unit\": \"K\",\n            \"value\": 1984646352\n         },\n         \"weatherTime\": \"1986-11-28T16:39:22Z\",\n         \"weatherType\": \"Windy\",\n         \"wind\": {\n            \"direction\": \"E\",\n            \"speed\": 762000299,\n            \"unit\": \"KPH\"\n         }\n      }\n   }'")
		}
	}
	var id int
	{
		var v int64
		v, err = strconv.ParseInt(logGetID, 10, 64)
		id = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT")
		}
	}
	v := &log.Log{
		CreatedAt: body.CreatedAt,
		UpdatedAt: body.UpdatedAt,
		Name:      body.Name,
		Size:      body.Size,
		FireID:    body.FireID,
		AddedAt:   body.AddedAt,
	}
	if body.Weather != nil {
		v.Weather = marshalWeatherRequestBodyToLogWeather(body.Weather)
	}
	v.ID = &id

	return v, nil
}

// BuildUpdatePayload builds the payload for the log update endpoint from CLI
// flags.
func BuildUpdatePayload(logUpdateBody string, logUpdateID string) (*log.Log, error) {
	var err error
	var body UpdateRequestBody
	{
		err = json.Unmarshal([]byte(logUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"addedAt\": \"1985-10-08T03:04:17Z\",\n      \"createdAt\": \"2015-07-05T16:03:55Z\",\n      \"fireID\": 9073437711020541164,\n      \"name\": \"Labore iusto.\",\n      \"size\": \"XL\",\n      \"updatedAt\": \"2016-01-31T04:23:21Z\",\n      \"weather\": {\n         \"createdAt\": \"2002-05-27T23:56:45Z\",\n         \"dewPoint\": {\n            \"unit\": \"K\",\n            \"value\": 1984646352\n         },\n         \"fireID\": 4618423612491386150,\n         \"humidity\": 1322058308,\n         \"id\": 743463055768621296,\n         \"logID\": 3982859360461558219,\n         \"temperature\": {\n            \"unit\": \"K\",\n            \"value\": 1984646352\n         },\n         \"weatherTime\": \"1986-11-28T16:39:22Z\",\n         \"weatherType\": \"Windy\",\n         \"wind\": {\n            \"direction\": \"E\",\n            \"speed\": 762000299,\n            \"unit\": \"KPH\"\n         }\n      }\n   }'")
		}
	}
	var id int
	{
		var v int64
		v, err = strconv.ParseInt(logUpdateID, 10, 64)
		id = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT")
		}
	}
	v := &log.Log{
		CreatedAt: body.CreatedAt,
		UpdatedAt: body.UpdatedAt,
		Name:      body.Name,
		Size:      body.Size,
		FireID:    body.FireID,
		AddedAt:   body.AddedAt,
	}
	if body.Weather != nil {
		v.Weather = marshalWeatherRequestBodyToLogWeather(body.Weather)
	}
	v.ID = &id

	return v, nil
}

// BuildDeletePayload builds the payload for the log delete endpoint from CLI
// flags.
func BuildDeletePayload(logDeleteBody string, logDeleteID string) (*log.Log, error) {
	var err error
	var body DeleteRequestBody
	{
		err = json.Unmarshal([]byte(logDeleteBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"addedAt\": \"2013-07-20T12:54:06Z\",\n      \"createdAt\": \"2000-03-07T22:36:50Z\",\n      \"fireID\": 3117523544869095399,\n      \"name\": \"Rerum et a qui.\",\n      \"size\": \"L\",\n      \"updatedAt\": \"2002-08-31T02:52:53Z\",\n      \"weather\": {\n         \"createdAt\": \"2002-05-27T23:56:45Z\",\n         \"dewPoint\": {\n            \"unit\": \"K\",\n            \"value\": 1984646352\n         },\n         \"fireID\": 4618423612491386150,\n         \"humidity\": 1322058308,\n         \"id\": 743463055768621296,\n         \"logID\": 3982859360461558219,\n         \"temperature\": {\n            \"unit\": \"K\",\n            \"value\": 1984646352\n         },\n         \"weatherTime\": \"1986-11-28T16:39:22Z\",\n         \"weatherType\": \"Windy\",\n         \"wind\": {\n            \"direction\": \"E\",\n            \"speed\": 762000299,\n            \"unit\": \"KPH\"\n         }\n      }\n   }'")
		}
	}
	var id int
	{
		var v int64
		v, err = strconv.ParseInt(logDeleteID, 10, 64)
		id = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT")
		}
	}
	v := &log.Log{
		CreatedAt: body.CreatedAt,
		UpdatedAt: body.UpdatedAt,
		Name:      body.Name,
		Size:      body.Size,
		FireID:    body.FireID,
		AddedAt:   body.AddedAt,
	}
	if body.Weather != nil {
		v.Weather = marshalWeatherRequestBodyToLogWeather(body.Weather)
	}
	v.ID = &id

	return v, nil
}

// BuildListPayload builds the payload for the log list endpoint from CLI flags.
func BuildListPayload(logListBody string) (*log.LogListPayload, error) {
	var err error
	var body ListRequestBody
	{
		err = json.Unmarshal([]byte(logListBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"filters\": {\n         \"end\": [\n            {\n               \"key\": \"2004-12-16T02:06:51Z\",\n               \"operator\": \"=\",\n               \"value\": 3707112604264870284\n            },\n            {\n               \"key\": \"2004-12-16T02:06:51Z\",\n               \"operator\": \"=\",\n               \"value\": 3707112604264870284\n            }\n         ],\n         \"name\": [\n            {\n               \"key\": \"In reprehenderit veniam asperiores aliquid.\",\n               \"operator\": \"!=\",\n               \"value\": \"Quod in iusto.\"\n            },\n            {\n               \"key\": \"In reprehenderit veniam asperiores aliquid.\",\n               \"operator\": \"!=\",\n               \"value\": \"Quod in iusto.\"\n            },\n            {\n               \"key\": \"In reprehenderit veniam asperiores aliquid.\",\n               \"operator\": \"!=\",\n               \"value\": \"Quod in iusto.\"\n            }\n         ],\n         \"start\": [\n            {\n               \"key\": \"2004-12-16T02:06:51Z\",\n               \"operator\": \"=\",\n               \"value\": 3707112604264870284\n            },\n            {\n               \"key\": \"2004-12-16T02:06:51Z\",\n               \"operator\": \"=\",\n               \"value\": 3707112604264870284\n            },\n            {\n               \"key\": \"2004-12-16T02:06:51Z\",\n               \"operator\": \"=\",\n               \"value\": 3707112604264870284\n            }\n         ]\n      },\n      \"pagination\": {\n         \"limit\": 6359683262686384173,\n         \"page\": 448613042018020648\n      },\n      \"search\": {\n         \"description\": \"Quos atque incidunt quo reiciendis nisi magnam.\",\n         \"name\": \"Qui aliquam.\"\n      },\n      \"sort\": {\n         \"end\": \"ASC, DESC\",\n         \"id\": \"ASC, DESC\",\n         \"start\": \"ASC, DESC\"\n      }\n   }'")
		}
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
		if err != nil {
			return nil, err
		}
	}
	v := &log.LogListPayload{}
	if body.Filters != nil {
		v.Filters = marshalLogFiltersRequestBodyToLogLogFilters(body.Filters)
	}
	if body.Search != nil {
		v.Search = marshalLogSearchRequestBodyToLogLogSearch(body.Search)
	}
	if body.Sort != nil {
		v.Sort = marshalLogSortsRequestBodyToLogLogSorts(body.Sort)
	}
	if body.Pagination != nil {
		v.Pagination = marshalLogPaginationRequestBodyToLogLogPagination(body.Pagination)
	}

	return v, nil
}
