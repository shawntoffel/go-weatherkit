# weatherkit

A [WeatherKit](https://developer.apple.com/weatherkit/) API client in Go.

*weatherkit* is a fork of [go-weatherkit](https://github.com/shawntoffel/go-weatherkit) that provides support for managed authentication. It is an open source project not affiliated with Apple Inc.

Notice: This client was created from documentation available here: https://developer.apple.com/documentation/weatherkitrestapi

## Installing

This assumes you already have a working Go environment, if not please see
[this page](https://golang.org/doc/install) first.

```sh
go get github.com/nicknassar/weatherkit
```

## Usage

Import the package into your project:

```go
import "github.com/nicknassar/weatherkit"
```

Create a new weatherkit client:

```go
client := weatherkit.NewClient(keyId, teamId, serviceId, privateKey)
```

Build a request:

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

Get a response:
```go
ctx := context.Background()

response, err := client.Weather(ctx, request)
```

## Authentication
A JWT developer `token` is used to authenticate requests. See the documentation [here](https://developer.apple.com/documentation/weatherkitrestapi/request_authentication_for_weatherkit_rest_api) for details on WeatherKit API authentication.

The client requires keyId, teamId, serviceId, and the private key from Apple to generate the JWT token.

## Documentation

- [![Go Reference](https://pkg.go.dev/badge/github.com/nicknassar/weatherkit.svg)](https://pkg.go.dev/github.com/nicknassar/weatherkit)

- [REST API](https://developer.apple.com/documentation/weatherkitrestapi)
- [Authentication](https://developer.apple.com/documentation/weatherkitrestapi/request_authentication_for_weatherkit_rest_api)

## Attribution
See Apple's documentation for Apple Weather and third-party attribution requirements.
- https://developer.apple.com/weatherkit/get-started#attribution-requirements

- https://weatherkit.apple.com/legal-attribution.html

## Examples
- [Current weather](https://github.com/nicknassar/weatherkit/tree/master/examples/current_weather/main.go)
- [Hourly forecast](https://github.com/nicknassar/weatherkit/tree/master/examples/hourly_forecast/main.go)
- [Multiple data sets](https://github.com/nicknassar/weatherkit/tree/master/examples/multiple_datasets/main.go)

## Troubleshooting
Please use the GitHub [Discussions](https://github.com/nicknassar/weatherkit/issues) to report issues. The Apple Developer forums are available for questions regarding the underlying API: https://developer.apple.com/forums/tags/weatherkit
