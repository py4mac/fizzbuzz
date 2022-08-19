package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"

	"github.com/py4mac/fizzbuzz/internal/fizzbuzz"
	"github.com/py4mac/fizzbuzz/internal/fizzbuzz/domain"
)

// v1Handlers handlers
type v1Handlers struct {
	uc fizzbuzz.UseCase
}

// NewV1Handlers is v1 handlers constructor
func NewV1Handlers(uc fizzbuzz.UseCase) fizzbuzz.Handlers {
	return &v1Handlers{uc: uc}
}

// Record godoc
// @Summary			Record
// @Description		Record fizzbuzz handler
// @Tags			fizzbuzz
// @Accept			json
// @Produce			json
// @Success			200 {object} string
// @Failure			400
// @Failure			500
// @Param			int1 query int true "Integer1" 1
// @Param			int2 query int true "Integer2" 1
// @Param			limit query int true "Limit" 100
// @Param			str1 query string true "String1" "fizz"
// @Param			str2 query string true "String2" "buzz"
// @Router			/fizzbuzz [get]
func (h v1Handlers) Record() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, span := otel.Tracer("").Start(c.Request().Context(), "v1Handlers.Record")
		defer span.End()

		n := new(domain.Fizzbuz)
		if err := c.Bind(n); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		response, err := h.uc.Record(ctx, n)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, response)
	}
}

// Process godoc
// @Summary			Process
// @Description		Process status handler
// @Tags			fizzbuzz
// @Produce			json
// @Success			200 {object} domain.Statistics
// @Failure			500
// @Router			/stats [get]
func (h v1Handlers) Process() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response *domain.Statistics

		ctx, span := otel.Tracer("").Start(c.Request().Context(), "v1Handlers.Process")
		defer span.End()

		response, err := h.uc.Process(ctx)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, response)
	}
}
