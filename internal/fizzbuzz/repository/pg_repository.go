package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/py4mac/fizzbuzz/internal/fizzbuzz"
	"github.com/py4mac/fizzbuzz/internal/fizzbuzz/domain"
	"github.com/py4mac/fizzbuzz/pkg/x/errorx"
)

// fbInPg holds internal structure for postgres repository
type fbInPg struct {
	db *sqlx.DB
}

// NewFBInPg return instance of postgres repository fiting Repository interface
func NewFBInPg(db *sqlx.DB) fizzbuzz.Repository {
	return &fbInPg{db: db}
}

// Record a fizzbuzz object into postgres repository
func (f *fbInPg) Record(ctx context.Context, e *domain.Fizzbuz) (string, error) {
	if _, err := f.db.ExecContext(
		ctx,
		insertRecord,
		e.Int1,
		e.Int2,
		e.Limit,
		e.Str1,
		e.Str2,
	); err != nil {
		return "", errorx.Wrap(err, "db record error")
	}

	return e.Process(ctx)
}

// Process compute statistic based on postgres db content
func (f *fbInPg) Process(ctx context.Context) (*domain.Statistics, error) {
	var (
		int1, int2, limit, count int
		str1, str2               string
	)

	if err := f.db.QueryRowContext(ctx, getStats).
		Scan(
			&int1,
			&int2,
			&limit,
			&str1,
			&str2,
			&count,
		); err != nil {
		return nil, errorx.Wrap(err, "db process error")
	}

	return &domain.Statistics{
		Hits: int32(count),
		Fizzbuz: &domain.Fizzbuz{
			Int1:  int1,
			Int2:  int2,
			Limit: limit,
			Str1:  str1,
			Str2:  str2,
		},
	}, nil
}
