// Code generated by goa v3.2.5, DO NOT EDIT.
//
// fire gRPC client CLI support package
//
// Command:
// $ goa gen github.com/NeedMoreVolume/FireTracker/design

package client

import (
	"encoding/json"
	"fmt"

	fire "github.com/NeedMoreVolume/FireTracker/gen/fire"
	firepb "github.com/NeedMoreVolume/FireTracker/gen/grpc/fire/pb"
)

// BuildCreatePayload builds the payload for the fire create endpoint from CLI
// flags.
func BuildCreatePayload(fireCreateMessage string) (*fire.Fire, error) {
	var err error
	var message firepb.CreateRequest
	{
		if fireCreateMessage != "" {
			err = json.Unmarshal([]byte(fireCreateMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"createdAt\": \"2001-07-26T15:42:42Z\",\n      \"deletedAt\": \"2008-01-23T13:39:15Z\",\n      \"description\": \"Vero dolorum enim et aut.\",\n      \"end\": \"2002-03-03T18:26:05Z\",\n      \"id\": 5117576016712403101,\n      \"name\": \"Aliquid odio asperiores totam dignissimos.\",\n      \"start\": \"1987-09-07T04:18:11Z\",\n      \"updatedAt\": \"1985-03-21T02:39:44Z\"\n   }'")
			}
		}
	}
	v := &fire.Fire{}
	if message.Id != 0 {
		idptr := int(message.Id)
		v.ID = &idptr
	}
	if message.CreatedAt != "" {
		v.CreatedAt = &message.CreatedAt
	}
	if message.UpdatedAt != "" {
		v.UpdatedAt = &message.UpdatedAt
	}
	if message.DeletedAt != "" {
		v.DeletedAt = &message.DeletedAt
	}
	if message.Name != "" {
		v.Name = &message.Name
	}
	if message.Description != "" {
		v.Description = &message.Description
	}
	if message.Start != "" {
		v.Start = &message.Start
	}
	if message.End != "" {
		v.End = &message.End
	}

	return v, nil
}

// BuildGetPayload builds the payload for the fire get endpoint from CLI flags.
func BuildGetPayload(fireGetMessage string) (*fire.GetPayload, error) {
	var err error
	var message firepb.GetRequest
	{
		if fireGetMessage != "" {
			err = json.Unmarshal([]byte(fireGetMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"id\": 3471621047790542047\n   }'")
			}
		}
	}
	v := &fire.GetPayload{}
	if message.Id != 0 {
		idptr := int(message.Id)
		v.ID = &idptr
	}

	return v, nil
}

// BuildUpdatePayload builds the payload for the fire update endpoint from CLI
// flags.
func BuildUpdatePayload(fireUpdateMessage string) (*fire.Fire, error) {
	var err error
	var message firepb.UpdateRequest
	{
		if fireUpdateMessage != "" {
			err = json.Unmarshal([]byte(fireUpdateMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"createdAt\": \"1979-12-25T16:28:11Z\",\n      \"deletedAt\": \"2009-02-13T05:16:56Z\",\n      \"description\": \"Consequuntur ea ipsum deleniti modi.\",\n      \"end\": \"1976-08-28T20:58:47Z\",\n      \"id\": 5069143882329837951,\n      \"name\": \"Dolor quis sit.\",\n      \"start\": \"1988-03-22T03:56:20Z\",\n      \"updatedAt\": \"2006-06-29T13:18:00Z\"\n   }'")
			}
		}
	}
	v := &fire.Fire{}
	if message.Id != 0 {
		idptr := int(message.Id)
		v.ID = &idptr
	}
	if message.CreatedAt != "" {
		v.CreatedAt = &message.CreatedAt
	}
	if message.UpdatedAt != "" {
		v.UpdatedAt = &message.UpdatedAt
	}
	if message.DeletedAt != "" {
		v.DeletedAt = &message.DeletedAt
	}
	if message.Name != "" {
		v.Name = &message.Name
	}
	if message.Description != "" {
		v.Description = &message.Description
	}
	if message.Start != "" {
		v.Start = &message.Start
	}
	if message.End != "" {
		v.End = &message.End
	}

	return v, nil
}

// BuildDeletePayload builds the payload for the fire delete endpoint from CLI
// flags.
func BuildDeletePayload(fireDeleteMessage string) (*fire.DeletePayload, error) {
	var err error
	var message firepb.DeleteRequest
	{
		if fireDeleteMessage != "" {
			err = json.Unmarshal([]byte(fireDeleteMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"id\": 1211853009398384327\n   }'")
			}
		}
	}
	v := &fire.DeletePayload{}
	if message.Id != 0 {
		idptr := int(message.Id)
		v.ID = &idptr
	}

	return v, nil
}

