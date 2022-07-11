package main

import (
	"context"
	"fmt"

	"github.com/shawntoffel/go-weatherkit"
	"golang.org/x/text/language"
)

// print hour 0 temp in new york
func main() {
	client := weatherkit.Client{}

	token := "my token"
	ctx := context.Background()

	request := weatherkit.WeatherRequest{
		Latitude:  38.960,
		Longitude: -104.506,
		Language:  language.English,
		DataSets: weatherkit.DataSets{
			weatherkit.DataSetForecastHourly,
		},
	}

	weather, err := client.Weather(ctx, token, request)
	if err != nil {
		fmt.Println("error", err.Error())
		return
	}

	fmt.Println(weather.ForcastHourly.Hours[0].Temperature)
}
