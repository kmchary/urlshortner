package redis

import (
	"github.com/go-redis/redis"
	"github.com/kmchary/urlshortner/pkg/urlshortener"
)

type Storage struct {
	client *redis.Client
}

func (s *Storage) Get(key string) string {
	value, _ := s.client.Get(key).Result()
	return value
}

func (s *Storage) Set(key string, value string) error {
	return s.client.Set(key, value, 0).Err()
}

func newRedisClient(redisURL string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     redisURL, // use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewRedisRepository(redisURL string) (urlshortener.Repository, error) {
	repo := &Storage{}
	client, err := newRedisClient(redisURL)
	if err != nil {
		return nil, err
	}
	repo.client = client
	return repo, nil
}
