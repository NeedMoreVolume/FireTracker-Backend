// Code generated by goa v3.2.5, DO NOT EDIT.
//
// fire gRPC server
//
// Command:
// $ goa gen github.com/NeedMoreVolume/FireTracker/design

package server

import (
	"context"

	fire "github.com/NeedMoreVolume/FireTracker/gen/fire"
	firepb "github.com/NeedMoreVolume/FireTracker/gen/grpc/fire/pb"
	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc/codes"
)

// Server implements the firepb.FireServer interface.
type Server struct {
	CreateH            goagrpc.UnaryHandler
	GetH               goagrpc.UnaryHandler
	UpdateH            goagrpc.UnaryHandler
	DeleteH            goagrpc.UnaryHandler
	ListH              goagrpc.UnaryHandler
	GetWeatherForFireH goagrpc.UnaryHandler
	GetLogsForFireH    goagrpc.UnaryHandler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the expr.
type ErrorNamer interface {
	ErrorName() string
}

// New instantiates the server struct with the fire service endpoints.
func New(e *fire.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		CreateH:            NewCreateHandler(e.Create, uh),
		GetH:               NewGetHandler(e.Get, uh),
		UpdateH:            NewUpdateHandler(e.Update, uh),
		DeleteH:            NewDeleteHandler(e.Delete, uh),
		ListH:              NewListHandler(e.List, uh),
		GetWeatherForFireH: NewGetWeatherForFireHandler(e.GetWeatherForFire, uh),
		GetLogsForFireH:    NewGetLogsForFireHandler(e.GetLogsForFire, uh),
	}
}

// NewCreateHandler creates a gRPC handler which serves the "fire" service
// "create" endpoint.
func NewCreateHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeCreateRequest, EncodeCreateResponse)
	}
	return h
}

// Create implements the "Create" method in firepb.FireServer interface.
func (s *Server) Create(ctx context.Context, message *firepb.CreateRequest) (*firepb.CreateResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "create")
	ctx = context.WithValue(ctx, goa.ServiceKey, "fire")
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
	return resp.(*firepb.CreateResponse), nil
}

// NewGetHandler creates a gRPC handler which serves the "fire" service "get"
// endpoint.
func NewGetHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeGetRequest, EncodeGetResponse)
	}
	return h
}

// Get implements the "Get" method in firepb.FireServer interface.
func (s *Server) Get(ctx context.Context, message *firepb.GetRequest) (*firepb.GetResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "get")
	ctx = context.WithValue(ctx, goa.ServiceKey, "fire")
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
	return resp.(*firepb.GetResponse), nil
}

// NewUpdateHandler creates a gRPC handler which serves the "fire" service
// "update" endpoint.
func NewUpdateHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeUpdateRequest, EncodeUpdateResponse)
	}
	return h
}

// Update implements the "Update" method in firepb.FireServer interface.
func (s *Server) Update(ctx context.Context, message *firepb.UpdateRequest) (*firepb.UpdateResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "update")
	ctx = context.WithValue(ctx, goa.ServiceKey, "fire")
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
	return resp.(*firepb.UpdateResponse), nil
}

// NewDeleteHandler creates a gRPC handler which serves the "fire" service
// "delete" endpoint.
func NewDeleteHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeDeleteRequest, EncodeDeleteResponse)
	}
	return h
}

// Delete implements the "Delete" method in firepb.FireServer interface.
func (s *Server) Delete(ctx context.Context, message *firepb.DeleteRequest) (*firepb.DeleteResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "delete")
	ctx = context.WithValue(ctx, goa.ServiceKey, "fire")
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
	return resp.(*firepb.DeleteResponse), nil
}

// NewListHandler creates a gRPC handler which serves the "fire" service "list"
// endpoint.
func NewListHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeListRequest, EncodeListResponse)
	}
	return h
}

// List implements the "List" method in firepb.FireServer interface.
func (s *Server) List(ctx context.Context, message *firepb.ListRequest) (*firepb.ListResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "list")
	ctx = context.WithValue(ctx, goa.ServiceKey, "fire")
	resp, err := s.ListH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*firepb.ListResponse), nil
}

// NewGetWeatherForFireHandler creates a gRPC handler which serves the "fire"
// service "getWeatherForFire" endpoint.
func NewGetWeatherForFireHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeGetWeatherForFireRequest, EncodeGetWeatherForFireResponse)
	}
	return h
}

// GetWeatherForFire implements the "GetWeatherForFire" method in
// firepb.FireServer interface.
func (s *Server) GetWeatherForFire(ctx context.Context, message *firepb.GetWeatherForFireRequest) (*firepb.GetWeatherForFireResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "getWeatherForFire")
	ctx = context.WithValue(ctx, goa.ServiceKey, "fire")
	resp, err := s.GetWeatherForFireH.Handle(ctx, message)
	if err != nil {
		if en, ok := err.(ErrorNamer); ok {
			switch en.ErrorName() {
			case "not_found":
				return nil, goagrpc.NewStatusError(codes.NotFound, err, goagrpc.NewErrorResponse(err))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*firepb.GetWeatherForFireResponse), nil
}

// NewGetLogsForFireHandler creates a gRPC handler which serves the "fire"
// service "getLogsForFire" endpoint.
func NewGetLogsForFireHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeGetLogsForFireRequest, EncodeGetLogsForFireResponse)
	}
	return h
}

// GetLogsForFire implements the "GetLogsForFire" method in firepb.FireServer
// interface.
func (s *Server) GetLogsForFire(ctx context.Context, message *firepb.GetLogsForFireRequest) (*firepb.GetLogsForFireResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "getLogsForFire")
	ctx = context.WithValue(ctx, goa.ServiceKey, "fire")
	resp, err := s.GetLogsForFireH.Handle(ctx, message)
	if err != nil {
		if en, ok := err.(ErrorNamer); ok {
			switch en.ErrorName() {
			case "not_found":
				return nil, goagrpc.NewStatusError(codes.NotFound, err, goagrpc.NewErrorResponse(err))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*firepb.GetLogsForFireResponse), nil
}