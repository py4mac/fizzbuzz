package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/py4mac/fizzbuzz/internal/fizzbuzz"
	"github.com/py4mac/fizzbuzz/internal/fizzbuzz/domain"
)

type fbInPg struct {
	db *sqlx.DB
}

func NewFBInPg(db *sqlx.DB) fizzbuzz.Repository {
	return &fbInPg{db: db}
}

func (f *fbInPg) Record(ctx context.Context, e domain.Fizzbuz) ([]string, error) {
	if _, err := f.db.ExecContext(
		ctx,
		insertRecord,
		e.Int1,
		e.Int2,
		e.Limit,
		e.Str1,
		e.Str2,
	); err != nil {
		return nil, err
	}

	return e.Process(ctx)
}

func (f *fbInPg) Process(ctx context.Context) (*domain.Statistics, error) {
	var int1, int2, limit int

	var str1, str2 string

	var count int

	if err := f.db.QueryRowContext(ctx, getStats).Scan(&count, &int1, &int2, &limit, &str1, &str2); err != nil {
		return nil, err
	}

	return &domain.Statistics{
		Hits: int32(count),
		Fizzbuz: domain.Fizzbuz{
			Int1:  int1,
			Int2:  int2,
			Limit: limit,
			Str1:  str1,
			Str2:  str2,
		},
	}, nil
}
