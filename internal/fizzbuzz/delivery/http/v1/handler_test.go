package v1

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"

	"github.com/py4mac/fizzbuzz/internal/fizzbuzz/domain"
	"github.com/py4mac/fizzbuzz/internal/fizzbuzz/mock"
)

func TestV1Handlers_Record(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFbUC := mock.NewMockUseCase(ctrl)
	newsHandlers := NewV1Handlers(mockFbUC)

	handlerFunc := newsHandlers.Record()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/fizzbuzz?int1=3&int2=5&limit=10&str1=fizz&str2=buzz", nil)
	res := httptest.NewRecorder()
	e := echo.New()
	ctx := e.NewContext(req, res)

	retExpected := "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz"
	mockFbUC.EXPECT().Record(gomock.Any(), gomock.Any()).Return(retExpected, nil)

	err := handlerFunc(ctx)
	require.NoError(t, err)
}

func TestV1Handlers_Process(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFbUC := mock.NewMockUseCase(ctrl)
	newsHandlers := NewV1Handlers(mockFbUC)

	handlerFunc := newsHandlers.Process()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/stats", nil)
	res := httptest.NewRecorder()
	e := echo.New()
	ctx := e.NewContext(req, res)

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
	mockFbUC.EXPECT().Process(gomock.Any()).Return(statExpected, nil)

	err := handlerFunc(ctx)
	require.NoError(t, err)
}