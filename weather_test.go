package weatherkit

import (
	"testing"
	"time"
)

func TestWeatherRequestFullUrlGeneration(t *testing.T) {
	ts, _ := time.Parse(dateTimeFormat, "2022-07-10T06:14:11Z")

	req := WeatherRequest{
		Language:    "en",
		Latitude:    40.713,
		Longitude:   -74.006,
		CountryCode: "US",
		CurrentAsOf: &ts,
		DailyEnd:    &ts,
		DailyStart:  &ts,
		HourlyEnd:   &ts,
		HourlyStart: &ts,
		Timezone:    "America/New_York",
		DataSets: DataSets{
			DataSetCurrentWeather,
			DataSetForecastDaily,
			DataSetForecastHourly,
			DataSetForecastNextHour,
			DataSetWeatherAlerts,
		},
	}

	want := BaseUrl + "/api/v1/weather/en/40.713/-74.006?countryCode=US&currentAsOf=2022-07-10T06%3A14%3A11Z&dailyEnd=2022-07-10T06%3A14%3A11Z&dailyStart=2022-07-10T06%3A14%3A11Z&dataSets=currentWeather%2CforecastDaily%2CforecastHourly%2CforecastNextHour%2CweatherAlerts&hourlyEnd=2022-07-10T06%3A14%3A11Z&hourlyStart=2022-07-10T06%3A14%3A11Z&timezone=America%2FNew_York"
	have := req.url()

	if want != have {
		t.Errorf("want: %s, have: %s", want, have)
	}
}
