package redis

import (
	"context"
	"fmt"
	"github.com/go-related/redis/settings"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type RedisService struct {
	client *redis.Client
}

func New() *RedisService {
	addr := fmt.Sprintf("%s:%s", settings.ApplicationConfiguration.Redis.Host, settings.ApplicationConfiguration.Redis.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: settings.ApplicationConfiguration.Redis.Password,
		DB:       settings.ApplicationConfiguration.Redis.Db,
	})

	return &RedisService{rdb}
}

func (rs *RedisService) PublishToChannel(ctx context.Context, channelName string, payload []byte) error {

	err := rs.client.Publish(ctx, channelName, payload).Err()
	if err != nil {
		logrus.WithError(err).Error("failed to send data to channel")
	}
	return err
}

func (rs *RedisService) SubscribeToChannel(ctx context.Context, channelName string) *redis.PubSub {
	return rs.client.Subscribe(ctx, channelName)
}

func (rs *RedisService) UnSubscribeToChannel(ctx context.Context, channelName string, data *redis.PubSub) error {
	if data != nil {
		return data.Unsubscribe(ctx, channelName)
	}
	return nil
}
