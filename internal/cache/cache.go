package cache

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/humtta/rosalind-cli/internal/problem"
)

const (
	cacheDir          = "rosalind-cli"
	listCacheFile     = "problems.json"
	statementCacheDir = "statements"

	defaultTTL = 24 * time.Hour
)

type cacheEntry interface {
	getWrittenAt() time.Time
	setWrittenAt(time.Time)
}

type cacheMetadata struct {
	WrittenAt time.Time `json:"written_at"`
}

func (c *cacheMetadata) getWrittenAt() time.Time { return c.WrittenAt }

func (c *cacheMetadata) setWrittenAt(t time.Time) { c.WrittenAt = t }

type listCacheEntry struct {
	Problems []problem.Problem `json:"problems"`
	cacheMetadata
}

type statementCacheEntry struct {
	Statement problem.ProblemStatement `json:"statement"`
	cacheMetadata
}

type Cache struct {
	dir string
	ttl time.Duration
}

func NewCache() (*Cache, error) {
	dir, err := os.UserCacheDir()
	if err != nil {
		return nil, fmt.Errorf("find cache directory: %w", err)
	}
	return &Cache{
		dir: filepath.Join(dir, cacheDir),
		ttl: defaultTTL,
	}, nil
}

// listCachePath returns the path to the list cache file.
func (c *Cache) listCachePath() string {
	return filepath.Join(c.dir, listCacheFile)
}

// statementCachePath returns the path to the statement cache file for the given
// problem ID.
func (c *Cache) statementCachePath(id string) (string, error) {
	if err := problem.ValidateID(id); err != nil {
		return "", fmt.Errorf("invalid id '%s': %w", id, err)
	}
	return filepath.Join(c.dir, statementCacheDir, id+".json"), nil
}
