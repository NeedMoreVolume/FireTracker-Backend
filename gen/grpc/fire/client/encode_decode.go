// Code generated by goa v3.3.1, DO NOT EDIT.
//
// fire gRPC client encoders and decoders
//
// Command:
// $ goa gen github.com/NeedMoreVolume/FireTracker/design

package client

import (
	"context"

	fire "github.com/NeedMoreVolume/FireTracker/gen/fire"
	firepb "github.com/NeedMoreVolume/FireTracker/gen/grpc/fire/pb"
	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// BuildCreateFunc builds the remote method to invoke for "fire" service
// "create" endpoint.
func BuildCreateFunc(grpccli firepb.FireClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Create(ctx, reqpb.(*firepb.CreateRequest), opts...)
		}
		return grpccli.Create(ctx, &firepb.CreateRequest{}, opts...)
	}
}

// EncodeCreateRequest encodes requests sent to fire create endpoint.
func EncodeCreateRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*fire.Fire)
	if !ok {
		return nil, goagrpc.ErrInvalidType("fire", "create", "*fire.Fire", v)
	}
	return NewCreateRequest(payload), nil
}

// DecodeCreateResponse decodes responses from the fire create endpoint.
func DecodeCreateResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*firepb.CreateResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("fire", "create", "*firepb.CreateResponse", v)
	}
	if err := ValidateCreateResponse(message); err != nil {
		return nil, err
	}
	res := NewCreateResult(message)
	return res, nil
}

// BuildGetFunc builds the remote method to invoke for "fire" service "get"
// endpoint.
func BuildGetFunc(grpccli firepb.FireClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Get(ctx, reqpb.(*firepb.GetRequest), opts...)
		}
		return grpccli.Get(ctx, &firepb.GetRequest{}, opts...)
	}
}

// EncodeGetRequest encodes requests sent to fire get endpoint.
func EncodeGetRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*fire.GetPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("fire", "get", "*fire.GetPayload", v)
	}
	return NewGetRequest(payload), nil
}

// DecodeGetResponse decodes responses from the fire get endpoint.
func DecodeGetResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*firepb.GetResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("fire", "get", "*firepb.GetResponse", v)
	}
	if err := ValidateGetResponse(message); err != nil {
		return nil, err
	}
	res := NewGetResult(message)
	return res, nil
}

// BuildUpdateFunc builds the remote method to invoke for "fire" service
// "update" endpoint.
func BuildUpdateFunc(grpccli firepb.FireClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Update(ctx, reqpb.(*firepb.UpdateRequest), opts...)
		}
		return grpccli.Update(ctx, &firepb.UpdateRequest{}, opts...)
	}
}

// EncodeUpdateRequest encodes requests sent to fire update endpoint.
func EncodeUpdateRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*fire.Fire)
	if !ok {
		return nil, goagrpc.ErrInvalidType("fire", "update", "*fire.Fire", v)
	}
	return NewUpdateRequest(payload), nil
}

// DecodeUpdateResponse decodes responses from the fire update endpoint.
func DecodeUpdateResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*firepb.UpdateResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("fire", "update", "*firepb.UpdateResponse", v)
	}
	if err := ValidateUpdateResponse(message); err != nil {
		return nil, err
	}
	res := NewUpdateResult(message)
	return res, nil
}

// BuildDeleteFunc builds the remote method to invoke for "fire" service
// "delete" endpoint.
func BuildDeleteFunc(grpccli firepb.FireClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Delete(ctx, reqpb.(*firepb.DeleteRequest), opts...)
		}
		return grpccli.Delete(ctx, &firepb.DeleteRequest{}, opts...)
	}
}

// EncodeDeleteRequest encodes requests sent to fire delete endpoint.
func EncodeDeleteRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*fire.DeletePayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("fire", "delete", "*fire.DeletePayload", v)
	}
	return NewDeleteRequest(payload), nil
}

// BuildListFunc builds the remote method to invoke for "fire" service "list"
// endpoint.
func BuildListFunc(grpccli firepb.FireClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.List(ctx, reqpb.(*firepb.ListRequest), opts...)
		}
		return grpccli.List(ctx, &firepb.ListRequest{}, opts...)
	}
}

// EncodeListRequest encodes requests sent to fire list endpoint.
func EncodeListRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*fire.FireListPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("fire", "list", "*fire.FireListPayload", v)
	}
	return NewListRequest(payload), nil
}

// DecodeListResponse decodes responses from the fire list endpoint.
func DecodeListResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*firepb.ListResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("fire", "list", "*firepb.ListResponse", v)
	}
	if err := ValidateListResponse(message); err != nil {
		return nil, err
	}
	res := NewListResult(message)
	return res, nil
}

// BuildGetWeatherForFireFunc builds the remote method to invoke for "fire"
// service "getWeatherForFire" endpoint.
func BuildGetWeatherForFireFunc(grpccli firepb.FireClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.GetWeatherForFire(ctx, reqpb.(*firepb.GetWeatherForFireRequest), opts...)
		}
		return grpccli.GetWeatherForFire(ctx, &firepb.GetWeatherForFireRequest{}, opts...)
	}
}

// EncodeGetWeatherForFireRequest encodes requests sent to fire
// getWeatherForFire endpoint.
func EncodeGetWeatherForFireRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*fire.GetWeatherForFirePayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("fire", "getWeatherForFire", "*fire.GetWeatherForFirePayload", v)
	}
	return NewGetWeatherForFireRequest(payload), nil
}

// DecodeGetWeatherForFireResponse decodes responses from the fire
// getWeatherForFire endpoint.
func DecodeGetWeatherForFireResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*firepb.GetWeatherForFireResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("fire", "getWeatherForFire", "*firepb.GetWeatherForFireResponse", v)
	}
	if err := ValidateGetWeatherForFireResponse(message); err != nil {
		return nil, err
	}
	res := NewGetWeatherForFireResult(message)
	return res, nil
}

// BuildGetLogsForFireFunc builds the remote method to invoke for "fire"
// service "getLogsForFire" endpoint.
func BuildGetLogsForFireFunc(grpccli firepb.FireClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.GetLogsForFire(ctx, reqpb.(*firepb.GetLogsForFireRequest), opts...)
		}
		return grpccli.GetLogsForFire(ctx, &firepb.GetLogsForFireRequest{}, opts...)
	}
}

// EncodeGetLogsForFireRequest encodes requests sent to fire getLogsForFire
// endpoint.
func EncodeGetLogsForFireRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*fire.GetLogsForFirePayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("fire", "getLogsForFire", "*fire.GetLogsForFirePayload", v)
	}
	return NewGetLogsForFireRequest(payload), nil
}

// DecodeGetLogsForFireResponse decodes responses from the fire getLogsForFire
// endpoint.
func DecodeGetLogsForFireResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*firepb.GetLogsForFireResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("fire", "getLogsForFire", "*firepb.GetLogsForFireResponse", v)
	}
	if err := ValidateGetLogsForFireResponse(message); err != nil {
		return nil, err
	}
	res := NewGetLogsForFireResult(message)
	return res, nil
}
