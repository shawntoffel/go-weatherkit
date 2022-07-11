package weatherkit

import (
	"testing"

	"golang.org/x/text/language"
)

func TestWeatherAlertRequestFullUrlGeneration(t *testing.T) {
	req := WeatherAlertRequest{
		ID:       "test",
		Language: language.English,
	}

	want := BaseUrl + "/weatherAlert/en/test"
	have := req.url()

	if want != have {
		t.Errorf("want: %s, have: %s", want, have)
	}
}
