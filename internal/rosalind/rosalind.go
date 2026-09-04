// Package rosalind provides access to Rosalind problems.
package rosalind

import (
	"fmt"

	"github.com/humtta/rosalind-cli/internal/cache"
	"github.com/humtta/rosalind-cli/internal/client"
)

// Rosalind provides access to Rosalind problems.
type Rosalind struct {
	client *client.Client
	cache  *cache.Cache
}

// NewRosalind returns a new [Rosalind] with the default client and cache.
func NewRosalind() (*Rosalind, error) {
	cache, err := cache.NewCache()
	if err != nil {
		return nil, fmt.Errorf("init cache: %w", err)
	}

	return &Rosalind{
		client: client.NewClient(),
		cache:  cache,
	}, nil
}
