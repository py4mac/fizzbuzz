package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"

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

// Create godoc
// @Summary Record
// @Description Record fizzbuzz handler
// @Tags v1
// @Accept json
// @Produce json
// @Success 200 {object} []string
// @Param int1 formData int true "Integer1" 1
// @Param int2 formData int true "Integer2" 1
// @Param limit formData int true "Limit" 100
// @Param str1 formData string true "String1" "fizz"
// @Param str2 formData string true "String2" "buzz"
// @Router /api/v1/fizzbuzz [get]
func (h v1Handlers) Record() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(c.Request().Context(), "v1Handlers.Record")
		defer span.Finish()

		n := new(domain.Fizzbuz)
		if err := c.Bind(n); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		response, err := h.uc.Record(ctx, *n)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, response)
	}
}

// Update godoc
// @Summary Process
// @Description Process status handler
// @Tags v1
// @Produce json
// @Success 200 {object} dto.Statistics
// @Router /api/v1/stats [get]
func (h v1Handlers) Process() echo.HandlerFunc {
	return func(c echo.Context) error {
		var response *domain.Statistics

		span, ctx := opentracing.StartSpanFromContext(c.Request().Context(), "v1Handlers.Process")
		defer span.Finish()

		response, err := h.uc.Process(ctx)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, response)
	}
}
