package main

import (
	"context"
	"fmt"

	"github.com/nicknassar/weatherkit"
)

// print data set availability in new york
func main() {
	client := weatherkit.Client{}

	token := "my token"
	ctx := context.Background()

	availability, err := client.Availability(ctx, token, weatherkit.AvailabilityRequest{
		Latitude:  38.960,
		Longitude: -104.506,
	})
	if err != nil {
		fmt.Println("error", err.Error())
		return
	}

	fmt.Println(availability)
}
