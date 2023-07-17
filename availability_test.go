package weatherkit

import (
	"testing"
)

func TestAvailabilityRequestFullUrlGeneration(t *testing.T) {
	req := AvailabilityRequest{
		Latitude:  40.713,
		Longitude: -74.006,
		Country:   "US",
	}

	want := BaseUrl + "/api/v1/availability/40.713/-74.006?country=US"
	have := req.url()

	if want != have {
		t.Errorf("want: %s, have: %s", want, have)
	}
}
