// Package main is the applicative entry point
package main

import (
	"flag"
	"fmt"

	"github.com/py4mac/fizzbuzz/config"
	"github.com/py4mac/fizzbuzz/internal/server"
	"github.com/py4mac/fizzbuzz/pkg/constants"
	"github.com/py4mac/fizzbuzz/pkg/postgres"
	"github.com/py4mac/fizzbuzz/pkg/tracing"
	"github.com/py4mac/fizzbuzz/pkg/x/pflagx"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
)

// init predefined flag for passing configuration file
func init() {
	flag.StringVar(&config.CfgFile, "config", pflagx.LookupEnvOrString("CONFIG", "/app/config.yaml"), "Fizzbuzz microservice config path")
}

// @title Go Fizzbuzz REST API
// @version 1.0
// @description Fizzbuzz REST API
// @contact.name Pierre-Yves BOISBUNON
// @contact.url https://github.com/py4mac
// @contact.email pierreyves.boisbunon@gmail.com
// @BasePath /api/v1
func main() {
	flag.Parse()
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)

	logger := log.WithFields(log.Fields{
		"module": "cmd",
	})

	logger.Info(
		"serve fizzbuzz",
		" Version:", constants.Version,
		" Built:", constants.Built,
		" Revision:", constants.Revision,
	)

	cfg, err := config.InitConfig()
	if err != nil {
		logger.Error(fmt.Sprintf("Error loading configuration file: %s", err.Error()))
		panic(err)
	}

	psqlDB, err := postgres.NewPsqlDB(cfg)
	if err != nil {
		logger.Error(fmt.Sprintf("Postgresql init: %s", err.Error()))
		panic(err)
	} else {
		logger.Infof("Postgres connected, Status: %#v", psqlDB.Stats())
	}
	defer psqlDB.Close()

	tp, err := tracing.NewTracing(cfg)
	if err != nil {
		log.Error(fmt.Sprintf("cannot create tracer %s", err.Error()))
		panic(err)
	}

	logger.Info("Jaeger connected")

	otel.SetTracerProvider(tp)

	logger.Info("Otel connected")

	s := server.NewServer(cfg, psqlDB)
	if err = s.Run(); err != nil {
		log.Error(fmt.Sprintf("server error %s", err.Error()))
		panic(err)
	}
}
