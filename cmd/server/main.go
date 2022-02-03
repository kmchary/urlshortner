package main

import (
	"fmt"
	"github.com/kmchary/urlshortner/pkg/storage/redis"
	"github.com/kmchary/urlshortner/pkg/urlshortener"
	"log"
)

func main() {
	redisRepo, err := redis.NewRedisRepository("localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	urlShortener := urlshortener.NewService(redisRepo)
	shortUrl, err := urlShortener.ShortenURL("www.google.com/something/something/something/text.htm", "kmchary")
	if err != nil {
		log.Fatal(err)
	}
	shortUrl = "localhost:8080/"+shortUrl
	fmt.Println("short url", shortUrl)
}
