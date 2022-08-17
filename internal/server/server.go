package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"

	"github.com/py4mac/fizzbuzz/config"
	_ "github.com/py4mac/fizzbuzz/docs" //nolint:golint
)

const (
	maxHeaderBytes = 1 << 20
	ctxTimeout     = 5
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

func (s *Server) Run() error {
	// Disable echo banner
	s.echo.HideBanner = true

	server := &http.Server{
		Addr:           s.cfg.Server.Port,
		ReadTimeout:    time.Second * s.cfg.Server.ReadTimeout,
		WriteTimeout:   time.Second * s.cfg.Server.WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	go func() {
		logger.Infof("Server is listening on PORT: %s", s.cfg.Server.Port)

		if err := s.echo.StartServer(server); err != nil {
			logger.Fatalf(fmt.Sprintf("Error starting Server: %s", err.Error()))
		}
	}()

	if err := s.MapHandlers(s.echo); err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer shutdown()

	logger.Info("Server Exited Properly")

	return s.echo.Server.Shutdown(ctx)
}
