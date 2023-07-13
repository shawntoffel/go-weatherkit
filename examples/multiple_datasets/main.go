package main

import (
	"context"
	"fmt"

	"github.com/nicknassar/weatherkit"
)

// print current, day 0 temp max, and hour 0 temp.
func main() {
	client := weatherkit.Client{}

	token := "my token"
	ctx := context.Background()

	request := weatherkit.WeatherRequest{
		Latitude:  38.960,
		Longitude: -104.506,
		Language:  "en",
		DataSets: weatherkit.DataSets{
			weatherkit.DataSetCurrentWeather,
			weatherkit.DataSetForecastDaily,
			weatherkit.DataSetForecastHourly,
		},
	}

	weather, err := client.Weather(ctx, token, request)
	if err != nil {
		fmt.Println("error", err.Error())
		return
	}

	fmt.Println(weather.CurrentWeather.Temperature)
	fmt.Println(weather.ForcastDaily.Days[0].TemperatureMax)
	fmt.Println(weather.ForcastHourly.Hours[0].Temperature)
}
