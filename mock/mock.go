package httpmock

import (
	"fmt"
	"net/http"

	gohttp_types "github.com/miknonny/http-client/types"
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
func (m *Mock) GetResponse() (*gohttp_types.Response, error) {
	if m.Error != nil {
		return nil, m.Error
	}

	response := gohttp_types.Response{
		Status:     fmt.Sprintf("%d %s", m.ResponseStatusCode, http.StatusText(m.ResponseStatusCode)),
		StatusCode: m.ResponseStatusCode,
		Body:       []byte(m.ResponseBody),
	}

	return &response, nil
}
