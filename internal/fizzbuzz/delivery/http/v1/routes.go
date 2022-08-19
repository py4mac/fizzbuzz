package v1

import (
	"github.com/labstack/echo/v4"

	"github.com/py4mac/fizzbuzz/internal/fizzbuzz"
)

// MapV1Routes maps fizzbuzz routes to their corresponding handlers
func MapV1Routes(v1Group *echo.Group, h fizzbuzz.Handlers) {
	v1Group.GET("/fizzbuzz", h.Record())
	v1Group.GET("/stats", h.Process())
}
