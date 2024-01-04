package gohttp

import (
	"net/http"
	"time"
)

// The clientBuilder is simply a way the configure a our http.Client with sensible defaults. so the builder is
// just a struct with methods to change the config
// why do we use a builder?
// so we can seperate our interface for building a client from the interface for making request with the built client.
type clientBuilder struct {
	headers http.Header

	maxIdleConnections int

	// timeout field on net.Dial; or time to wait for creating a new connection.
	connectionTimeout time.Duration

	// ResponseHeaderTimeout on http.Transport
	responseTimeout time.Duration

	disableTimeouts bool
}

type ClientBuilder interface {
	SetHeaders(headers http.Header) ClientBuilder
	SetConnectionTimeout(timeout time.Duration) ClientBuilder
	SetResponseTimeout(timeout time.Duration) ClientBuilder
	SetMaxIdleConnections(i int) ClientBuilder
	DisableTimeouts(disable bool) ClientBuilder
	Build() Client
}

func NewBuilder() ClientBuilder {
	// note the we can still set sensible defaults in our builder config.
	return &clientBuilder{}
}

// Here we take all the configuration from our config/builder and create the httpClient.
func (b *clientBuilder) Build() Client {
	return &httpClient{
		builder: b,
	}
}

func (b *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	b.headers = headers

	return b
}

func (b *clientBuilder) SetConnectionTimeout(timeout time.Duration) ClientBuilder {
	b.connectionTimeout = timeout

	return b
}

func (b *clientBuilder) SetResponseTimeout(timeout time.Duration) ClientBuilder {
	b.responseTimeout = timeout

	return b
}

func (b *clientBuilder) SetMaxIdleConnections(connections int) ClientBuilder {
	b.maxIdleConnections = connections

	return b
}

func (b *clientBuilder) DisableTimeouts(disable bool) ClientBuilder {
	b.disableTimeouts = disable

	return b
}
