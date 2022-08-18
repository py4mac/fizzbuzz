package usecase

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/py4mac/fizzbuzz/internal/fizzbuzz/domain"
	"github.com/py4mac/fizzbuzz/internal/fizzbuzz/mock"
)

func TestFbUC_Record(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFbRepo := mock.NewMockRepository(ctrl)
	fbUC := NewFBUseCase(mockFbRepo)

	fb := &domain.Fizzbuz{
		Int1:  3,
		Int2:  5,
		Limit: 10,
		Str1:  "fizz",
		Str2:  "buzz",
	}

	retExpected := "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz"
	mockFbRepo.EXPECT().Record(gomock.Any(), gomock.Any()).Return(retExpected, nil)

	ret, err := fbUC.Record(context.Background(), fb)
	require.NoError(t, err)
	require.Equal(t, retExpected, ret)
}

func TestFbUC_Process(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFbRepo := mock.NewMockRepository(ctrl)
	fbUC := NewFBUseCase(mockFbRepo)

	statExpected := &domain.Statistics{
		Hits: 10,
		Fizzbuz: domain.Fizzbuz{
			Int1:  3,
			Int2:  5,
			Limit: 10,
			Str1:  "fizz",
			Str2:  "buzz",
		},
	}

	mockFbRepo.EXPECT().Process(gomock.Any()).Return(statExpected, nil)

	ret, err := fbUC.Process(context.Background())
	require.NoError(t, err)
	require.Equal(t, statExpected, ret)
}
