package domain

import (
	"context"
	"fmt"
	"testing"
)

func TestValidate(t *testing.T) {
	var tests = []struct {
		in      *Fizzbuz
		wantErr bool
	}{
		{&Fizzbuz{-1, -1, 10, "fizz", "buzz"}, true},
		{&Fizzbuz{1, -1, 10, "fizz", "buzz"}, true},
		{&Fizzbuz{2, 1, 10, "fizz", "buzz"}, true},
		{&Fizzbuz{1, 2, 0, "fizz", "buzz"}, true},
		{&Fizzbuz{1, 2, 101, "fizz", "buzz"}, true},
		{&Fizzbuz{1, 2, 10, "fizz", "buzz"}, false},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := tc.in.validate(context.Background())
			if tc.wantErr != (got != nil) {
				t.Fatalf("got error %v; but expected no error", got)
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
