package cache

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/humtta/rosalind-cli/internal/model"
)

const (
	cacheSubpath = "/rosalind-cli/cache.json"
	ttl          = 24 * time.Hour
)

type cacheData struct {
	Problems  []model.Problem `json:"problems"`
	WrittenAt time.Time       `json:"written_at"`
}

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
