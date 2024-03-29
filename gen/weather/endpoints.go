// Code generated by goa v3.3.1, DO NOT EDIT.
//
// weather endpoints
//
// Command:
// $ goa gen github.com/NeedMoreVolume/FireTracker/design

package weather

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "weather" service endpoints.
type Endpoints struct {
	Create goa.Endpoint
	Get    goa.Endpoint
	Update goa.Endpoint
	Delete goa.Endpoint
	List   goa.Endpoint
}

// NewEndpoints wraps the methods of the "weather" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Create: NewCreateEndpoint(s),
		Get:    NewGetEndpoint(s),
		Update: NewUpdateEndpoint(s),
		Delete: NewDeleteEndpoint(s),
		List:   NewListEndpoint(s),
	}
}

// Use applies the given middleware to all the "weather" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Create = m(e.Create)
	e.Get = m(e.Get)
	e.Update = m(e.Update)
	e.Delete = m(e.Delete)
	e.List = m(e.List)
}

// NewCreateEndpoint returns an endpoint function that calls the method
// "create" of service "weather".
func NewCreateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*Weather)
		return s.Create(ctx, p)
	}
}

// NewGetEndpoint returns an endpoint function that calls the method "get" of
// service "weather".
func NewGetEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*Weather)
		return s.Get(ctx, p)
	}
}

// NewUpdateEndpoint returns an endpoint function that calls the method
// "update" of service "weather".
func NewUpdateEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*Weather)
		return s.Update(ctx, p)
	}
}

// NewDeleteEndpoint returns an endpoint function that calls the method
// "delete" of service "weather".
func NewDeleteEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*Weather)
		return nil, s.Delete(ctx, p)
	}
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "weather".
func NewListEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*WeatherListPayload)
		return s.List(ctx, p)
	}
}
