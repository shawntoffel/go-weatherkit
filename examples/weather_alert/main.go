package main

import (
	"context"
	"fmt"

	"github.com/shawntoffel/go-weatherkit"
	"golang.org/x/text/language"
)

// print event text for an alert id
func main() {
	client := weatherkit.Client{}

	token := "my token"
	ctx := context.Background()

	request := weatherkit.WeatherAlertRequest{
		ID:       "alert id",
		Language: language.English,
	}

	response, err := client.Alert(ctx, token, request)
	if err != nil {
		fmt.Println("error", err.Error())
		return
	}

	fmt.Println(response.EventText[0].Text)
}
