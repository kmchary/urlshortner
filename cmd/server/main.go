package main

import (
	"context"
	"github.com/kmchary/urlshortner/pkg/http/rest"
	"github.com/kmchary/urlshortner/pkg/storage/redis"
	"github.com/kmchary/urlshortner/pkg/urlshortener"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	redisRepo, err := redis.NewRedisRepository(ctx, "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	urlShortener := urlshortener.NewService(redisRepo)
	shortUrl, err := urlShortener.ShortenURL("www.google.com/something/something/something/text.htm", "kmchary")
	if err != nil {
		log.Fatal(err)
	}
	shortUrl = "localhost:8080/" + shortUrl
	r := rest.NewHandler(urlShortener)
	log.Fatal(http.ListenAndServe("localhost:8080", r))

}
