package client

import (
	"context"
	"github.com/alexeyzer/user-api/config"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type RedisClient interface {
	Set(ctx context.Context, key, value string) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
}

type redisClient struct {
	redisClient *redis.Client
}

func (c *redisClient) Delete(ctx context.Context, key string) error {
	err := c.redisClient.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *redisClient) Set(ctx context.Context, key, value string) error {
	err := c.redisClient.Set(ctx, key, value, time.Hour*24).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c *redisClient) Get(ctx context.Context, key string) (string, error) {
	val, err := c.redisClient.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", status.Errorf(codes.Unauthenticated, "sessionID does`t exists")
		}
		return "", err
	}
	return val, nil
}

func NewRedisClient(ctx context.Context) (RedisClient, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Host + ":" + config.Config.Redis.Port,
		Password: config.Config.Redis.Password,
		DB:       0, // use default DB
	})
	err := rdb.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}
	return &redisClient{redisClient: rdb}, nil
}
