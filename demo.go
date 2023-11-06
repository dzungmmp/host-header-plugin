// Package plugindemo a demo plugin.
package host_header_plugin

import (
	"context"
	"net/http"
	"strings"
)

// Config the plugin configuration.
type Config struct {
	Headers      map[string]string `json:"headers"`
	AllowedHosts []string          `json:"allowed_hosts"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		Headers: make(map[string]string),
	}
}

// Demo a Demo plugin.
type Demo struct {
	next         http.Handler
	headers      map[string]string
	allowedHosts []string
	name         string
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

	hostHeader := req.Header.Get("Host")
	if len(a.allowedHosts) > 0 {
		// Check allowed hosts
		if !isSliceStringContains(a.allowedHosts, hostHeader) {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

	}

	// For local testing: just allows whoami.localhost in Host
	if !isSliceStringContains([]string{"whoami.localhost"}, hostHeader) {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	for key, value := range a.headers {
		req.Header.Set(key, value)
	}

	req.Header.Set("From-Host-Header", "OK")
	a.next.ServeHTTP(rw, req)
}

func isSliceStringContains(sl []string, val string) bool {
	if len(sl) == 0 {
		return false
	}

	for _, i := range sl {
		if strings.EqualFold(i, val) {
			return true
		}
	}

	return false
}
