package risken

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"runtime"
	"strings"
	"time"
)

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

var defaultHTTPClient HTTPClient = newDefaultHTTPClient()

func newDefaultHTTPClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   60 * time.Second,
				KeepAlive: 60 * time.Second,
			}).DialContext,
			MaxIdleConns:          10,
			IdleConnTimeout:       60 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
		},
	}
}

type Client struct {
	apiEndpoint string
	apiToken    string
	HTTPClient  HTTPClient
	logger      *slog.Logger
}

func NewClient(apiToken string, options ...ClientOptions) *Client {
	client := &Client{
		apiToken:   apiToken,
		HTTPClient: defaultHTTPClient,
		logger:     slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}
	for _, opt := range options {
		opt(client)
	}
	return client
}

type ClientOptions func(*Client)

func WithAPIEndpoint(endpoint string) ClientOptions {
	return func(c *Client) {
		c.apiEndpoint = endpoint
	}
}

// Do sets some headers on the request, before actioning it using the internal
// HTTPClient. This also assumes any request body is in JSON format and sets
// the Content-Type to application/json.
func (c *Client) Do(r *http.Request) (*http.Response, error) {
	c.prepRequest(r, nil)
	resp, err := c.HTTPClient.Do(r)
	return c.checkResponse(resp, err)
}

func (c *Client) checkResponse(resp *http.Response, err error) (*http.Response, error) {
	if err != nil {
		return nil, fmt.Errorf("error calling the API endpoint: %v", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, c.getErrorFromResponse(resp)
	}
	return resp, nil
}

func (c *Client) getErrorFromResponse(resp *http.Response) APIError {
	// check whether the error response is declared as JSON
	if !strings.HasPrefix(resp.Header.Get("Content-Type"), "application/json") {
		defer resp.Body.Close()
		aerr := APIError{
			Status:  resp.StatusCode,
			Message: fmt.Sprintf("HTTP response with status code %d does not contain Content-Type: application/json", resp.StatusCode),
		}
		return aerr
	}

	aerr := APIError{}
	aerr.Status = resp.StatusCode
	// because of above check this probably won't fail, but it's possible...
	if err := decodeBody(resp, &aerr); err != nil {
		aerr.Message = fmt.Sprintf("HTTP response with status code %d, JSON error object decode failed: %s", resp.StatusCode, err)
		return aerr
	}
	return aerr
}

type APIError struct {
	// StatusCode is the HTTP response status code
	Status  int    `json:"status"`
	Message string `json:"error"`
}

// Error satisfies the error interface, and should contain the StatusCode,
// APIErrorObject.Message, and APIErrorObject.Code.
func (a APIError) Error() string {
	if len(a.Message) > 0 {
		return a.Message
	}
	return fmt.Sprintf("HTTP response with status code %d", a.Status)
}

const (
	VERSION             = "0.0.1"
	USER_AGENT          = "go-risken/" + VERSION
	ACCEPT_HEADER       = "application/json"
	CONTENT_TYPE_HEADER = "application/json"
)

func (c *Client) prepRequest(req *http.Request, headers map[string]string) {
	req.Header.Set("Accept", ACCEPT_HEADER)
	req.Header.Set("Authorization", "Bearer "+c.apiToken)
	req.Header.Set("User-Agent", USER_AGENT)
	req.Header.Set("Content-Type", CONTENT_TYPE_HEADER)

	for k, v := range headers {
		req.Header.Set(k, v)
	}
}

func (c *Client) NewRequest(ctx context.Context, method, path string, param interface{}) (*http.Request, error) {
	parsedURL, err := url.Parse(c.apiEndpoint + path)
	if err != nil {
		return nil, err
	}

	var req *http.Request
	switch method {
	case http.MethodGet:
		if param != nil {
			parsedURL.RawQuery = structToQueryParams(param).Encode()
		}
		req, err = http.NewRequestWithContext(ctx, method, parsedURL.String(), nil)
		if err != nil {
			return nil, err
		}
	case http.MethodPost:
		var buf []byte
		if param != nil {
			buf, err = structToJSONRequestBody(param)
			if err != nil {
				return nil, err
			}
		}
		req, err = http.NewRequestWithContext(ctx, method, parsedURL.String(), bytes.NewBuffer(buf))
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("invalid method: %s", method)
	}
	return req, nil
}

func structToQueryParams(s interface{}) url.Values {
	v := reflect.ValueOf(s)

	// Check if the passed interface is a pointer, and if so, dereference it
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	// Check if the value is a struct
	if v.Kind() != reflect.Struct {
		fmt.Printf("Expected a struct, got %s\n", v.Kind())
		return nil
	}
	t := v.Type()
	params := url.Values{}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")
		// Skip if tag is empty or has "-"
		if tag == "" || tag == "-" {
			continue
		}

		jsonTagParts := strings.Split(tag, ",")
		jsonName := jsonTagParts[0]

		fieldValue := v.Field(i)

		switch fieldValue.Kind() {
		case reflect.Slice:
			for j := 0; j < fieldValue.Len(); j++ {
				params.Add(jsonName, fmt.Sprintf("%v", fieldValue.Index(j)))
			}
		case reflect.String, reflect.Int, reflect.Uint32, reflect.Int32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.Bool:
			// Add more types as needed
			if fieldValue.Interface() != reflect.Zero(fieldValue.Type()).Interface() {
				params.Add(jsonName, fmt.Sprintf("%v", fieldValue.Interface()))
			}
		}
	}
	return params
}

func structToJSONRequestBody(s interface{}) ([]byte, error) {
	requestBody, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return requestBody, nil
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}

func decodeBodyWithDataKey(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	type Temp struct {
		Data json.RawMessage `json:"data"`
	}

	var temp Temp
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&temp); err != nil {
		return err
	}
	return json.Unmarshal(temp.Data, out)
}
