package client

import (
	"net/http"
	"time"
)

const (
	baseURL = "https://rosalind.info"
	timeout = 10 * time.Second
)

type Client struct {
	httpClient *http.Client
	baseURL    string
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: timeout,
		},
		baseURL: baseURL,
	}
}
