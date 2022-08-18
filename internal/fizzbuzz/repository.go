//go:generate mockgen -source repository.go -destination mock/pg_repository_mock.go -package mock
package fizzbuzz

import (
	"context"

	"github.com/py4mac/fizzbuzz/internal/fizzbuzz/domain"
)

// Repository holds repositories interface
type Repository interface {
	Record(ctx context.Context, e *domain.Fizzbuz) (string, error)
	Process(ctx context.Context) (*domain.Statistics, error)
}
