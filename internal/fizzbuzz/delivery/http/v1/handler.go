package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"

	"github.com/py4mac/fizzbuzz/internal/fizzbuzz"
	"github.com/py4mac/fizzbuzz/internal/fizzbuzz/domain"
	"github.com/py4mac/fizzbuzz/pkg/x/errorx"
)

// v1Handlers handlers
type v1Handlers struct {
	uc fizzbuzz.UseCase
}

// NewV1Handlers is v1 handlers constructor
func NewV1Handlers(uc fizzbuzz.UseCase) fizzbuzz.Handlers {
	return &v1Handlers{uc: uc}
}

// Fizzbuzz godoc
// @Summary			Record and compute request
// @Description		Record request inside persistent repository and returns a list of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2
// @Tags			Fizzbuzz
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
func (h *v1Handlers) Record() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx, span := otel.Tracer("").Start(c.Request().Context(), "v1Handlers.Record")
		defer span.End()

		n := new(domain.Fizzbuz)
		if err := c.Bind(n); err != nil {
			return c.JSON(http.StatusBadRequest, errorx.Wrap(err, "wrong parameters").Error())
		}

		response, err := h.uc.Record(ctx, n)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, response)
	}
}

// Stats godoc
// @Summary			Process the most frequent request
// @Description		Return the most used request, as well as the number of hits for this request
// @Tags			Fizzbuzz
// @Produce			json
// @Success			200 {object} domain.Statistics
// @Failure			500
// @Router			/stats [get]
func (h *v1Handlers) Process() echo.HandlerFunc {
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
