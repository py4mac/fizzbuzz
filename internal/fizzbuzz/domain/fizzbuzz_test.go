package domain

import (
	"fmt"
	"testing"
)

func TestValidate(t *testing.T) {
	var tests = []struct {
		in   *Fizzbuz
		want error
	}{
		{&Fizzbuz{-1, -1, 0, "", ""}, ErrFizzbuzzIntsMustBePositive},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("SumNumbers=%d", i), func(t *testing.T) {
			got := tc.in.validate()
			if got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			} else {
				t.Logf("Success !")
			}
		})
	}
}
