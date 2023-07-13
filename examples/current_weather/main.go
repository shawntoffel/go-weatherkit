package main

import (
	"context"
	"fmt"

	"github.com/nicknassar/weatherkit"
)

// print the current temp in new york
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
		},
	}

	weather, err := client.Weather(ctx, token, request)
	if err != nil {
		fmt.Println("error", err.Error())
		return
	}

	fmt.Println(weather.CurrentWeather.Temperature)
}
