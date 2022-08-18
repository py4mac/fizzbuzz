package repository

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/py4mac/fizzbuzz/internal/fizzbuzz/domain"
	"github.com/stretchr/testify/require"
)

func TestFbRepo_Record(t *testing.T) {
	t.Run("Create", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		fbRepo := NewFBInPg(sqlx.NewDb(db, "sqlmock"))

		int1 := 3
		int2 := 5
		limit := 10
		str1 := "fizz"
		str2 := "buzz"
		mock.ExpectExec(regexp.QuoteMeta(insertRecord)).WillReturnResult(sqlmock.NewResult(1, 1))

		fb := &domain.Fizzbuz{
			Int1:  int1,
			Int2:  int2,
			Limit: limit,
			Str1:  str1,
			Str2:  str2,
		}
		ret, err := fbRepo.Record(context.Background(), fb)
		require.NoError(t, err)
		require.Equal(t, "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz", ret)
	})
}

func TestFbRepo_Process(t *testing.T) {
	t.Run("Process", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		fbRepo := NewFBInPg(sqlx.NewDb(db, "sqlmock"))
		int1 := 3
		int2 := 5
		limit := 10
		str1 := "fizz"
		str2 := "buzz"
		count := 10
		rows := sqlmock.NewRows([]string{"int1", "int2", "max_limit", "str1", "str2", "count"})
		mock.ExpectQuery(regexp.QuoteMeta(getStats)).WillReturnRows(rows.AddRow(int1, int2, limit, str1, str2, count))

		stats, err := fbRepo.Process(context.Background())
		require.NoError(t, err)
		require.Equal(t, int32(count), stats.Hits)
	})
}
