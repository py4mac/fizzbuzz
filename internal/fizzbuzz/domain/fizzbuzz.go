package domain

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/py4mac/fizzbuzz/pkg/x/errorx"
)

const (
	MaxLimit = 100
)

var (
	ErrFizzbuzzInt1MustBePositive       = errors.New("int1 must be positive")
	ErrFizzbuzzInt2MustBePositive       = errors.New("int2 must be positive")
	ErrFizzbuzzInt2MustBeHigherThanInt1 = errors.New("int2 must be higher than int1")
	ErrFizzbuzzLimitMustBePositive      = errors.New("limit must be positive")
	ErrFizzbuzzLimitExceeded            = fmt.Errorf("limit must be below or equal %d", MaxLimit)
)

// FizzbuzReq holds query parameters for the rest endpoint
type Fizzbuz struct {
	Int1  int    `query:"int1" json:"int1"`
	Int2  int    `query:"int2" json:"int2"`
	Limit int    `query:"limit" json:"limit"`
	Str1  string `query:"str1" json:"str1"`
	Str2  string `query:"str2" json:"str2"`
}

// validate fizzbuzz struct fields
func (f *Fizzbuz) validate() error {
	if f.Int1 <= 0 {
		return ErrFizzbuzzInt1MustBePositive
	}

	if f.Int2 <= 0 {
		return ErrFizzbuzzInt2MustBePositive
	}

	if f.Int2 <= f.Int1 {
		return ErrFizzbuzzInt2MustBeHigherThanInt1
	}

	if f.Limit <= 0 {
		return ErrFizzbuzzLimitMustBePositive
	}

	if f.Limit > MaxLimit {
		return ErrFizzbuzzLimitExceeded
	}

	return nil
}

// Process serialize fizzbuzz struct
func (f *Fizzbuz) Process(ctx context.Context) (string, error) {
	if err := f.validate(); err != nil {
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
