package stats

import "github.com/py4mac/fizzbuzz/pkg/fizzbuzz"

type Stats interface {
	Record(e fizzbuzz.Fizzbuz) error
	Process() (Statistics, error)
}
