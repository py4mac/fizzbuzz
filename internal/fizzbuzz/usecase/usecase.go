package usecase

import (
	"context"

	"github.com/py4mac/fizzbuzz/internal/fizzbuzz"
	"go.opentelemetry.io/otel"

	"github.com/py4mac/fizzbuzz/internal/fizzbuzz/domain"
)

// fbUC UseCase
type fbUC struct {
	repo fizzbuzz.Repository
}

// NewFBUseCase repo constructor
func NewFBUseCase(repo fizzbuzz.Repository) fizzbuzz.UseCase {
	return &fbUC{repo: repo}
}

// Record fizzbuzz user entry inside persistent repository
func (u *fbUC) Record(ctx context.Context, e *domain.Fizzbuz) (string, error) {
	ctx, span := otel.Tracer("").Start(ctx, "fbUC.Record")
	defer span.End()

	return u.repo.Record(ctx, e)
}

// Process fizzbuzz statistics
func (u *fbUC) Process(ctx context.Context) (*domain.Statistics, error) {
	ctx, span := otel.Tracer("").Start(ctx, "fbUC.Process")
	defer span.End()

	return u.repo.Process(ctx)
}
