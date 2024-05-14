package utils

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var ctx = context.Background()

func NewRedisClient(lc fx.Lifecycle) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.address"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	})

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return rdb.Close()
		},
	})

	return rdb
}

func GetContext() context.Context {
	return ctx
}
