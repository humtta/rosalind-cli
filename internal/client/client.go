package client

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	baseURL = "https://rosalind.info"
	timeout = 10 * time.Second

	problemListPageEndpoint = "/problems/list-view"
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

func (c *Client) GetProblemListPage() ([]byte, error) {
	return c.Get(problemListPageEndpoint)
}

func (c *Client) Get(elem ...string) ([]byte, error) {
	reqURL, err := url.JoinPath(c.baseURL, elem...)
	if err != nil {
		return nil, fmt.Errorf("build request URL: %w", err)
	}

	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("create GET request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("GET %s: %w", reqURL, err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET %s: %s", reqURL, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body from %s: %w", reqURL, err)
	}

	return body, nil
}
