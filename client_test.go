package weatherkit

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestParseCurrentWeatherResponse(t *testing.T) {
	weather(t, "testdata/current_weather.json")
}

func TestParseForecastDailyResponse(t *testing.T) {
	weather(t, "testdata/forecast_daily.json")
}

func TestParseForecastHourlyResponse(t *testing.T) {
	weather(t, "testdata/forecast_hourly.json")
}

func TestParseNextHourForecastResponse(t *testing.T) {
	weather(t, "testdata/next_hour_forecast.json")
}

func TestParseFullForecastResponse(t *testing.T) {
	weather(t, "testdata/full_weather.json")
}

func TestParseFullAvailabilityResponse(t *testing.T) {
	availability(t, "testdata/availability.json")
}

func TestWeatherErrorResponse(t *testing.T) {
	client := Client{}

	server, expected, err := getMockServerWithFileData("testdata/weather_error.json", http.StatusNotFound)
	if err != nil {
		t.Error(err.Error())
	}

	defer server.Close()
	BaseUrl = server.URL

	_, err = client.Weather(context.TODO(), "", WeatherRequest{})
	if err == nil {
		t.Errorf("expected request to error")
		return
	}

	restError, ok := err.(*RestError)
	if !ok {
		t.Errorf("expected error to be of type %s", reflect.TypeOf(RestError{}))
		return
	}

	actual, err := json.Marshal(restError.ErrorResponse)
	if err != nil {
		t.Errorf(err.Error())
	}

	assertJsonEqual(t, expected, actual)
}

func TestWeatherAlertErrorResponse(t *testing.T) {
	client := Client{}

	server, expected, err := getMockServerWithFileData("testdata/weather_alert_error.json", http.StatusBadRequest)
	if err != nil {
		t.Error(err.Error())
	}

	defer server.Close()
	BaseUrl = server.URL

	_, err = client.Alert(context.TODO(), "", WeatherAlertRequest{})
	if err == nil {
		t.Errorf("expected request to error")
		return
	}

	restError, ok := err.(*RestError)
	if !ok {
		t.Errorf("expected error to be of type %s", reflect.TypeOf(RestError{}))
		return
	}

	actual, err := json.Marshal(restError.ErrorResponse)
	if err != nil {
		t.Errorf(err.Error())
	}

	assertJsonEqual(t, expected, actual)
}

func TestAttributionResponse(t *testing.T) {
	client := Client{}

	server, expected, err := getMockServerWithFileData("testdata/attribution.json", http.StatusOK)
	if err != nil {
		t.Error(err.Error())
	}

	defer server.Close()
	BaseUrl = server.URL

	response, err := client.Attribution(context.TODO(), AttributionRequest{
		Language: "en",
	})
	if err != nil {
		t.Error(err.Error())
	}

	actual, err := json.Marshal(response)
	if err != nil {
		t.Errorf(err.Error())
	}

	assertJsonEqual(t, expected, actual)
}

func weather(t *testing.T, filename string) {
	pk, err := createPrivateKeyPEM()
	if err != nil {
		t.Error(err)
	}

	client := NewCredentialedClient(Credentials{
		KeyID:      "key",
		TeamID:     "team",
		ServiceID:  "service",
		PrivateKey: pk,
	})

	server, expected, err := getMockServerWithFileData(filename, http.StatusOK)
	if err != nil {
		t.Error(err.Error())
	}

	defer server.Close()
	BaseUrl = server.URL

	response, err := client.Weather(context.TODO(), WeatherRequest{})
	if err != nil {
		t.Error(err.Error())
	}

	actual, err := json.Marshal(response)
	if err != nil {
		t.Errorf(err.Error())
	}

	assertJsonEqual(t, expected, actual)
}

func availability(t *testing.T, filename string) {
	pk, err := createPrivateKeyPEM()
	if err != nil {
		t.Error(err)
	}

	client := NewCredentialedClient(Credentials{
		KeyID:      "key",
		TeamID:     "team",
		ServiceID:  "service",
		PrivateKey: pk,
	})

	server, expected, err := getMockServerWithFileData(filename, http.StatusOK)
	if err != nil {
		t.Error(err.Error())
	}

	defer server.Close()
	BaseUrl = server.URL

	response, err := client.Availability(context.TODO(), AvailabilityRequest{})
	if err != nil {
		t.Error(err.Error())
	}

	actual, err := json.Marshal(response)
	if err != nil {
		t.Errorf(err.Error())
	}

	assertJsonEqual(t, expected, actual)
}

func assertJsonEqual(t *testing.T, expected []byte, actual []byte) {
	var a, b interface{}

	err := json.Unmarshal(expected, &a)
	if err != nil {
		t.Error(err.Error())
	}

	err = json.Unmarshal(actual, &b)
	if err != nil {
		t.Error(err.Error())
	}

	if !reflect.DeepEqual(a, b) {
		t.Errorf("\nexpected: %+v\nactual: %+v", a, b)
	}
}

func getMockServerWithFileData(filename string, statusCode int) (*httptest.Server, []byte, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, bytes, err
	}

	return getMockServer(bytes, statusCode), bytes, nil
}

func getMockServer(data []byte, statusCode int) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		fmt.Fprintln(w, string(data))
	}))

	return server
}
