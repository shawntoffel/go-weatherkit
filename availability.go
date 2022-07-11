package weatherkit

import (
	"net/url"
)

type AvailabilityRequest struct {
	// The latitude of the desired location.
	Latitude float64

	// The longitude of the desired location.
	Longitude float64

	// (Required) The ISO Alpha-2 country code for the requested location.
	// This parameter is necessary for air quality and weather alerts.
	Country string
}

func (o AvailabilityRequest) url() string {
	q := url.Values{}

	if o.Country != "" {
		q.Add("country", o.Country)
	}

	return availabilityEndpoint(o.Latitude, o.Longitude, q)
}

// AvailabilityResponse has the data sets available for the specified location.
type AvailabilityResponse DataSets
