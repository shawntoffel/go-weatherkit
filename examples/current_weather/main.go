package main

import (
	"context"
	"fmt"
	"os"

	"github.com/shawntoffel/go-weatherkit"
)

// print the current temp in new york
func main() {
	privateKeyBytes, err := os.ReadFile("/path/to/AuthKey_ABCDE12345.p8")
	if err != nil {
		fmt.Println("failed to load private key", err.Error())
		return
	}

	client := weatherkit.NewCredentialedClient(weatherkit.Credentials{
		KeyID:      "key ID",
		TeamID:     "team ID",
		ServiceID:  "service ID",
		PrivateKey: privateKeyBytes,
	})

	ctx := context.Background()

	request := weatherkit.WeatherRequest{
		Latitude:  38.960,
		Longitude: -104.506,
		Language:  "en",
		DataSets: weatherkit.DataSets{
			weatherkit.DataSetCurrentWeather,
		},
	}

	weather, err := client.Weather(ctx, request)
	if err != nil {
		fmt.Println("error", err.Error())
		return
	}

	fmt.Println(weather.CurrentWeather.Temperature)
}
