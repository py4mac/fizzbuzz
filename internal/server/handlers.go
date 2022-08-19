package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/py4mac/fizzbuzz/docs"
	v1 "github.com/py4mac/fizzbuzz/internal/fizzbuzz/delivery/http/v1"
	repository "github.com/py4mac/fizzbuzz/internal/fizzbuzz/repository"
	useCase "github.com/py4mac/fizzbuzz/internal/fizzbuzz/usecase"
	apiMiddlewares "github.com/py4mac/fizzbuzz/internal/middleware"
	"github.com/py4mac/fizzbuzz/pkg/metric"
)

// Map Server Handlers
func (s *Server) MapHandlers(e *echo.Echo) error {
	metrics, err := metric.NewPrometheusMetric(s.cfg.Metrics.URL, s.cfg.Metrics.ServiceName)
	if err != nil {
		logger.Errorf("CreateMetrics Error: %s", err)
	}

	logger.Info(
		fmt.Sprintf("Metrics available URL: %s, ServiceName: %s",
			s.cfg.Metrics.URL,
			s.cfg.Metrics.ServiceName),
	)

	// Init repositories
	fbRepo := repository.NewFBInPg(s.db)

	// Init useCases
	fbUC := useCase.NewFBUseCase(fbRepo)

	// Init handlers
	v1Handlers := v1.NewV1Handlers(fbUC)

	mw := apiMiddlewares.NewMiddlewareManager(s.cfg)

	docs.SwaggerInfo.Title = "Fizzbuzz REST API"

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
	})

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.Use(mw.MetricsMiddleware(metrics))

	apiV1 := e.Group("/api/v1")

	v1.MapFbRoutes(apiV1, v1Handlers)

	return nil
}
