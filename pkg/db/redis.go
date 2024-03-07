package db

import (
	"WarpGPT/pkg/env"
	"WarpGPT/pkg/logger"
	"context"
)

func GetRedisClient() (*redis.Client, error) {
	logger.Log.Info("RedisAddress为：", env.E.RedisAddress)
	if env.E.RedisAddress == "" {
		logger.Log.Info("不启动redis")
		return nil, nil
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr:           env.E.RedisAddress,
		Password:       env.E.RedisPasswd,
		DB:             env.E.RedisDB,
		MaxRetries:     3,
		MaxActiveConns: 20,
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	logger.Log.Info("成功连接到Redis")

	return redisClient, nil
}
