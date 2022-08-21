package domain

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/py4mac/fizzbuzz/pkg/x/errorx"
)

const (
	MaxLimit = 100
)

var (
	fbValidator *validator.Validate
)

func init() {
	fbValidator = validator.New()
}

// FizzbuzReq holds query parameters for the rest endpoint
type Fizzbuz struct {
	Int1  int    `query:"int1" json:"int1" validate:"gt=0"`
	Int2  int    `query:"int2" json:"int2" validate:"gt=0,gtfield=Int1"`
	Limit int    `query:"limit" json:"limit" validate:"gt=0,lte=100"`
	Str1  string `query:"str1" json:"str1" validate:"required"`
	Str2  string `query:"str2" json:"str2" validate:"required"`
}

func (f *Fizzbuz) validate(ctx context.Context) error {
	return fbValidator.StructCtx(ctx, f)
}

// Process serialize fizzbuzz struct
func (f *Fizzbuz) Process(ctx context.Context) (string, error) {
	if err := f.validate(ctx); err != nil {
		return "", errorx.Wrap(err, "validation error")
	}

	response := make([]string, f.Limit)

	for i := 0; i < f.Limit; i++ {
		number := i + 1

		switch {
		case (number%f.Int1 == 0) && (number%f.Int2 == 0):
			response[i] = f.Str1 + f.Str2
		case (number%f.Int1 == 0):
			response[i] = f.Str1
		case (number%f.Int2 == 0):
			response[i] = f.Str2
		default:
			response[i] = fmt.Sprintf("%d", number)
		}
	}

	return strings.Join(response, ","), nil
}
