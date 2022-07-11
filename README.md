# go-weatherkit

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

Import the package into your project:

```go
import "github.com/shawntoffel/go-weatherkit"
```

Create a new weatherkit client:

```go
client := weatherkit.Client{}
```

Build a request:

```go
request := weatherkit.WeatherRequest{
		Latitude:  38.960,
		Longitude: -104.506,
		Language:  language.English,
		DataSets: weatherkit.DataSets{
			weatherkit.DataSetCurrentWeather,
		},
	}
```

Get a response:
```go
ctx := context.Background()

response, err := client.Weather(ctx, token, request)
```
The `token` parameter is a JWT developer token used for authentication. See the documentation [here](https://developer.apple.com/documentation/weatherkitrestapi/request_authentication_for_weatherkit_rest_api) for details on JWT creation.

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

## Troubleshooting
Please use the GitHub discussions tab for questions regarding this client library. The Apple Developer forums are available for questions regarding the underlying API: https://developer.apple.com/forums/tags/weatherkit