package weatherkit

import "strings"

// The collection of weather information for a location.
type DataSet string

// The collection of weather information for a location.
type DataSets []DataSet

// String returns a comma delimited list of the data sets.
func (d DataSets) String() string {
	dataSets := make([]string, len(d))
	for i, v := range d {
		dataSets[i] = string(v)
	}
	return strings.Join(dataSets, ",")
}

const (
	// The current weather for the requested location.
	DataSetCurrentWeather DataSet = "currentWeather"

	// The daily forecast for the requested location.
	DataSetForecastDaily DataSet = "forecastDaily"

	// The hourly forecast for the requested location.
	DataSetForecastHourly DataSet = "forecastHourly"

	// The next hour forecast for the requested location.
	DataSetForecastNextHour DataSet = "forecastNextHour"

	// Weather alerts for the requested location.
	DataSetWeatherAlerts DataSet = "weatherAlerts"
)
