// Package main is the applicative entry point
package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/py4mac/fizzbuzz/config"
	delivery "github.com/py4mac/fizzbuzz/internal/fizzbuzz/delivery/http"
	v1 "github.com/py4mac/fizzbuzz/internal/fizzbuzz/delivery/http/v1"
	"github.com/py4mac/fizzbuzz/internal/fizzbuzz/repository"
	"github.com/py4mac/fizzbuzz/internal/fizzbuzz/usecase"
	"github.com/py4mac/fizzbuzz/pkg/constants"
	"github.com/py4mac/fizzbuzz/pkg/db/postgres"
	"github.com/py4mac/fizzbuzz/pkg/server"
	"github.com/py4mac/fizzbuzz/pkg/tracing"
	"github.com/py4mac/fizzbuzz/pkg/x/pflagx"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
)

var (
	ctxTimeout = 10
)

// init predefined flag for passing configuration file
func init() {
	flag.StringVar(&config.CfgFile, "config", pflagx.LookupEnvOrString("CONFIG", "/app/config.yaml"), "Fizzbuzz microservice config path")
}

// @title 		Go Fizzbuzz REST API
// @version 	1.0
// @description Fizzbuzz REST API
// @BasePath  	/api/v1
func main() {
	flag.Parse()
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)

	logger := log.WithFields(log.Fields{
		"module": "cmd",
	})

	logger.Info(
		"Fizzbuzz",
		" Version:", constants.Version,
		" Built:", constants.Built,
		" Revision:", constants.Revision,
	)

	cfg, err := config.InitConfig()
	if err != nil {
		logger.Error(fmt.Sprintf("Error loading configuration file: %s", err.Error()))
		panic(err)
	}

	pgClient, err := postgres.NewPgClient(cfg)
	if err != nil {
		logger.Error(fmt.Sprintf("Postgresql init: %s", err.Error()))
		panic(err)
	} else {
		logger.Infof("Postgres connected")
	}
	defer pgClient.Close()

	// Init repositories
	fbRepo := repository.NewFBInPg(pgClient)

	// Init useCases
	fbUC := usecase.NewFBUseCase(fbRepo)

	// Tracing
	tp, err := tracing.NewTracing(cfg)
	if err != nil {
		logger.Error(fmt.Sprintf("cannot create tracer %s", err.Error()))
		panic(err)
	}

	logger.Info("Jaeger connected")

	otel.SetTracerProvider(tp)

	logger.Info("Otel connected")

	// Server
	s := server.NewServer(cfg, pgClient)

	// Init handlers
	v1Handlers := v1.NewV1Handlers(fbUC)
	apiV1 := s.NewGroup("/api/v1")
	delivery.MapRoutes(apiV1, v1Handlers)

	go func() {
		if err = s.Run(); err != nil {
			if err == http.ErrServerClosed {
				logger.Info("Shuting down server")
			} else {
				logger.Error(fmt.Sprintf("server error %s", err.Error()))
				panic(err)
			}
		}
	}()

	logger.Printf("Server is running on port: %s", cfg.Server.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(ctxTimeout)*time.Second)
	defer cancel()

	logger.Info("Server Exited Properly")

	_ = s.Stop(ctx)
}
