// Code generated by goa v3.3.1, DO NOT EDIT.
//
// weather gRPC server
//
// Command:
// $ goa gen github.com/NeedMoreVolume/FireTracker/design

package server

import (
	"context"

	weatherpb "github.com/NeedMoreVolume/FireTracker/gen/grpc/weather/pb"
	weather "github.com/NeedMoreVolume/FireTracker/gen/weather"
	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc/codes"
)

// Server implements the weatherpb.WeatherServer interface.
type Server struct {
	CreateH goagrpc.UnaryHandler
	GetH    goagrpc.UnaryHandler
	UpdateH goagrpc.UnaryHandler
	DeleteH goagrpc.UnaryHandler
	ListH   goagrpc.UnaryHandler
	weatherpb.UnimplementedWeatherServer
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the expr.
type ErrorNamer interface {
	ErrorName() string
}

// New instantiates the server struct with the weather service endpoints.
func New(e *weather.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		CreateH: NewCreateHandler(e.Create, uh),
		GetH:    NewGetHandler(e.Get, uh),
		UpdateH: NewUpdateHandler(e.Update, uh),
		DeleteH: NewDeleteHandler(e.Delete, uh),
		ListH:   NewListHandler(e.List, uh),
	}
}

// NewCreateHandler creates a gRPC handler which serves the "weather" service
// "create" endpoint.
func NewCreateHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeCreateRequest, EncodeCreateResponse)
	}
	return h
}

// Create implements the "Create" method in weatherpb.WeatherServer interface.
func (s *Server) Create(ctx context.Context, message *weatherpb.CreateRequest) (*weatherpb.CreateResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "create")
	ctx = context.WithValue(ctx, goa.ServiceKey, "weather")
	resp, err := s.CreateH.Handle(ctx, message)
	if err != nil {
		if en, ok := err.(ErrorNamer); ok {
			switch en.ErrorName() {
			case "bad_request":
				return nil, goagrpc.NewStatusError(codes.InvalidArgument, err, goagrpc.NewErrorResponse(err))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*weatherpb.CreateResponse), nil
}

// NewGetHandler creates a gRPC handler which serves the "weather" service
// "get" endpoint.
func NewGetHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeGetRequest, EncodeGetResponse)
	}
	return h
}

// Get implements the "Get" method in weatherpb.WeatherServer interface.
func (s *Server) Get(ctx context.Context, message *weatherpb.GetRequest) (*weatherpb.GetResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "get")
	ctx = context.WithValue(ctx, goa.ServiceKey, "weather")
	resp, err := s.GetH.Handle(ctx, message)
	if err != nil {
		if en, ok := err.(ErrorNamer); ok {
			switch en.ErrorName() {
			case "not_found":
				return nil, goagrpc.NewStatusError(404, err, goagrpc.NewErrorResponse(err))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*weatherpb.GetResponse), nil
}

// NewUpdateHandler creates a gRPC handler which serves the "weather" service
// "update" endpoint.
func NewUpdateHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeUpdateRequest, EncodeUpdateResponse)
	}
	return h
}

// Update implements the "Update" method in weatherpb.WeatherServer interface.
func (s *Server) Update(ctx context.Context, message *weatherpb.UpdateRequest) (*weatherpb.UpdateResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "update")
	ctx = context.WithValue(ctx, goa.ServiceKey, "weather")
	resp, err := s.UpdateH.Handle(ctx, message)
	if err != nil {
		if en, ok := err.(ErrorNamer); ok {
			switch en.ErrorName() {
			case "not_found":
				return nil, goagrpc.NewStatusError(codes.NotFound, err, goagrpc.NewErrorResponse(err))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*weatherpb.UpdateResponse), nil
}

// NewDeleteHandler creates a gRPC handler which serves the "weather" service
// "delete" endpoint.
func NewDeleteHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeDeleteRequest, EncodeDeleteResponse)
	}
	return h
}

// Delete implements the "Delete" method in weatherpb.WeatherServer interface.
func (s *Server) Delete(ctx context.Context, message *weatherpb.DeleteRequest) (*weatherpb.DeleteResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "delete")
	ctx = context.WithValue(ctx, goa.ServiceKey, "weather")
	resp, err := s.DeleteH.Handle(ctx, message)
	if err != nil {
		if en, ok := err.(ErrorNamer); ok {
			switch en.ErrorName() {
			case "not_found":
				return nil, goagrpc.NewStatusError(codes.NotFound, err, goagrpc.NewErrorResponse(err))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*weatherpb.DeleteResponse), nil
}

// NewListHandler creates a gRPC handler which serves the "weather" service
// "list" endpoint.
func NewListHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeListRequest, EncodeListResponse)
	}
	return h
}

// List implements the "List" method in weatherpb.WeatherServer interface.
func (s *Server) List(ctx context.Context, message *weatherpb.ListRequest) (*weatherpb.ListResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "list")
	ctx = context.WithValue(ctx, goa.ServiceKey, "weather")
	resp, err := s.ListH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*weatherpb.ListResponse), nil
}
