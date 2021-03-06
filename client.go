package resource

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

// HTTPClient is the interface for the http client used by the resource package.
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// NewClient is the interface for the resource client.
func NewClient(skipTLSValidation bool) HTTPClient {
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: skipTLSValidation,
		},
	}

	return &http.Client{
		Transport: transport,
	}
}
