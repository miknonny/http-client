package gohttp

import (
	"fmt"
	"net/http"
)

// We will be using method + url + RequestBody to create a mocking key.
// we do not use headers because we might have common headers or different
// headers for thesame set of request.

// The mock structure provides a clean way  to configure http mocks  based on the
// combination between  request Method, URL and request Body.
type Mock struct {
	// request
	Method      string
	Url         string
	RequestBody string

	// response
	Error              error
	ResponseBody       string
	ResponseStatusCode int
}

// GetResponse Returns the response object based on the mock configuration.
func (m *Mock) GetResponse() (*Response, error) {
	if m.Error != nil {
		return nil, m.Error
	}

	response := Response{
		status:     fmt.Sprintf("%d %s", m.ResponseStatusCode, http.StatusText(m.ResponseStatusCode)),
		statusCode: m.ResponseStatusCode,
		body:       []byte(m.ResponseBody),
	}

	return &response, nil
}
