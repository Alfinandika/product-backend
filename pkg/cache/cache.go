package cache

import "context"

// RedisClientItf is a client to interact with Redis.
type RedisClientItf interface {
	// Get gets the value from redis in []byte form.
	Get(ctx context.Context, key string) ([]byte, error)
	// SetEX sets the value to a key with timeout in seconds.
	SetEX(ctx context.Context, key string, seconds int64, value string) error
	// Exists checks whether the key exists in redis.
	Exists(ctx context.Context, key string) (bool, error)
}
