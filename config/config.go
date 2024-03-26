package config

import (
	"sync"
)

var (
	mu     sync.Mutex
	keyIdx int
)

func GetNextAPIKey() string {
	apiKeys := GetConfigValues().YoutubeApiKeys

	mu.Lock()
	defer mu.Unlock()
	key := apiKeys[keyIdx]
	keyIdx = (keyIdx + 1) % len(apiKeys)
	return key
}
