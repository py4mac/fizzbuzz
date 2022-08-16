package fizzbuzz

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFizzbuzzIntsMustBePositive  = errors.New("int1 and int2 must be positives")
	ErrFizzbuzzLimitMustBePositive = errors.New("limit must be positive")
)

// FizzbuzReq are input parameters for the rest api endpoint
type Fizzbuz struct {
	// Int1  int    `form:"int1" json:"int1" binding:"required,gte=1,lte=100,gtfield=Int2"` /*  */
	Int1  int    `form:"int1" json:"int1"`
	Int2  int    `form:"int2" json:"int2"`
	Limit int    `form:"limit" json:"limit"`
	Str1  string `form:"str1" json:"str1"`
	Str2  string `form:"str2" json:"str2"`
}

func (f *Fizzbuz) validate() error {
	if f.Int1 < 0 || f.Int2 < 0 {
		return ErrFizzbuzzIntsMustBePositive
	}

	if f.Limit < 0 {
		return ErrFizzbuzzLimitMustBePositive
	}

	return nil
}

func (f *Fizzbuz) Process(ctx context.Context) ([]string, error) {
	if err := f.validate(); err != nil {
		return nil, err
	}

	response := make([]string, f.Limit)

	for i := 0; i < f.Limit; i++ {
		number := i + 1

		if ctx.Err() != nil {
			return nil, ctx.Err()
		}

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

	return response, nil
}
