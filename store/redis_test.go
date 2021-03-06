package store

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestNewRedisStore(t *testing.T) {

	s := NewRedisStore(redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	}), "test:", 5*time.Minute)
	assert.IsType(t, &redisStore{}, s)
}

func TestRedisStore_Get(t *testing.T) {

	c := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	id := fmt.Sprintf("%d", rand.Int63())
	value := fmt.Sprintf("%d", rand.Int63())
	s := NewRedisStore(c, "test:", 5*time.Minute)
	c.Set("test:"+id, value, 5*time.Minute)
	result, err := s.Get(id, false)
	assert.Nil(t, err)
	assert.Equal(t, value, result)
	assert.Equal(t, value, c.Get("test:" + id).Val())
}

func TestRedisStore_GetAndClear(t *testing.T) {

	c := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	id := fmt.Sprintf("%d", rand.Int63())
	value := fmt.Sprintf("%d", rand.Int63())
	s := NewRedisStore(c, "test:", 5*time.Minute)
	c.Set("test:"+id, value, 5*time.Minute)
	result, err := s.Get(id, true)
	assert.Nil(t, err)
	assert.Equal(t, value, result)
	assert.Equal(t, int64(0), c.Exists("test:" + id).Val())
}

func TestNewRedisStoreByClient(t *testing.T) {

	s := NewRedisStore(redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	}), "test:", 5*time.Minute)
	assert.IsType(t, &redisStore{}, s)
}

func TestRedisStore_Set(t *testing.T) {
	c := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	id := fmt.Sprintf("%d", rand.Int63())
	value := fmt.Sprintf("%d", rand.Int63())
	s := NewRedisStore(redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	}), "test:", 5*time.Minute)
	err := s.Set(id, value)
	assert.Nil(t, err)
	verifyCode, err := c.Get("test:" + id).Result()
	assert.Nil(t, err)
	assert.Equal(t, value, verifyCode)
}

func TestRedisStore_GetFailed(t *testing.T) {

	id := fmt.Sprintf("%d", rand.Int63())
	s := NewRedisStore(redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	}), "test:", 5*time.Minute)
	value, err := s.Get(id, false)
	assert.NotNil(t, err)
	assert.Equal(t, "", value)
}

func TestRedisStore_GetFailed2(t *testing.T) {

	id := fmt.Sprintf("%d", rand.Int63())
	s := NewRedisStore(redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	}), "test:", 5*time.Minute)
	value, err := s.Get(id, true)
	assert.NotNil(t, err)
	assert.Equal(t, "", value)
}
