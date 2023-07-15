package main

import (
	"context"
	"fmt"
	"os"

	"github.com/shawntoffel/go-weatherkit"
)

// print event text for an alert id
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

	request := weatherkit.WeatherAlertRequest{
		ID:       "alert id",
		Language: "en",
	}

	response, err := client.Alert(ctx, request)
	if err != nil {
		fmt.Println("error", err.Error())
		return
	}

	fmt.Println(response.EventText[0].Text)
}
