package main

import (
	"context"
	"fmt"
	"os"

	"github.com/shawntoffel/go-weatherkit"
)

// print data set availability in new york
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

	availability, err := client.Availability(ctx, weatherkit.AvailabilityRequest{
		Latitude:  38.960,
		Longitude: -104.506,
	})
	if err != nil {
		fmt.Println("error", err.Error())
		return
	}

	fmt.Println(availability)
}
