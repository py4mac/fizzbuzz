package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	server *http.Server
}

// NewServer create a new HTTP server instance
func NewServer(port string) *Server {
	server := &http.Server{
		Addr: port,
	}

	return &Server{
		server: server,
	}
}

func (s *Server) SetupHandlers(handlers *gin.Engine) {
	s.server.Handler = handlers
}

// Run the http server
func (s *Server) Run() error {
	return s.server.ListenAndServe()
}

// Stop the http server
func (s *Server) Stop() error {
	if s.server != nil {
		return s.server.Close()
	}

	return nil
}
