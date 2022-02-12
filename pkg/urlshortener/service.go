package urlshortener

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"math/big"
)

type Service interface {
	ShortenURL(url string, userId string) (string, error)
}

type Repository interface {
	Get(ctx context.Context, key string) string
	Set(ctx context.Context, key string, value string) error
}

type service struct {
	repo Repository
}

func (s *service) ShortenURL(url string, userId string) (string, error) {
	ctx := context.Background()
	hf := sha256.New()
	_, err := hf.Write([]byte(url + userId))
	if err != nil {
		return "", errors.New("failed to generate hash")
	}

	hashBytes := hf.Sum(nil)

	redisKey := string(hashBytes)
	shortUrl := s.repo.Get(ctx, redisKey)
	log.Println("repo.Get shortUrl", shortUrl)
	if shortUrl != "" {
		return shortUrl, nil
	}

	bigNumber := new(big.Int).SetBytes(hashBytes).Uint64()
	shortUrl = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d", bigNumber)))

	if err := s.repo.Set(ctx, redisKey, shortUrl); err != nil {
		return "", errors.New("failed to store in the cache")
	}
	log.Println("repo.Set shortUrl", shortUrl)
	return shortUrl, nil
}

func NewService(repo Repository) Service {
	return &service{repo}
}
