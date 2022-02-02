package service

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"math/big"
)

type UrlShortener interface {
	ShortenURL(url string, userId string) (string, error)
}

type urlShortener struct {
}

func (us *urlShortener) ShortenURL(url string, userId string) (string, error)  {
	hf := sha256.New()
	_, err := hf.Write([]byte( url + userId ))
	if err != nil {
		return "", errors.New("failed to generate hash")
	}
	hashBytes := hf.Sum(nil)
	bigNumber := new(big.Int).SetBytes(hashBytes).Uint64()
	shortUrl := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d",bigNumber)))
	return shortUrl, nil
}

func NewUrlShortener() *urlShortener {
	return &urlShortener{}
}