package v1

import (
	"github.com/labstack/echo/v4"

	"github.com/py4mac/fizzbuzz/internal/fizzbuzz"
)

// MapFbRoutes maps fizzbuzz routes
func MapFbRoutes(newsGroup *echo.Group, h fizzbuzz.Handlers) {
	newsGroup.GET("/fizzbuzz", h.Record())
	newsGroup.GET("/stats", h.Process())
}