// BuildListPayload builds the payload for the fire list endpoint from CLI
// flags.
func BuildListPayload(fireListMessage string) (*fire.FireListPayload, error) {
	var err error
	var message firepb.ListRequest
	{
		if fireListMessage != "" {
			err = json.Unmarshal([]byte(fireListMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"filters\": {\n         \"end\": [\n            {\n               \"key\": \"1996-06-12T08:40:44Z\",\n               \"operator\": \"\\u003e=\",\n               \"value\": 5573547598936979342\n            },\n            {\n               \"key\": \"1996-06-12T08:40:44Z\",\n               \"operator\": \"\\u003e=\",\n               \"value\": 5573547598936979342\n            },\n            {\n               \"key\": \"1996-06-12T08:40:44Z\",\n               \"operator\": \"\\u003e=\",\n               \"value\": 5573547598936979342\n            },\n            {\n               \"key\": \"1996-06-12T08:40:44Z\",\n               \"operator\": \"\\u003e=\",\n               \"value\": 5573547598936979342\n            }\n         ],\n         \"name\": [\n            {\n               \"key\": \"Qui hic.\",\n               \"operator\": \"!=\",\n               \"value\": \"Quia sint eligendi eum architecto sint est.\"\n            },\n            {\n               \"key\": \"Qui hic.\",\n               \"operator\": \"!=\",\n               \"value\": \"Quia sint eligendi eum architecto sint est.\"\n            },\n            {\n               \"key\": \"Qui hic.\",\n               \"operator\": \"!=\",\n               \"value\": \"Quia sint eligendi eum architecto sint est.\"\n            }\n         ],\n         \"start\": [\n            {\n               \"key\": \"1996-06-12T08:40:44Z\",\n               \"operator\": \"\\u003e=\",\n               \"value\": 5573547598936979342\n            },\n            {\n               \"key\": \"1996-06-12T08:40:44Z\",\n               \"operator\": \"\\u003e=\",\n               \"value\": 5573547598936979342\n            }\n         ]\n      },\n      \"pagination\": {\n         \"limit\": 7720515675719359881,\n         \"page\": 5747059457513862389\n      },\n      \"search\": {\n         \"description\": \"Rem quaerat vel vero velit et.\",\n         \"name\": \"Quo reiciendis omnis ut.\"\n      },\n      \"sort\": {\n         \"end\": \"ASC, DESC\",\n         \"id\": \"ASC, DESC\",\n         \"start\": \"ASC, DESC\"\n      }\n   }'")
			}
		}
	}
	v := &fire.FireListPayload{}
	if message.Filters != nil {
		v.Filters = protobufFirepbFireFiltersToFireFireFilters(message.Filters)
	}
	if message.Search != nil {
		v.Search = protobufFirepbFireSearchToFireFireSearch(message.Search)
	}
	if message.Sort != nil {
		v.Sort = protobufFirepbFireSortsToFireFireSorts(message.Sort)
	}
	if message.Pagination != nil {
		v.Pagination = protobufFirepbFirePaginationToFireFirePagination(message.Pagination)
	}

	return v, nil
}

// BuildGetWeatherForFirePayload builds the payload for the fire
// getWeatherForFire endpoint from CLI flags.
func BuildGetWeatherForFirePayload(fireGetWeatherForFireMessage string) (*fire.GetWeatherForFirePayload, error) {
	var err error
	var message firepb.GetWeatherForFireRequest
	{
		if fireGetWeatherForFireMessage != "" {
			err = json.Unmarshal([]byte(fireGetWeatherForFireMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"id\": 141765740622424101\n   }'")
			}
		}
	}
	v := &fire.GetWeatherForFirePayload{}
	if message.Id != 0 {
		idptr := int(message.Id)
		v.ID = &idptr
	}

	return v, nil
}

// BuildGetLogsForFirePayload builds the payload for the fire getLogsForFire
// endpoint from CLI flags.
func BuildGetLogsForFirePayload(fireGetLogsForFireMessage string) (*fire.GetLogsForFirePayload, error) {
	var err error
	var message firepb.GetLogsForFireRequest
	{
		if fireGetLogsForFireMessage != "" {
			err = json.Unmarshal([]byte(fireGetLogsForFireMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"id\": 5391076212118225793\n   }'")
			}
		}
	}
	v := &fire.GetLogsForFirePayload{}
	if message.Id != 0 {
		idptr := int(message.Id)
		v.ID = &idptr
	}

	return v, nil
}