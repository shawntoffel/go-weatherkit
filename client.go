package weatherkit

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

// DefaultUserAgent to send along with requests.
const DefaultUserAgent = "shawntoffel/go-weatherkit"

// NewCredentialedClient creates a new client with creds.
func NewCredentialedClient(credentials Credentials, opts ...CredentialedClientOption) *CredentialedClient {
	cc := &CredentialedClient{
		credentials: credentials,
		options:     defaultcredentialedClientOptions(),
	}

	if opts == nil || len(opts) < 1 {
		return cc
	}

	for _, opt := range opts {
		opt.apply(cc.options)
	}

	return cc
}

// CredentialedClient is a WeatherKit API client.
// Construct with NewCredentialedClient.
type CredentialedClient struct {
	options     *credentialedClientOptions
	credentials Credentials
	mu          sync.Mutex
	token       string
	exp         time.Time
}

func (c *CredentialedClient) getToken() (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.options == nil {
		c.options = defaultcredentialedClientOptions()
	}

	if c.options.disableCache {
		signed, _, err := c.credentials.SignedJWT(c.options.tokenDuration)
		return signed, err
	}

	// Use a minute buffer to allow for req/resp time.
	if len(c.token) > 0 && c.exp.After(time.Now().UTC().Add(time.Minute)) {
		return c.token, nil
	}

	signed, exp, err := c.credentials.SignedJWT(c.options.tokenDuration)
	if err != nil {
		return "", err
	}

	c.token = signed
	c.exp = exp

	return c.token, nil
}

// Weather obtains weather data for the specified location.
func (d *CredentialedClient) Weather(ctx context.Context, request WeatherRequest) (*WeatherResponse, error) {
	token, err := d.getToken()
	if err != nil {
		return nil, err
	}
	return d.options.client.Weather(ctx, token, request)
}

// Availability determines the data sets available for the specified location.
func (d *CredentialedClient) Availability(ctx context.Context, request AvailabilityRequest) (*AvailabilityResponse, error) {
	token, err := d.getToken()
	if err != nil {
		return nil, err
	}
	return d.options.client.Availability(ctx, token, request)
}

// Alert receives information on an active weather alert.
func (d *CredentialedClient) Alert(ctx context.Context, request WeatherAlertRequest) (*WeatherAlertResponse, error) {
	token, err := d.getToken()
	if err != nil {
		return nil, err
	}
	return d.options.client.Alert(ctx, token, request)
}

// Attribution retrieves official attribution branding.
func (d *CredentialedClient) Attribution(ctx context.Context, request AttributionRequest) (*AttributionResponse, error) {
	return d.options.client.Attribution(ctx, request)
}

// CredentialedClientOption configures a CredentialedClient.
type CredentialedClientOption interface {
	apply(*credentialedClientOptions)
}

type credentialedClientOptions struct {
	disableCache  bool
	client        *Client
	tokenDuration time.Duration
}

type funcOption struct {
	f func(*credentialedClientOptions)
}

func (fo *funcOption) apply(o *credentialedClientOptions) {
	fo.f(o)
}

func newFuncOption(f func(*credentialedClientOptions)) *funcOption {
	return &funcOption{
		f: f,
	}
}

func defaultcredentialedClientOptions() *credentialedClientOptions {
	return &credentialedClientOptions{
		client:        &Client{},
		tokenDuration: defaultTokenDuration,
	}
}

// WithoutCache returns an Option which disables token caching.
// A new JWT will be generated for each request.
func WithoutCache() CredentialedClientOption {
	return newFuncOption(func(o *credentialedClientOptions) {
		o.disableCache = true
	})
}

// WithClient returns an Option which configures a custom Client.
func WithClient(client *Client) CredentialedClientOption {
	return newFuncOption(func(o *credentialedClientOptions) {
		o.client = client
	})
}

// WithTokenDuration returns an Option which configures the expiration duration of the internally generated JWTs.
// The default duration is 10 minutes.
func WithTokenDuration(duration time.Duration) CredentialedClientOption {
	return newFuncOption(func(o *credentialedClientOptions) {
		o.tokenDuration = duration
	})
}

// Client is a WeatherKit API client without Credentials.
// Use NewCredentialedClient for automatic JWT handling.
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

// Attribution retrieves official attribution branding.
func (d *Client) Attribution(ctx context.Context, request AttributionRequest) (*AttributionResponse, error) {
	response := AttributionResponse{}
	err := d.get(ctx, "", request, &response)
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

	if len(token) > 0 {
		req.Header.Add("Authorization", "Bearer "+token)
	}

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
