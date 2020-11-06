package util

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type RedisConfig struct {
	Addrs    string `toml:"addrs"`
	Pwd      string `toml:"pwd"`
	PoolSize int    `toml:"pool_size"`
}

type RedisClient struct {
	RedisInstance *redis.Client
}

func NewRedisClient(c *RedisConfig) (*RedisClient, error) {
	options := &redis.Options{
		Password:     c.Pwd,
		PoolSize:     c.PoolSize,
		MinIdleConns: 10,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
		PoolTimeout:  2 * time.Second,
		IdleTimeout:  300 * time.Second,
	}
	redisClient := RedisClient{redis.NewClient(options)}
	return &redisClient, nil
}

func (client *RedisClient) Set(key string, value string, expire int) error {
	err := client.RedisInstance.Set(ctx, key, value, time.Duration(expire)*time.Second).Err()
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func (client *RedisClient) Get(key string) (string, error) {
	val, err := client.RedisInstance.Get(ctx, key).Result()
	if err != nil {
		log.Fatal(err)
	}
	return val, err
}
