package client

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	defaultBaseURL = "https://rosalind.info"
	defaultTimeout = 10 * time.Second

	listEndpoint    = "/problems/list-view"
	problemEndpoint = "/problems"
)

type Client struct {
	httpClient *http.Client
	baseURL    string
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{Timeout: defaultTimeout},
		baseURL:    defaultBaseURL,
	}
}

func (c *Client) BaseURL() string {
	return c.baseURL
}

func (c *Client) GetList() ([]byte, error) {
	return c.get(listEndpoint)
}

func (c *Client) GetProblem(id string) ([]byte, error) {
	if id == "" {
		return nil, fmt.Errorf("problem id is empty")
	}
	return c.get(problemEndpoint, id)
}

func (c *Client) get(segments ...string) ([]byte, error) {
	reqURL, err := url.JoinPath(c.baseURL, segments...)
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
