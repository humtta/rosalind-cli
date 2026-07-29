package cache

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/humtta/rosalind-cli/internal/problem"
)

const (
	cacheDir      = "rosalind-cli"
	listCacheFile = "problems.json"

	defaultTTL = 24 * time.Hour
)

type listCacheData struct {
	Problems  []problem.Problem `json:"problems"`
	WrittenAt time.Time         `json:"written_at"`
}

type statementCacheData struct {
	Statement problem.ProblemStatement `json:"statement"`
	WrittenAt time.Time                `json:"written_at"`
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
