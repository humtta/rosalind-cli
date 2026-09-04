// Package rosalind provides access to Rosalind problems.
package rosalind

import (
	"github.com/humtta/rosalind-cli/internal/cache"
	"github.com/humtta/rosalind-cli/internal/client"
)

// Rosalind provides access to Rosalind problems.
type Rosalind struct {
	client *client.Client
	cache  *cache.Cache
}
