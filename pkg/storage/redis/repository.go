package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/kmchary/urlshortner/pkg/urlshortener"
)

type Storage struct {
	client *redis.Client
}

func (s *Storage) Get(ctx context.Context, key string) string {
	value, _ := s.client.Get(ctx, key).Result()
	return value
}

func (s *Storage) Set(ctx context.Context, key string, value string) error {
	return s.client.Set(ctx, key, value, 0).Err()
}

func newRedisClient(ctx context.Context, redisURL string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     redisURL, // use default Addr
		Password: "",       // no password set
		DB:       0,        // use default DB
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewRedisRepository(ctx context.Context, redisURL string) (urlshortener.Repository, error) {
	repo := &Storage{}
	client, err := newRedisClient(ctx, redisURL)
	if err != nil {
		return nil, err
	}
	repo.client = client
	return repo, nil
}
