package cache

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const (
	cacheSubpath = "/rosalind-cli/cache.json"
	ttl          = 24 * time.Hour
)

type Cache struct {
	path string
	ttl  time.Duration
}

func NewCache() (*Cache, error) {
	path, err := cachePath()
	if err != nil {
		return nil, fmt.Errorf("get cache path: %w", err)
	}
	return &Cache{
		path: path,
		ttl:  ttl,
	}, nil
}

func cachePath() (string, error) {
	dir, err := os.UserCacheDir()
	if err != nil {
		return "", fmt.Errorf("locate user cache dir: %w", err)
	}
	return filepath.Join(dir, cacheSubpath), nil
}
