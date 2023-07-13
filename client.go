package weatherkit

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// DefaultUserAgent to send along with requests.
const DefaultUserAgent = "nicknassar-weatherkit"

// Client is a WeatherKit API client.
type Client struct {
	HttpClient *http.Client

	// The UserAgent header value to send along with requests.
	UserAgent string
}

// Weather obtains weather data for the specified location.
// The token parameter is a JWT developer token.
func (d *Client) Weather(ctx context.Context, token string, request WeatherRequest) (*WeatherResponse, error) {
	response := WeatherResponse{}
	err := d.get(ctx, token, request, &response)
	return &response, err
}

// Availability determines the data sets available for the specified location.
// The token parameter is a JWT developer token.
func (d *Client) Availability(ctx context.Context, token string, request AvailabilityRequest) (*AvailabilityResponse, error) {
	response := AvailabilityResponse{}
	err := d.get(ctx, token, request, &response)
	return &response, err
}

// Alert receives information on an active weather alert.
// The token parameter is a JWT developer token.
func (d *Client) Alert(ctx context.Context, token string, request WeatherAlertRequest) (*WeatherAlertResponse, error) {
	response := WeatherAlertResponse{}
	err := d.get(ctx, token, request, &response)
	return &response, err
}

func (d *Client) get(ctx context.Context, token string, request urlBuilder, output interface{}) error {
	if d.HttpClient == nil {
		d.HttpClient = &http.Client{}
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, request.url(), nil)
	if err != nil {
		return err
	}

	req.Header.Add("User-Agent", d.userAgent())
	req.Header.Add("Accept", "application/json; charset=utf-8")
	req.Header.Add("Accept-Encoding", "gzip")
	req.Header.Add("Authorization", "Bearer "+token)

	response, err := d.HttpClient.Do(req)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	err = validateResponse(response)
	if err != nil {
		return err
	}

	return decode(response, &output)
}

func (d *Client) userAgent() string {
	if len(d.UserAgent) > 0 {
		return d.UserAgent
	}

	return DefaultUserAgent
}

func validateResponse(response *http.Response) error {
	if response.StatusCode == http.StatusOK {
		return nil
	}

	errorResponse := ErrorResponse{}

	err := decode(response, &errorResponse)
	if err != nil {
		return err
	}

	return &RestError{
		Response:      response,
		ErrorResponse: &errorResponse,
	}
}

func decode(response *http.Response, into interface{}) error {
	body, err := decompress(response)
	if err != nil {
		return err
	}

	return unmarshal(body, into)
}

func decompress(response *http.Response) (io.Reader, error) {
	header := response.Header.Get("Content-Encoding")
	if len(header) < 1 {
		return response.Body, nil
	}

	return gzip.NewReader(response.Body)
}

func unmarshal(body io.Reader, into interface{}) error {
	bytes, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}

	if bytes == nil || len(bytes) < 1 {
		return nil
	}

	return json.Unmarshal(bytes, &into)
}
