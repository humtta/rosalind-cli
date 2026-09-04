// Package cache provides a local file-based cache for Rosalind data.
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
	cacheDir          = "rosalind-cli"
	listCacheFile     = "problems.json"
	statementCacheDir = "statements"

	defaultTTL = 24 * time.Hour
)

// cacheEntry defines a cache entry with methods to get and set its write time.
type cacheEntry interface {
	getWrittenAt() time.Time
	setWrittenAt(time.Time)
}

// cacheMetadata provides the written time metadata for cache entries.
type cacheMetadata struct {
	WrittenAt time.Time `json:"written_at"`
}

// getWrittenAt returns the time the cache entry was written.
func (c *cacheMetadata) getWrittenAt() time.Time { return c.WrittenAt }

// setWrittenAt sets the time the cache entry was written.
func (c *cacheMetadata) setWrittenAt(t time.Time) { c.WrittenAt = t }

// listCacheEntry represents the cache entry for the Rosalind problem list.
type listCacheEntry struct {
	Problems []problem.Problem `json:"problems"`
	cacheMetadata
}

// statementCacheEntry represents the cache entry for a Rosalind problem
// statement.
type statementCacheEntry struct {
	Statement problem.ProblemStatement `json:"statement"`
	cacheMetadata
}

// Cache is a local file-based cache for Rosalind data.
type Cache struct {
	dir string
	ttl time.Duration
}

// NewCache returns a new [Cache] with the default directory and TTL.
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

// GetList returns the cached Rosalind problem list, if available and valid.
func (c *Cache) GetList() ([]problem.Problem, error) {
	var entry listCacheEntry

	found, err := c.read(c.listCachePath(), &entry)
	if err != nil {
		return nil, fmt.Errorf("read problem list: %w", err)
	}
	if !found {
		return nil, nil
	}

	return entry.Problems, nil
}

func (c *Cache) SetList(problems []problem.Problem) error {
	entry := &listCacheEntry{Problems: problems}

	if err := c.write(c.listCachePath(), entry); err != nil {
		return fmt.Errorf("write problem list: %w", err)
	}

	return nil
}

func (c *Cache) GetStatement(id string) (*problem.ProblemStatement, error) {
	var entry statementCacheEntry

	found, err := c.read(c.statementCachePath(id), &entry)
	if err != nil {
		return nil, fmt.Errorf("read problem statement: %w", err)
	}
	if !found {
		return nil, nil
	}

	return &entry.Statement, nil
}

func (c *Cache) SetStatement(
	id string,
	statement problem.ProblemStatement,
) error {
	entry := &statementCacheEntry{Statement: statement}

	if err := c.write(c.statementCachePath(id), entry); err != nil {
		return fmt.Errorf("write problem statement: %w", err)
	}

	return nil
}

// listCachePath returns the path to the list cache file.
func (c *Cache) listCachePath() string {
	return filepath.Join(c.dir, listCacheFile)
}

// statementCachePath returns the path to the statement cache file for the given
// problem ID.
func (c *Cache) statementCachePath(id string) string {
	return filepath.Join(c.dir, statementCacheDir, id+".json")
}

// read reads the cache file at the given path into the given cache entry. The
// boolean reports whether a valid entry was found.
func (c *Cache) read(path string, entry cacheEntry) (bool, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil // Cache missing
		}
		return false, fmt.Errorf("read '%s': %w", path, err)
	}

	if err := json.Unmarshal(data, entry); err != nil {
		return false, fmt.Errorf("unmarshal '%s': %w", path, err)
	}

	if time.Since(entry.getWrittenAt()) > c.ttl {
		return false, nil // Cache expired
	}

	return true, nil
}

// write writes the given cache entry to the cache file at the given path.
func (c *Cache) write(path string, entry cacheEntry) error {
	entry.setWrittenAt(time.Now())

	data, err := json.Marshal(entry)
	if err != nil {
		return fmt.Errorf("marshal entry: %w", err)
	}

	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, 0o700); err != nil {
		return fmt.Errorf("create directory '%s': %w", dir, err)
	}

	tmp, err := os.CreateTemp(dir, ".tmp-*")
	if err != nil {
		return fmt.Errorf("create temporary file: %w", err)
	}

	defer tmp.Close()
	defer os.Remove(tmp.Name())

	if _, err := tmp.Write(data); err != nil {
		return fmt.Errorf("write '%s': %w", tmp.Name(), err)
	}

	if err := tmp.Close(); err != nil {
		return fmt.Errorf("close '%s': %w", tmp.Name(), err)
	}

	if err := os.Rename(tmp.Name(), path); err != nil {
		return fmt.Errorf("rename '%s' to '%s': %w", tmp.Name(), path, err)
	}

	return nil
}
