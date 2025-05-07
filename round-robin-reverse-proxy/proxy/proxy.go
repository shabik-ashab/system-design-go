package proxy

import (
	"io"
	"net/http"
	"net/url"
	"sync"
)

type ReverseProxy struct {
	backends []*url.URL
	curr     int
	mu       sync.Mutex
}

// NewReverseProxy initializes the proxy with a list of backend URLs
func NewReverseProxy(backendUrls []string) *ReverseProxy {
	var parsed []*url.URL
	for _, raw := range backendUrls {
		u, err := url.Parse(raw)
		if err == nil {
			parsed = append(parsed, u)
		}
	}
	return &ReverseProxy{
		backends: parsed,
		curr:     0,
	}
}

// ServeHTTP forwards incoming requests to the selected backend
func (rp *ReverseProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	target := rp.nextBackend()

	// Clone the incoming request to modify it
	req := r.Clone(r.Context())
	req.URL.Scheme = target.Scheme
	req.URL.Host = target.Host
	req.RequestURI = "" // Required for client requests

	// Forward the request
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		http.Error(w, "Bad Gateway", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	// Copy response headers
	for key, values := range resp.Header {
		for _, v := range values {
			w.Header().Add(key, v)
		}
	}

	// Write response status and body
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

// nextBackend picks the next backend using round-robin
func (rp *ReverseProxy) nextBackend() *url.URL {
	rp.mu.Lock()
	defer rp.mu.Unlock()

	backend := rp.backends[rp.curr]
	rp.curr = (rp.curr + 1) % len(rp.backends)
	return backend
}
