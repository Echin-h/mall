package cache

import "testing"

func TestInitCache(t *testing.T) {
	InitCache()
	if RedisClient == nil {
		t.Errorf("RedisClient is nil")
	}
}
