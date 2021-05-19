package helpers

import (
	"context"
	"io"
	"net/http"
	"time"
)

// HttpClient
type HttpClient struct {
	Context context.Context
	URL     string
	Body    io.Reader
	Method  string
	Headers map[string]interface{}
}

// Get
func (h HttpClient) Get() (*http.Response, error) {
	client := http.Client{
		Timeout: time.Second * 3,
	}

	r, err := http.NewRequestWithContext(h.Context, h.Method, h.URL, h.Body)
	if err != nil {
		return nil, err
	}

	// TODO: add custom headers

	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Post
func (h HttpClient) Post() (*http.Response, error) {
	client := http.Client{
		Timeout: time.Second * 3,
	}

	r, err := http.NewRequestWithContext(h.Context, h.Method, h.URL, h.Body)
	if err != nil {
		return nil, err
	}

	// TODO: add custom headers

	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Patch
func (h HttpClient) Patch() (*http.Response, error) {
	client := http.Client{
		Timeout: time.Second * 3,
	}

	r, err := http.NewRequestWithContext(h.Context, h.Method, h.URL, h.Body)
	if err != nil {
		return nil, err
	}

	// TODO: add custom headers

	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
