package main

import (
	"context"
	"github.com/rs/zerolog"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/NeedMoreVolume/FireTracker/gen/fire"
	firesvr "github.com/NeedMoreVolume/FireTracker/gen/http/fire/server"
	logsvr "github.com/NeedMoreVolume/FireTracker/gen/http/log/server"
	weathersvr "github.com/NeedMoreVolume/FireTracker/gen/http/weather/server"
	"github.com/NeedMoreVolume/FireTracker/gen/log"
	"github.com/NeedMoreVolume/FireTracker/gen/weather"
	zerologmdlwr "github.com/NeedMoreVolume/FireTracker/middleware"
	goahttp "goa.design/goa/v3/http"
	httpmdlwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"
)

// handleHTTPServer starts configures and starts a HTTP server on the given
// URL. It shuts down the server if any error is received in the error channel.
func handleHTTPServer(ctx context.Context, u *url.URL, fireEndpoints *fire.Endpoints, weatherEndpoints *weather.Endpoints, logEndpoints *log.Endpoints, wg *sync.WaitGroup, errc chan error, logger zerolog.Logger, debug bool) {

	// Setup goa log adapter.
	var (
		adapter zerologmdlwr.Logger
	)
	{
		adapter = zerologmdlwr.NewLogger(logger.With().Str("serverType", "HTTP").Logger())
	}

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/implement/encoding.
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	var mux goahttp.Muxer
	{
		mux = goahttp.NewMuxer()
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to HTTP requests and
	// responses.
	var (
		fireServer *firesvr.Server
		weatherServer *weathersvr.Server
		logServer *logsvr.Server
	)
	{
		eh := errorHandler(logger)
		fireServer = firesvr.New(fireEndpoints, mux, dec, enc, eh, nil)
		logServer = logsvr.New(logEndpoints, mux, dec, enc, eh, nil)
		weatherServer = weathersvr.New(weatherEndpoints, mux, dec, enc, eh, nil)
		if debug {
			servers := goahttp.Servers{
				fireServer,
				weatherServer,
				logServer,
			}
			servers.Use(httpmdlwr.Debug(mux, os.Stdout))
		}
	}

	// Configure the mux.
	firesvr.Mount(mux, fireServer)
	logsvr.Mount(mux, logServer)
	weathersvr.Mount(mux, weatherServer)

	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints.
	var handler http.Handler = mux
	{
		handler = httpmdlwr.Log(adapter)(handler)
		handler = httpmdlwr.RequestID()(handler)
	}

	// Start HTTP server using default configuration, change the code to
	// configure the server as required by your service.
	srv := &http.Server{Addr: u.Host, Handler: handler}
	for _, m := range fireServer.Mounts {
		logger.Debug().Msgf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern)
	}

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start HTTP server in a separate goroutine.
		go func() {
			logger.Info().Msgf("HTTP server listening on %q", u.Host)
			errc <- srv.ListenAndServe()
		}()

		<-ctx.Done()
		logger.Info().Msgf("shutting down HTTP server at %q", u.Host)

		// Shutdown gracefully with a 30s timeout.
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()
}

// errorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func errorHandler(logger zerolog.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		_, _ = w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		logger.Info().Err(err).Msgf("request %s resulted in error", id)
	}
}
