// Package plugindemo a demo plugin.
package host_header_plugin

import (
	"context"
	"net/http"
)

// Config the plugin configuration.
type Config struct {
	Headers map[string]string `json:"headers"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		Headers: make(map[string]string),
	}
}

// Demo a Demo plugin.
type Demo struct {
	next    http.Handler
	headers map[string]string
	name    string
}

// New created a new Demo plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	// if len(config.Headers) == 0 {
	// 	return nil, fmt.Errorf("headers cannot be empty")
	// }

	return &Demo{
		headers: config.Headers,
		next:    next,
		name:    name,
	}, nil
}

func (a *Demo) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	for key, value := range a.headers {
		_ = value
		if key == "Host" {
			req.Header.Set(key, "test.test")
		}
	}

	a.next.ServeHTTP(rw, req)
}
