package middleware

import (
	"github.com/py4mac/fizzbuzz/config"
)

// Middleware manager
type Manager struct {
	cfg *config.Config
}

// NewMiddlewareManager create a new middleware manager
func NewMiddlewareManager(cfg *config.Config) *Manager {
	return &Manager{cfg: cfg}
}
