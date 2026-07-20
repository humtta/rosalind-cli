package cache

import (
	"encoding/json"
	"errors"
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

func (c *Cache) Get() ([]model.Problem, error) {
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

func (c *Cache) Set(problems []model.Problem) error {
	data := cacheData{
		Problems:  problems,
		WrittenAt: time.Now(),
	}

	raw, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("encode cache: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(c.path), 0o755); err != nil {
		return fmt.Errorf("create cache dir: %w", err)
	}

	if err := os.WriteFile(c.path, raw, 0o644); err != nil {
		return fmt.Errorf("write cache: %w", err)
	}

	return nil
}
