package examples

import (
	"time"

	"github.com/miknonny/http-client/gohttp"
)

var (
	httpClient = getHttpClient()
)

func getHttpClient() gohttp.Client {
	client := gohttp.NewBuilder().
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(3 * time.Second).
		Build()

	return client
}

// TODO. Examine the go std library for definition of packages then add more work to your package.
