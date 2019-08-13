package service_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestColourService_FetchColourFromHexbot(t *testing.T) {

	tests  := []struct{
		Desc string
		ExpectedError error
		StartTestServer bool
	}{
		{
			Desc: "fails due to GET error",
			ExpectedError: errors.New("problem getting hex from hexbot"),
			StartTestServer: true,

		},
	}

	for tt := range tests {
		t.Run(tt.Desc, func(t *testing.T) {

			if StartTestServer == true {
				server := httptest.NewServer()

		})

	}

}