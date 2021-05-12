// Code generated by goa v3.3.1, DO NOT EDIT.
//
// weather gRPC client
//
// Command:
// $ goa gen github.com/NeedMoreVolume/FireTracker/design

package client

import (
	"context"

	weatherpb "github.com/NeedMoreVolume/FireTracker/gen/grpc/weather/pb"
	goagrpc "goa.design/goa/v3/grpc"
	goapb "goa.design/goa/v3/grpc/pb"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli weatherpb.WeatherClient
	opts    []grpc.CallOption
}

// NewClient instantiates gRPC client for all the weather service servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: weatherpb.NewWeatherClient(cc),
		opts:    opts,
	}
}

// Create calls the "Create" function in weatherpb.WeatherClient interface.
func (c *Client) Create() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildCreateFunc(c.grpccli, c.opts...),
			EncodeCreateRequest,
			DecodeCreateResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault(err.Error())
			}
		}
		return res, nil
	}
}

// Get calls the "Get" function in weatherpb.WeatherClient interface.
func (c *Client) Get() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildGetFunc(c.grpccli, c.opts...),
			EncodeGetRequest,
			DecodeGetResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault(err.Error())
			}
		}
		return res, nil
	}
}

// Update calls the "Update" function in weatherpb.WeatherClient interface.
func (c *Client) Update() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildUpdateFunc(c.grpccli, c.opts...),
			EncodeUpdateRequest,
			DecodeUpdateResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault(err.Error())
			}
		}
		return res, nil
	}
}

// Delete calls the "Delete" function in weatherpb.WeatherClient interface.
func (c *Client) Delete() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildDeleteFunc(c.grpccli, c.opts...),
			EncodeDeleteRequest,
			nil)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault(err.Error())
			}
		}
		return res, nil
	}
}

// List calls the "List" function in weatherpb.WeatherClient interface.
func (c *Client) List() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildListFunc(c.grpccli, c.opts...),
			EncodeListRequest,
			DecodeListResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}
