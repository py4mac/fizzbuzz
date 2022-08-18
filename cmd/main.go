// Package main is the applicative entry point
package main

import (
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/py4mac/fizzbuzz/config"
	"github.com/py4mac/fizzbuzz/internal/server"
	"github.com/py4mac/fizzbuzz/pkg/constants"
	"github.com/py4mac/fizzbuzz/pkg/postgres"
	log "github.com/sirupsen/logrus"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
)

// @title Go Fizzbuzz REST API
// @version 1.0
// @description Fizzbuzz REST API
// @contact.name Pierre-Yves BOISBUNON
// @contact.url https://github.com/py4mac
// @contact.email pierreyves.boisbunon@gmail.com
// @BasePath /api/v1
func main() {
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

	jaegerCfgInstance := jaegercfg.Configuration{
		ServiceName: cfg.Jaeger.ServiceName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           false,
			LocalAgentHostPort: cfg.Jaeger.Host,
		},
	}

	tracer, closer, err := jaegerCfgInstance.NewTracer(
		jaegercfg.Logger(jaegerlog.StdLogger),
		jaegercfg.Metrics(metrics.NullFactory),
	)
	if err != nil {
		log.Error(fmt.Sprintf("cannot create tracer %s", err.Error()))
		panic(err)
	}

	logger.Info("Jaeger connected")

	opentracing.SetGlobalTracer(tracer)

	defer closer.Close()

	logger.Info("Opentracing connected")

	s := server.NewServer(cfg, psqlDB)
	if err = s.Run(); err != nil {
		log.Error(fmt.Sprintf("server error %s", err.Error()))
		panic(err)
	}
}
