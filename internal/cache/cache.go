package cache

import (
	"time"
)

type Cache struct {
	path string
	ttl  time.Duration
}
