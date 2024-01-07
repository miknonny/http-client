package gohttp_types

import (
	"encoding/json"
	"net/http"
)

// Response defines a http Server response.
// Response is put in a seperate package `gohttp_types` to avoid cyclic dependency `go_http` and `httpmock_server` packages.
type Response struct {
	Status     string
	StatusCode int
	Headers    http.Header
	Body       []byte
}

func (r *Response) Bytes() []byte {
	return r.Body
}

func (r *Response) String() string {
	return string(r.Body)
}

func (r *Response) UnmarshalJson(target interface{}) error {
	return json.Unmarshal(r.Bytes(), target)
}
