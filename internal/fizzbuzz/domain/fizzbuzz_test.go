package domain

import (
	"context"
	"fmt"
	"testing"
)

func TestValidate(t *testing.T) {
	var tests = []struct {
		in   *Fizzbuz
		want error
	}{
		{&Fizzbuz{-1, -1, 10, "fizz", "buzz"}, ErrFizzbuzzIntsMustBePositive},
		{&Fizzbuz{2, 1, 10, "fizz", "buzz"}, ErrFizzbuzzInt2MustBeHigherThanInt1},
		{&Fizzbuz{1, 2, 0, "fizz", "buzz"}, ErrFizzbuzzLimitMustBePositive},
		{&Fizzbuz{1, 2, 101, "fizz", "buzz"}, ErrFizzbuzzLimitExceeded},
		{&Fizzbuz{1, 2, 10, "fizz", "buzz"}, nil},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := tc.in.validate()
			if got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			} else {
				t.Logf("Success !")
			}
		})
	}
}

func TestProcess(t *testing.T) {
	var tests = []struct {
		in   *Fizzbuz
		want string
	}{
		{&Fizzbuz{-1, -1, 10, "fizz", "buzz"}, ""},
		{&Fizzbuz{3, 5, 10, "fizz", "buzz"}, "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz"},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got, _ := tc.in.Process(context.Background())
			if got != tc.want {
				t.Fatalf("got %v; want %v", got, tc.want)
			} else {
				t.Logf("Success !")
			}
		})
	}
}
