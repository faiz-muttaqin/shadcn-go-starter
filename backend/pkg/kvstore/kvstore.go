package kvstore

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/pkg/clr"
	"github.com/go-redis/redis/v8"
)

var (
	RDB       *redis.Client
	redisUp   atomic.Bool
	shardMaps [256]*sync.Map
)

type valueWithTTL struct {
	value string
	ttl   time.Time
}

func init() {
	for i := range shardMaps {
		shardMaps[i] = &sync.Map{}
	}
	redisUp.Store(false) // Initialize redisUp to false
}

// InitRedis initializes and returns a Redis client
func InitRedis(addr, password string, db any) *redis.Client {
	var dbInt int

	switch v := db.(type) {
	case int: // already int
		dbInt = v
	case int8, int16, int32, int64: // other integer types
		dbInt = int(fmt.Sprintf("%v", v)[0]) // or simpler: dbInt = int(reflect.ValueOf(v).Int())
	case uint, uint8, uint16, uint32, uint64:
		dbInt = int(fmt.Sprintf("%v", v)[0]) // same as above
	case string: // parse string to int
		if n, err := strconv.Atoi(v); err == nil {
			dbInt = n
		} else {
			dbInt = 0
		}
	default: // unsupported type
		dbInt = 0
	}

	RDB := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       dbInt,
	})

	// Ping the Redis server to check the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		redisUp.Store(false)
		fmt.Printf("failed to connect to Redis: %v", err)
	}

	redisUp.Store(true)
	fmt.Println(clr.TextBlue("Connected to the Redis server successfully"))
	return RDB
}

func getShard(key string) *sync.Map {
	return shardMaps[uint(fnv32(key))%uint(len(shardMaps))]
}

func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	for i := 0; i < len(key); i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}

// SetKey sets a key with a value and an expiration time
func SetKey(key string, value string, ttl time.Duration) error {
	if redisUp.Load() {
		err := RDB.Set(context.Background(), key, value, ttl).Err()
		if err == nil {
			return nil
		}
		redisUp.Store(false)
	}

	shard := getShard(key)
	expiration := time.Now().Add(ttl)
	shard.Store(key, valueWithTTL{value: value, ttl: expiration})
	time.AfterFunc(ttl, func() {
		shard.Delete(key)
	})
	return nil
}

// GetKey retrieves a value by key
func GetKey(key string) (string, error) {
	if redisUp.Load() {
		val, err := RDB.Get(context.Background(), key).Result()
		if err == nil {
			return val, nil
		}
		redisUp.Store(false)
	}

	shard := getShard(key)
	if val, ok := shard.Load(key); ok {
		v := val.(valueWithTTL)
		if time.Now().Before(v.ttl) {
			return v.value, nil
		}
		shard.Delete(key)
	}
	return "", fmt.Errorf("key not found")
}

// ExistsIn checks if a key exists
func ExistsIn(key string) (bool, error) {
	if redisUp.Load() {
		count, err := RDB.Exists(context.Background(), key).Result()
		if err == nil {
			return count > 0, nil
		}
		redisUp.Store(false)
	}

	shard := getShard(key)
	if val, ok := shard.Load(key); ok {
		v := val.(valueWithTTL)
		if time.Now().Before(v.ttl) {
			return true, nil
		}
		shard.Delete(key)
	}
	return false, nil
}

// ExtendKeyTTL extends the expiration of a key
func ExtendKeyTTL(key string, ttl time.Duration) error {
	if redisUp.Load() {
		err := RDB.Expire(context.Background(), key, ttl).Err()
		if err == nil {
			return nil
		}
		redisUp.Store(false)
	}

	shard := getShard(key)
	if val, ok := shard.Load(key); ok {
		v := val.(valueWithTTL)
		v.ttl = time.Now().Add(ttl)
		shard.Store(key, v)
		time.AfterFunc(ttl, func() {
			shard.Delete(key)
		})
		return nil
	}
	return fmt.Errorf("key not found")
}

// DeleteKey removes a key
func DeleteKey(key string) error {
	if redisUp.Load() {
		err := RDB.Del(context.Background(), key).Err()
		if err == nil {
			return nil
		}
		redisUp.Store(false)
	}

	shard := getShard(key)
	shard.Delete(key)
	return nil
}

// DeleteKeysWithPrefix removes all keys that start with the given prefix
func DeleteKeysWithPrefix(prefix string) error {
	if redisUp.Load() {
		ctx := context.Background()
		iter := RDB.Scan(ctx, 0, prefix+"*", 0).Iterator()
		for iter.Next(ctx) {
			err := RDB.Del(ctx, iter.Val()).Err()
			if err != nil {
				redisUp.Store(false)
				break
			}
		}
		if err := iter.Err(); err == nil {
			return nil
		}
	} else {
		for _, shard := range shardMaps {
			shard.Range(func(key, value interface{}) bool {
				if k, ok := key.(string); ok && len(k) >= len(prefix) && k[:len(prefix)] == prefix {
					shard.Delete(key)
				}
				return true
			})
		}
	}

	return nil
}

// GetKeyTTL retrieves the remaining expiration time of a key
func GetKeyTTL(key string) (time.Duration, error) {
	if redisUp.Load() {
		ttl, err := RDB.TTL(context.Background(), key).Result()
		if err == nil {
			return ttl, nil
		}
		redisUp.Store(false)
	}

	shard := getShard(key)
	if val, ok := shard.Load(key); ok {
		v := val.(valueWithTTL)
		if time.Now().Before(v.ttl) {
			return time.Until(v.ttl), nil
		}
		shard.Delete(key)
	}
	return 0, fmt.Errorf("key not found")
}
