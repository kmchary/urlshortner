package urlshortener

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"time"
)

var FailedToStoreError = errors.New("failed to store in the cache")

type Service interface {
	GenerateURLUsingBase62RandomChars(url string, userId string) (string, error)
	GetActualUrl(shortUrl string) string
}

type Repository interface {
	Get(ctx context.Context, key string) string
	Set(ctx context.Context, key string, value string) error
	HGetAll(ctx context.Context, key string) map[string]string
	HSet(ctx context.Context, key string, data map[string]interface{}) error
}

type RandomCharsGenerator interface {
	GetRandomChars(len int) string
}

type service struct {
	repo Repository
	rcg  RandomCharsGenerator
}

func NewService(repo Repository) Service {
	return &service{repo, NewRandomCharsGenerator()}
}

type base62RandomChars struct {
}

func NewRandomCharsGenerator() RandomCharsGenerator {
	return &base62RandomChars{}
}

func (b *base62RandomChars) GetRandomChars(length int) string {
	const ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	var url []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		url = append(url, ALPHABET[r.Intn(62)])
	}
	return string(url)
}

func (s *service) GenerateURLUsingBase62RandomChars(url string, userId string) (string, error) {
	ctx := context.Background()
	var shortUrl string

	redisKey := url + userId
	redisData := s.repo.HGetAll(ctx, redisKey)
	log.Println("repo.HGetAll shortUrl", redisData["SHORT_URL"])

	if shortUrl, ok := redisData["SHORT_URL"]; ok {
		return shortUrl, nil
	}

	for {
		shortUrl = s.rcg.GetRandomChars(7)
		log.Println("GetRandomChars shortUrl", shortUrl)
		existingUrl := s.repo.Get(ctx, shortUrl)
		log.Println("existingUrl:", existingUrl)
		if existingUrl == "" {
			log.Println("storing key,value:", shortUrl, url)
			s.repo.Set(ctx, shortUrl, url)
			break
		}
	}

	data := map[string]interface{}{
		"SHORT_URL": shortUrl,
		"URL":       url,
	}

	if err := s.repo.HSet(ctx, redisKey, data); err != nil {
		log.Println(FailedToStoreError)
		log.Println(err)
		return "", FailedToStoreError
	}
	log.Println("repo.HMSet shortUrl", shortUrl)
	return shortUrl, nil
}

func (s *service) GetActualUrl(shortUrl string) string {
	ctx := context.Background()
	url := s.repo.Get(ctx, shortUrl)
	log.Println("actual url:", url)
	return url
}
