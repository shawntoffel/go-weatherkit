package weatherkit

import (
	"fmt"
	"net/url"
)

// WeatherKit API base URL
var BaseUrl = "https://weatherkit.apple.com"

func weatherEndpoint(lang string, latitude float64, longitude float64, values url.Values) string {
	return BaseUrl + "/api/v1/weather/" + fmt.Sprintf("%s/%g/%g", lang, latitude, longitude) + encodeUrlParameters(values)
}

func availabilityEndpoint(latitude, longitude float64, values url.Values) string {
	return BaseUrl + "/api/v1/availability/" + fmt.Sprintf("%g/%g", latitude, longitude) + encodeUrlParameters(values)
}

func weatherAlertEndpoint(lang string, id string) string {
	return BaseUrl + "/api/v1/weatherAlert/" + fmt.Sprintf("%s/%s", lang, id)
}

func attribution(lang string) string {
	return BaseUrl + "/attribution/" + lang
}

func encodeUrlParameters(values url.Values) string {
	queryString := values.Encode()

	if queryString == "" {
		return queryString
	}

	return "?" + queryString
}

type urlBuilder interface {
	url() string
}
