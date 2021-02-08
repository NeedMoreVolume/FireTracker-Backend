package main

import (
	"context"
	"github.com/rs/zerolog"
	"net"
	"net/url"
	"sync"

	"github.com/NeedMoreVolume/FireTracker/gen/fire"
	firepb "github.com/NeedMoreVolume/FireTracker/gen/grpc/fire/pb"
	firesvr "github.com/NeedMoreVolume/FireTracker/gen/grpc/fire/server"
	logpb "github.com/NeedMoreVolume/FireTracker/gen/grpc/log/pb"
	logsvr "github.com/NeedMoreVolume/FireTracker/gen/grpc/log/server"
	weatherpb "github.com/NeedMoreVolume/FireTracker/gen/grpc/weather/pb"
	weathersvr "github.com/NeedMoreVolume/FireTracker/gen/grpc/weather/server"
	"github.com/NeedMoreVolume/FireTracker/gen/log"
	"github.com/NeedMoreVolume/FireTracker/gen/weather"
	"github.com/NeedMoreVolume/FireTracker/middleware"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcmdlwr "goa.design/goa/v3/grpc/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// handleGRPCServer starts configures and starts a gRPC server on the given
// URL. It shuts down the server if any error is received in the error channel.
func handleGRPCServer(ctx context.Context, u *url.URL, fireEndpoints *fire.Endpoints, weatherEndpoints *weather.Endpoints, logEndpoints *log.Endpoints, wg *sync.WaitGroup, errc chan error, logger zerolog.Logger, debug bool) {

	// Setup goa log adapter.
	var (
		adapter middleware.Logger
	)
	{
		adapter = middleware.NewLogger(logger)
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to gRPC requests and
	// responses.
	var (
		fireServer *firesvr.Server
		weatherServer *weathersvr.Server
		logServer *logsvr.Server
	)
	{
		fireServer = firesvr.New(fireEndpoints, nil)
		weatherServer = weathersvr.New(weatherEndpoints, nil)
		logServer = logsvr.New(logEndpoints, nil)
	}

	// Initialize gRPC server with the middleware.
	srv := grpc.NewServer(
		grpcmiddleware.WithUnaryServerChain(
			grpcmdlwr.UnaryRequestID(),
			grpcmdlwr.UnaryServerLog(adapter),
		),
	)

	// Register the servers.
	firepb.RegisterFireServer(srv, fireServer)
	weatherpb.RegisterWeatherServer(srv, weatherServer)
	logpb.RegisterLogServer(srv, logServer)

	for svc, info := range srv.GetServiceInfo() {
		for _, m := range info.Methods {
			logger.Debug().Msgf("serving gRPC method %s", svc+"/"+m.Name)
		}
	}

	// Register the server reflection service on the server.
	// See https://grpc.github.io/grpc/core/md_doc_server-reflection.html.
	reflection.Register(srv)

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start gRPC server in a separate goroutine.
		go func() {
			lis, err := net.Listen("tcp", u.Host)
			if err != nil {
				errc <- err
			}
			logger.Info().Msgf("gRPC server listening on %q", u.Host)
			errc <- srv.Serve(lis)
		}()

		<-ctx.Done()
		logger.Info().Msgf("shutting down gRPC server at %q", u.Host)
		srv.Stop()
	}()
}
