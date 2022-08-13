package weatherkit

import (
	"testing"
)

func TestWeatherAlertRequestFullUrlGeneration(t *testing.T) {
	req := WeatherAlertRequest{
		ID:       "test",
		Language: "en",
	}

	want := BaseUrl + "/weatherAlert/en/test"
	have := req.url()

	if want != have {
		t.Errorf("want: %s, have: %s", want, have)
	}
}
