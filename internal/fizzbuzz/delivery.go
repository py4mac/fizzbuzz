package fizzbuzz

import "github.com/labstack/echo/v4"

// Handlers holds REST APIs interface
type Handlers interface {
	Record() echo.HandlerFunc
	Process() echo.HandlerFunc
}
