package server

import (
	"context"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/py4mac/fizzbuzz/config"
	_ "github.com/py4mac/fizzbuzz/docs"
)

// Server struct
type Server struct {
	echo *echo.Echo
	cfg  *config.Config
	db   *sqlx.DB
}

// NewServer New Server constructor
func NewServer(cfg *config.Config, db *sqlx.DB) *Server {
	return &Server{echo: echo.New(), cfg: cfg, db: db}
}

func (s *Server) NewGroup(group string) *echo.Group {
	return s.echo.Group(group)
}

func (s *Server) Run() error {
	// Disable echo banner
	s.echo.HideBanner = true

	// Setup Middlewares
	s.echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	s.echo.Use(metricsMiddleware())

	s.echo.GET("/swagger/*", echoSwagger.WrapHandler)

	// Setup Health Endpoint
	s.echo.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{})
	})

	// Run metrics
	go func() {
		router := echo.New()
		router.HideBanner = true
		router.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
		logger.Printf("Metrics server is running on port: %s", s.cfg.Metrics.Port)

		if err := router.Start(s.cfg.Metrics.Port); err != nil {
			logger.Fatal(err)
		}
	}()

	return s.echo.Start(s.cfg.Server.Port)

}

// Stop the http server
func (s *Server) Stop(ctx context.Context) error {
	if s.echo != nil {
		return s.echo.Shutdown(ctx)
	}
	return nil
}
