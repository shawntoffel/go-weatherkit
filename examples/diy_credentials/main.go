package main

import (
	"context"
	"fmt"

	"github.com/shawntoffel/go-weatherkit"
)

// Demonstrates how to use the plain Client if you generate JWT developer tokens yourself.
func main() {
	client := weatherkit.Client{}

	request := weatherkit.WeatherRequest{
		Latitude:  38.960,
		Longitude: -104.506,
		Language:  "en",
		DataSets: weatherkit.DataSets{
			weatherkit.DataSetCurrentWeather,
		},
	}

	ctx := context.Background()
	token := "your JWT developer token generated elsewhere"

	weather, err := client.Weather(ctx, token, request)
	if err != nil {
		fmt.Println("error", err.Error())
		return
	}

	fmt.Println(weather.CurrentWeather.Temperature)
}
