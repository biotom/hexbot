package service_test

import (
	"errors"
	"github.com/River-Island/product-backbone-v2/logging"
	"github.com/golang/mock/gomock"
	"hexbot/internal/service"
	"testing"
)


type testService struct {
	ctrl *gomock.Controller
	log *logging.Logger
	mockDB *service.MockDatabase
	httpClient *service.MockHTTPClient
	service	*service.ColourService
}

func newTestService(t *testing.T) *testService {
	ctrl := gomock.NewController(t)
	db := service.NewMockDatabase(ctrl)
	hc := service.NewMockHTTPClient(ctrl)
	s := &testService{
		ctrl:    ctrl,
		mockDB:   db,
		httpClient: hc,
		service: service.NewColourService(logging.NopLogger, db, hc),
		//needs a hexbot client
	}
	return s
}

func TestColourService_FetchColourFromHexbot(t *testing.T) {

	tests := []struct {
		Desc            string
		ExpectedError   error
	}{
		{
			Desc:            "fails due to GET error",
			ExpectedError:   errors.New("problem getting hex from hexbot"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.Desc, func(t *testing.T) {




		})

	}

}


