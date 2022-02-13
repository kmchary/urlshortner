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
	urlHandler := rest.NewHandler(urlShortener)
	urlHandler.InitRoutes()

	log.Fatal(http.ListenAndServe("localhost:8081", urlHandler.Router))
}
