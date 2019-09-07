package service_test

import (
	"context"
	"errors"
	"hexbot/internal/service"
	"testing"

	"github.com/River-Island/product-backbone-v2/logging"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

type testService struct {
	ctrl       *gomock.Controller
	log        *logging.Logger
	mockDB     *service.MockDatabase
	httpClient *service.MockHTTPClient
	service    *service.Service
}

func newTestService(t *testing.T) *testService {
	ctrl := gomock.NewController(t)
	db := service.NewMockDatabase(ctrl)
	hc := service.NewMockHTTPClient(ctrl)
	s := &testService{
		ctrl:       ctrl,
		mockDB:     db,
		httpClient: hc,
		service:    service.NewService(logging.NopLogger, db, hc),
	}
	return s
}

func (s *testService) Finish() {
	s.ctrl.Finish()
}

func TestColourService_FetchColourFromHexbot(t *testing.T) {

	tests := []struct {
		Desc          string
		Colour        string
		GETerr        error
		ExpectedError error
	}{
		{
			Desc:          "fails due to GET error",
			Colour:        "0000ff",
			GETerr:        errors.New("GET unsuccessful"),
			ExpectedError: errors.New("problem getting hex from hexbot: GET unsuccessful"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.Desc, func(t *testing.T) {

			s := newTestService(t)
			defer s.Finish()

			ctx := context.Background()

			s.httpClient.EXPECT().GetColour(ctx).Return(tt.Colour, nil)

			_, err := s.service.FetchColourFromHexbot(ctx)

			if tt.ExpectedError != nil {
				require.Error(t, err)
				require.EqualError(t, err, tt.ExpectedError.Error())

			} else {
				require.NoError(t, err)
			}

		})

	}

}
