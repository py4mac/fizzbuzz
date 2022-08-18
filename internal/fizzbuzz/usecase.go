//go:generate mockgen -source usecase.go -destination mock/usecase_mock.go -package mock
package fizzbuzz

import (
	"context"

	"github.com/py4mac/fizzbuzz/internal/fizzbuzz/domain"
)

// UseCase holds use cases interface
type UseCase interface {
	Record(ctx context.Context, e domain.Fizzbuz) (string, error)
	Process(ctx context.Context) (*domain.Statistics, error)
}
