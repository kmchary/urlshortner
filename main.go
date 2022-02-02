package main

import (
	"fmt"
	"github.com/kmchary/urlshortner/service"
)

func main() {
	urlShortener := service.NewUrlShortener()
	url := "www.google.com/kmchary/test/something/something/text.html"
	userId := "kmchary"
	shortUrl, err := urlShortener.ShortenURL(url, userId)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("original url", url)
	fmt.Println("shortened url", "localhost:8080"+shortUrl)
}