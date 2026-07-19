package cache

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const cacheSubpath = "/rosalind-cli/cache.json"

type Cache struct {
	path string
	ttl  time.Duration
}

func cachePath() (string, error) {
	dir, err := os.UserCacheDir()
	if err != nil {
		return "", fmt.Errorf("locate user cache dir: %w", err)
	}
	return filepath.Join(dir, cacheSubpath), nil
}
