package http

import (
	"github.com/labstack/echo/v4"

	"github.com/py4mac/fizzbuzz/internal/fizzbuzz"
)

// MapRoutes maps fizzbuzz routes to their corresponding handlers
func MapRoutes(g *echo.Group, h fizzbuzz.Handlers) {
	g.GET("/fizzbuzz", h.Record())
	g.GET("/stats", h.Process())
}
