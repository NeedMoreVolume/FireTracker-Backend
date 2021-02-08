package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/NeedMoreVolume/FireTracker/controllers"
	"github.com/NeedMoreVolume/FireTracker/services"
	"github.com/rs/zerolog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/NeedMoreVolume/FireTracker/gen/fire"
	"github.com/NeedMoreVolume/FireTracker/gen/log"
	"github.com/NeedMoreVolume/FireTracker/gen/weather"
)

func main() {
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		hostF     = flag.String("host", "development", "Server host (valid values: development, stage, production)")
		domainF   = flag.String("domain", "0.0.0.0", "Host domain name (overrides host domain specified in service design)")
		httpPortF = flag.String("http-port", "9010", "HTTP port (overrides host HTTP port specified in service design)")
		grpcPortF = flag.String("grpc-port", "9011", "gRPC port (overrides host gRPC port specified in service design)")
		secureF   = flag.Bool("secure", false, "Use secure scheme (https or grpcs)")
		dbgF      = flag.Bool("debug", false, "Log request and response bodies")
		sqlHost   = flag.String("sql-host", "127.0.0.1", "SQL host")
		sqlPort   = flag.String("sql-port", "3306", "SQL port")
		sqlUser   = flag.String("sql-user", "root", "SQL user")
		sqlPass   = flag.String("sql-pass", "R0ckets#1", "SQL pass")
		sqlSchema = flag.String("sql-schema", "fires", "SQL schema")
	)
	flag.Parse()

	// Setup logger. Replace logger with your own log package of choice.
	var (
		logger zerolog.Logger
	)
	{
		// default to info level logging
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		if *dbgF {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		}
		logger = zerolog.New(os.Stderr).With().Str("server", "firetracker").Timestamp().Logger()
	}

	// Setup DB.
	var (
		db *gorm.DB
	)
	{
		var err error
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime&parseTime=True&loc=Local", *sqlUser, *sqlPass, *sqlHost, *sqlPort, *sqlSchema)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to open db connection")
		}
	}

	// Initialize the services.
	var (
		fireSvc    fire.Service
		weatherSvc weather.Service
		logSvc     log.Service
	)
	{
		tempUnitService := services.NewTempUnit(logger, db)
		speedUnitService := services.NewSpeedUnit(logger, db)

		logService := services.NewLogService(logger, db)
		logSvc = controllers.NewLog(logger, logService)

		weatherService := services.NewWeatherService(logger, db)
		weatherSvc = controllers.NewWeather(logger, weatherService, tempUnitService, speedUnitService)

		fireService := services.NewFireService(logger, db)
		fireSvc = controllers.NewFire(logger, fireService, logService, weatherService)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		fireEndpoints    *fire.Endpoints
		weatherEndpoints *weather.Endpoints
		logEndpoints     *log.Endpoints
	)
	{
		fireEndpoints = fire.NewEndpoints(fireSvc)
		weatherEndpoints = weather.NewEndpoints(weatherSvc)
		logEndpoints = log.NewEndpoints(logSvc)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Start the servers and send errors (if any) to the error channel.
	switch *hostF {
	case "development":
		{
			addr := "http://0.0.0.0:8080"
			u, err := url.Parse(addr)
			if err != nil {
				logger.Error().Err(err).Msgf("invalid URL %s", addr)
				os.Exit(1)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h := strings.Split(u.Host, ":")[0]
				u.Host = h + ":" + *httpPortF
			} else if u.Port() == "" {
				u.Host += ":80"
			}
			handleHTTPServer(ctx, u, fireEndpoints, weatherEndpoints, logEndpoints, &wg, errc, logger, *dbgF)
		}

		{
			addr := "grpc://0.0.0.0:9080"
			u, err := url.Parse(addr)
			if err != nil {
				logger.Error().Err(err).Msgf("invalid URL %s", addr)
				os.Exit(1)
			}
			if *secureF {
				u.Scheme = "grpcs"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *grpcPortF != "" {
				h := strings.Split(u.Host, ":")[0]
				u.Host = h + ":" + *grpcPortF
			} else if u.Port() == "" {
				u.Host += ":8080"
			}
			handleGRPCServer(ctx, u, fireEndpoints, weatherEndpoints, logEndpoints, &wg, errc, logger, *dbgF)
		}

	case "stage":
		{
			addr := "https://test.tmf.com"
			u, err := url.Parse(addr)
			if err != nil {
				logger.Error().Err(err).Msgf("invalid URL %s", addr)
				os.Exit(1)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h := strings.Split(u.Host, ":")[0]
				u.Host = h + ":" + *httpPortF
			} else if u.Port() == "" {
				u.Host += ":443"
			}
			handleHTTPServer(ctx, u, fireEndpoints, weatherEndpoints, logEndpoints, &wg, errc, logger, *dbgF)
		}

		{
			addr := "grpc://test.tmf.com"
			u, err := url.Parse(addr)
			if err != nil {
				logger.Error().Err(err).Msgf("invalid URL %s", addr)
				os.Exit(1)
			}
			if *secureF {
				u.Scheme = "grpcs"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *grpcPortF != "" {
				h := strings.Split(u.Host, ":")[0]
				u.Host = h + ":" + *grpcPortF
			} else if u.Port() == "" {
				u.Host += ":8080"
			}
			handleGRPCServer(ctx, u, fireEndpoints, weatherEndpoints, logEndpoints, &wg, errc, logger, *dbgF)
		}

	case "production":
		{
			addr := "https://trackmyfire.com"
			u, err := url.Parse(addr)
			if err != nil {
				logger.Error().Err(err).Msgf("invalid URL %s", addr)
				os.Exit(1)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h := strings.Split(u.Host, ":")[0]
				u.Host = h + ":" + *httpPortF
			} else if u.Port() == "" {
				u.Host += ":443"
			}
			handleHTTPServer(ctx, u, fireEndpoints, weatherEndpoints, logEndpoints, &wg, errc, logger, *dbgF)
		}

		{
			addr := "grpc://trackmyfire.com"
			u, err := url.Parse(addr)
			if err != nil {
				logger.Error().Err(err).Msgf("invalid URL %s", addr)
				os.Exit(1)
			}
			if *secureF {
				u.Scheme = "grpcs"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *grpcPortF != "" {
				h := strings.Split(u.Host, ":")[0]
				u.Host = h + ":" + *grpcPortF
			} else if u.Port() == "" {
				u.Host += ":8080"
			}
			handleGRPCServer(ctx, u, fireEndpoints, weatherEndpoints, logEndpoints, &wg, errc, logger, *dbgF)
		}

	default:
		logger.Error().Msgf("invalid host argument: %q (valid hosts: development|stage|production)", *hostF)
	}

	// Wait for signal.
	logger.Info().Msgf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Info().Msg("exited")
}
