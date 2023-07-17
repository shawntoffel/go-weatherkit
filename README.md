# go-weatherkit

[![Go Reference](https://pkg.go.dev/badge/github.com/shawntoffel/go-weatherkit.svg)](https://pkg.go.dev/github.com/shawntoffel/go-weatherkit) 
 [![Go Report Card](https://goreportcard.com/badge/github.com/shawntoffel/go-weatherkit)](https://goreportcard.com/report/github.com/shawntoffel/go-weatherkit) [![Build status](https://github.com/shawntoffel/go-weatherkit/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/shawntoffel/go-weatherkit/actions/workflows/go.yml)

A [WeatherKit](https://developer.apple.com/weatherkit/) API client in Go. WeatherKit is powered by the Apple Weather service.

Notice: The WeatherKit REST API is currently in beta and is subject to change. This client was created from documentation available here: https://developer.apple.com/documentation/weatherkitrestapi

*go-weatherkit* is an open source project not affiliated with Apple Inc.

## Installing

This assumes you already have a working Go environment, if not please see
[this page](https://golang.org/doc/install) first.

```sh
go get github.com/shawntoffel/go-weatherkit
```

## Usage

### Import the package into your project:

```go
import "github.com/shawntoffel/go-weatherkit"
```

### Create a new weatherkit client:

```go
// See Authentication documentation below.
privateKeyBytes, _ := os.ReadFile("/path/to/AuthKey_ABCDE12345.p8")

client := weatherkit.NewCredentialedClient(weatherkit.Credentials{
	KeyID:      "key ID",
	TeamID:     "team ID",
	ServiceID:  "service ID",
	PrivateKey: privateKeyBytes,
})
```
Locating your identifiers:
* **Key ID (kid)**: An identifier associated with your private key. It can be found on the [Certificates, Identifiers & Profiles](https://developer.apple.com/account/resources/authkeys/list) page under Keys. Click on the appropriate key to view the ID. 
* **Team ID (tid)**: Found on the [account](https://developer.apple.com/account) page under Membership details.
* **Service ID (sid)**: Found on the [Certificates, Identifiers & Profiles](https://developer.apple.com/account/resources/identifiers/list/serviceId) page under Identifiers. Make sure "Services IDs" is selected from the dropdown. 

### Build a request:

```go
request := weatherkit.WeatherRequest{
	Latitude:  38.960,
	Longitude: -104.506,
	Language:  "en",
	DataSets: weatherkit.DataSets{
		weatherkit.DataSetCurrentWeather,
	},
}
```

### Get a response:
```go
ctx := context.Background()

response, err := client.Weather(ctx, request)
```

## Documentation

- [![Go Reference](https://pkg.go.dev/badge/github.com/shawntoffel/go-weatherkit.svg)](https://pkg.go.dev/github.com/shawntoffel/go-weatherkit) 

- [REST API](https://developer.apple.com/documentation/weatherkitrestapi)
- [Authentication](https://developer.apple.com/documentation/weatherkitrestapi/request_authentication_for_weatherkit_rest_api)

## Attribution
See Apple's documentation for Apple Weather and third-party attribution requirements.
- https://developer.apple.com/weatherkit/get-started#attribution-requirements

- https://weatherkit.apple.com/legal-attribution.html

## Examples
- [Current weather](https://github.com/shawntoffel/go-weatherkit/tree/master/examples/current_weather/main.go)
- [Hourly forecast](https://github.com/shawntoffel/go-weatherkit/tree/master/examples/hourly_forecast/main.go)
- [Multiple data sets](https://github.com/shawntoffel/go-weatherkit/tree/master/examples/multiple_datasets/main.go)
- [DIY Credentials](https://github.com/shawntoffel/go-weatherkit/tree/master/examples/diy_credentials/main.go)

## Troubleshooting
Please use the GitHub [Discussions](https://github.com/shawntoffel/go-weatherkit/discussions) tab for questions regarding this client library. The Apple Developer forums are available for questions regarding the underlying API: https://developer.apple.com/forums/tags/weatherkit
