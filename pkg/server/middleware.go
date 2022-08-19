package server

import (
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

// metricsMiddleware return middleware handler func
func metricsMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			err := next(c)

			var status int

			if err != nil {
				status = err.(*echo.HTTPError).Code
			} else {
				status = c.Response().Status
			}

			// Response time
			times.WithLabelValues(strconv.Itoa(status), c.Request().Method, c.Path()).Observe(time.Since(start).Seconds())
			// Counter
			hitsTotal.Inc()
			hits.WithLabelValues(strconv.Itoa(status), c.Request().Method, c.Path()).Inc()

			return err
		}
	}
}
