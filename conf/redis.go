package conf

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"time"
)

type RedisClient struct {
}

var rdClient *redis.Client
var nDuration = 30 * time.Hour

func InitRedis() (*RedisClient, error) {
	rdClient = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})

	_, err := rdClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return &RedisClient{}, nil
}

func (rc *RedisClient) Set(key string, value any) error {
	return rdClient.Set(context.Background(), key, value, nDuration).Err()
}

func (rc *RedisClient) Get(key string) (any, error) {
	return rdClient.Get(context.Background(), key).Result()
}

func (rc *RedisClient) Delete(key ...string) error {
	return rdClient.Del(context.Background(), key...).Err()
}
