package service_test

import (
	"errors"
	"fmt"
	"github.com/River-Island/product-backbone-v2/logging"
	"github.com/golang/mock/gomock"
	"hexbot/internal/service"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)


type testService struct {
	ctrl *gomock.Controller
	log *logging.Logger
	mockDB *service.MockDatabase
	service	*service.ColourService
}

func newTestService(t *testing.T) *testService {
	ctrl := gomock.NewController(t)
	db := service.NewMockDatabase(ctrl)
	s := &testService{
		ctrl:    ctrl,
		mockDB:   db,
		service: service.NewColourService(logging.NopLogger, ),
		//needs a hexbot client
	}
	return s
}

func TestColourService_FetchColourFromHexbot(t *testing.T) {

	tests := []struct {
		Desc            string
		ExpectedError   error
		StartTestServer bool
	}{
		{
			Desc:            "fails due to GET error",
			ExpectedError:   errors.New("problem getting hex from hexbot"),
			StartTestServer: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Desc, func(t *testing.T) {
			if tt.StartTestServer == true {

				ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					_, err  := fmt.Fprintln(w, "#228B22")
					if err != nil {
						t.Fatal(err)
					}
				}))
				defer ts.Close()


				//this happens inside the colour service
				resp, err := http.Get(ts.URL)
				if err != nil {
					t.Fatal(err)
				}


				colour, err := ioutil.ReadAll(resp.Body)
				resp.Body.Close()
				if err != nil{
					t.Fatal(err)
				}



		}

	})

}

}
