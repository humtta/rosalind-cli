package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/humtta/rosalind-cli/internal/problem"
)

const (
	cacheDir  = "rosalind-cli"
	cacheFile = "cache.json"

	defaultTTL = 24 * time.Hour
)

type cacheData struct {
	Problems  []problem.Problem `json:"problems"`
	WrittenAt time.Time         `json:"written_at"`
}

type Cache struct {
	path string
	ttl  time.Duration
}

func NewCache() (*Cache, error) {
	dir, err := os.UserCacheDir()
	if err != nil {
		return nil, fmt.Errorf("find cache directory: %w", err)
	}
	return &Cache{
		path: filepath.Join(dir, cacheDir, cacheFile),
		ttl:  defaultTTL,
	}, nil
}

func (c *Cache) Get() ([]problem.Problem, error) {
	raw, err := os.ReadFile(c.path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil
		}
		return nil, fmt.Errorf("read cache: %w", err)
	}

	var data cacheData
	if err := json.Unmarshal(raw, &data); err != nil {
		return nil, fmt.Errorf("decode cache: %w", err)
	}

	if time.Since(data.WrittenAt) > c.ttl {
		return nil, nil
	}

	return data.Problems, nil
}
