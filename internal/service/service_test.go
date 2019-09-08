package service_test

import (
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

func TestService_FetchColourFromHexbot(t *testing.T) {

	tests := []struct {
		Desc          string
		Colour        string
		GETerr        error
		ExpectedError error
	}{
		{
			Desc:          "fails due to GET error",
			GETerr:        errors.New("GET unsuccessful"),
			ExpectedError: errors.New("problem getting hex from hexbot: GET unsuccessful"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.Desc, func(t *testing.T) {

			s := newTestService(t)
			defer s.Finish()

			s.httpClient.EXPECT().GetColourFromHexbot().Return(tt.Colour, tt.GETerr)

			_, err := s.service.FetchColour()

			if tt.ExpectedError != nil {
				require.Error(t, err)
				require.EqualError(t, err, tt.ExpectedError.Error())

			} else {
				require.NoError(t, err)
			}

		})

	}

}

func TestService_SaveColour(t *testing.T) {
	tests := []struct {
		Desc string
		Colour string
		ExpectedErr error
	}{
		{
			Desc: "should fail because of an empty  colour  string",
			Colour: "",
			ExpectedErr: errors.New("trying to save an empty colour string"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.Desc, func(t *testing.T) {
			s := newTestService(t)
			defer s.Finish()

			err  := s.service.SaveColour(tt.Colour)

			if tt.ExpectedErr != nil  {
				require.Error(t, err)
				require.EqualError(t, err, tt.ExpectedErr.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
