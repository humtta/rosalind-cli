// Package client provides an HTTP client for fetching Rosalind pages.
package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/humtta/rosalind-cli/internal/problem"
)

const (
	defaultBaseURL = "https://rosalind.info"
	defaultTimeout = 10 * time.Second

	listEndpoint    = "/problems/list-view"
	problemEndpoint = "/problems"

	userAgent = "rosalind-cli"
)

// Client is an HTTP client for fetching Rosalind pages.
type Client struct {
	httpClient *http.Client
	baseURL    string
}

// NewClient returns a new [Client] with the default base URL and timeout.
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{Timeout: defaultTimeout},
		baseURL:    defaultBaseURL,
	}
}

// BaseURL returns the [Client]'s base URL.
func (c *Client) BaseURL() string { return c.baseURL }

// GetList fetches the Rosalind problem list page.
func (c *Client) GetList(ctx context.Context) ([]byte, error) {
	return c.get(ctx, listEndpoint)
}

// GetProblem fetches the Rosalind problem page for the given problem ID.
func (c *Client) GetProblem(ctx context.Context, id string) ([]byte, error) {
	if err := problem.ValidateID(id); err != nil {
		return nil, fmt.Errorf("invalid id '%s': %w", id, err)
	}
	return c.get(ctx, problemEndpoint, id)
}

// get performs a GET request to the base URL joined with the given path
// segments.
func (c *Client) get(ctx context.Context, segments ...string) ([]byte, error) {
	url, err := url.JoinPath(c.baseURL, segments...)
	if err != nil {
		return nil, fmt.Errorf("build URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}

	req.Header.Set("User-Agent", userAgent)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("get '%s': %w", url, err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get '%s': unexpected status '%s'", url, res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("get '%s': read body: %w", url, err)
	}

	return body, nil
}
