package gohttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	httpmock_server "github.com/miknonny/http-client/v3/mock"
	gohttp_types "github.com/miknonny/http-client/v3/types"
)

const (
	defaultMaxIdleConnection = 5
	defaultResponseTimeout   = 5 * time.Second
	defaultConnectionTimeout = 1 * time.Second
)

// Returns the body as bytes based on the Content-Type set in the headers.
func (c *httpClient) marshalRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}

	switch strings.ToLower(contentType) {
	case "application/json":
		return json.Marshal(body)
	case "application/xml":
		return xml.Marshal(body)
	default:
		return json.Marshal(body)
	}
}

// This core function sends the request to the server. // Prepares the request body and headers fo flight.
func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*gohttp_types.Response, error) {

	fullHeaders := c.getAllRequestHeaders(headers)
	requestBody, err := c.marshalRequestBody(fullHeaders.Get("Content-Type"), body)
	if err != nil {
		return nil, err
	}

	// We are still yet to remove the mock from our source code.
	if mock := httpmock_server.GetMock(method, url, string(requestBody)); mock != nil {
		return mock.GetResponse()
	}

	request, err := http.NewRequest(method, url, bytes.NewReader(requestBody)) // bytes.NewBuffer(requestBody) still works.
	if err != nil {
		return nil, errors.New("unable to create a request")
	}

	// setting all headers returned from `getAllReqeustHeaders(headers)` on the request.
	request.Header = fullHeaders

	client := c.getHttpClient()

	// dispatching the http call using go http.Client
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return &gohttp_types.Response{
		Status:     response.Status,
		StatusCode: response.StatusCode,
		Headers:    response.Header, Body: responseBody,
	}, nil
}

// returns a map of type headers containing both the common headers like authorization Header
// set on the httpClient and a request specific header e.g timeout or Content-Type
func (c *httpClient) getAllRequestHeaders(requestHeaders http.Header) http.Header {
	result := make(http.Header)
	for k, v := range c.builder.headers {
		if len(v) > 0 {
			result.Set(k, v[0])
		}
	}

	// Add custom headers:
	for k, v := range requestHeaders {
		if len(v) > 0 {
			result.Set(k, v[0])
		}
	}

	// set the userAgent if not present.
	if c.builder.userAgent != "" {
		if result.Get("User-Agent") != "" {
			return result
		}
		result.Set("User-Agent", c.builder.userAgent)
	}

	return result
}

// http.Client is created on the first request and all subsequent request use same instance.
func (c *httpClient) getHttpClient() *http.Client {
	c.clientOnce.Do(func() {
		// created by calling SetHTTPClient on builder.
		if c.builder.client != nil {
			c.client = c.builder.client
			return
		}

		c.client = &http.Client{
			// setting this timeout enables us to disable the timeout totally.
			Timeout: c.GetConnectionTimeout() + c.GetResponseTimeout(),
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   c.GetMaxIdleConnections(),
				ResponseHeaderTimeout: c.GetResponseTimeout(),
				DialContext: (&net.Dialer{
					Timeout: c.GetConnectionTimeout(),
				}).DialContext,
			},
		}
	})

	return c.client
}

// check our httpclient if particular fields has been set if not use the default constants.
func (c *httpClient) GetMaxIdleConnections() int {
	if c.builder.maxIdleConnections > 0 {
		return c.builder.maxIdleConnections
	}
	return defaultMaxIdleConnection
}

func (c *httpClient) GetResponseTimeout() time.Duration {
	if c.builder.responseTimeout > 0 {
		return c.builder.responseTimeout
	}

	if c.builder.disableTimeouts {
		return 0
	}

	return defaultResponseTimeout
}

func (c *httpClient) GetConnectionTimeout() time.Duration {
	if c.builder.connectionTimeout > 0 {
		return c.builder.connectionTimeout
	}

	if c.builder.disableTimeouts {
		return 0
	}

	return defaultConnectionTimeout
}
