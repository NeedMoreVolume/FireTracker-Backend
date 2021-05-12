// Code generated by goa v3.3.1, DO NOT EDIT.
//
// fire client
//
// Command:
// $ goa gen github.com/NeedMoreVolume/FireTracker/design

package fire

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "fire" service client.
type Client struct {
	CreateEndpoint            goa.Endpoint
	GetEndpoint               goa.Endpoint
	UpdateEndpoint            goa.Endpoint
	DeleteEndpoint            goa.Endpoint
	ListEndpoint              goa.Endpoint
	GetWeatherForFireEndpoint goa.Endpoint
	GetLogsForFireEndpoint    goa.Endpoint
}

// NewClient initializes a "fire" service client given the endpoints.
func NewClient(create, get, update, delete_, list, getWeatherForFire, getLogsForFire goa.Endpoint) *Client {
	return &Client{
		CreateEndpoint:            create,
		GetEndpoint:               get,
		UpdateEndpoint:            update,
		DeleteEndpoint:            delete_,
		ListEndpoint:              list,
		GetWeatherForFireEndpoint: getWeatherForFire,
		GetLogsForFireEndpoint:    getLogsForFire,
	}
}

// Create calls the "create" endpoint of the "fire" service.
func (c *Client) Create(ctx context.Context, p *Fire) (res *Fire, err error) {
	var ires interface{}
	ires, err = c.CreateEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Fire), nil
}

// Get calls the "get" endpoint of the "fire" service.
func (c *Client) Get(ctx context.Context, p *GetPayload) (res *Fire, err error) {
	var ires interface{}
	ires, err = c.GetEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Fire), nil
}

// Update calls the "update" endpoint of the "fire" service.
func (c *Client) Update(ctx context.Context, p *Fire) (res *Fire, err error) {
	var ires interface{}
	ires, err = c.UpdateEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Fire), nil
}

// Delete calls the "delete" endpoint of the "fire" service.
func (c *Client) Delete(ctx context.Context, p *DeletePayload) (err error) {
	_, err = c.DeleteEndpoint(ctx, p)
	return
}

// List calls the "list" endpoint of the "fire" service.
func (c *Client) List(ctx context.Context, p *FireListPayload) (res *FireList, err error) {
	var ires interface{}
	ires, err = c.ListEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*FireList), nil
}

// GetWeatherForFire calls the "getWeatherForFire" endpoint of the "fire"
// service.
func (c *Client) GetWeatherForFire(ctx context.Context, p *GetWeatherForFirePayload) (res *WeatherList, err error) {
	var ires interface{}
	ires, err = c.GetWeatherForFireEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*WeatherList), nil
}

// GetLogsForFire calls the "getLogsForFire" endpoint of the "fire" service.
func (c *Client) GetLogsForFire(ctx context.Context, p *GetLogsForFirePayload) (res *LogList, err error) {
	var ires interface{}
	ires, err = c.GetLogsForFireEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*LogList), nil
}
