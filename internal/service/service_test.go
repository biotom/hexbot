package service_test

import "testing"

func TestColourService_FetchColourFromHexbot(t *testing.T) {
	for tt := range tests {
		t.run(tt.Desc, func (t *testing.T))
	}

}


