package gohttp

import (
	"net/http"
	"sync"

	gohttp_types "github.com/miknonny/http-client/v3/types"
)

type httpClient struct {
	builder    *clientBuilder
	client     *http.Client
	clientOnce sync.Once
}

type Client interface {
	Get(url string, headers http.Header) (*gohttp_types.Response, error)
	Post(url string, headers http.Header, body interface{}) (*gohttp_types.Response, error)
	Put(url string, headers http.Header, body interface{}) (*gohttp_types.Response, error)
	Patch(url string, headers http.Header, body interface{}) (*gohttp_types.Response, error)
	Delete(url string, headers http.Header) (*gohttp_types.Response, error)
	Options(url string, headers http.Header) (*gohttp_types.Response, error)
}

func (c *httpClient) Get(url string, headers http.Header) (*gohttp_types.Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}

func (c *httpClient) Post(url string, headers http.Header, body interface{}) (*gohttp_types.Response, error) {
	return c.do(http.MethodPost, url, headers, body)
}

func (c *httpClient) Put(url string, headers http.Header, body interface{}) (*gohttp_types.Response, error) {
	return c.do(http.MethodPut, url, headers, body)
}

func (c *httpClient) Patch(url string, headers http.Header, body interface{}) (*gohttp_types.Response, error) {
	return c.do(http.MethodPatch, url, headers, body)
}

func (c *httpClient) Delete(url string, headers http.Header) (*gohttp_types.Response, error) {
	return c.do(http.MethodDelete, url, headers, nil)
}

func (c *httpClient) Options(url string, headers http.Header) (*gohttp_types.Response, error) {
	return c.do(http.MethodOptions, url, headers, nil)
}
