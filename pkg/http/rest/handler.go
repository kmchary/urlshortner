package rest

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kmchary/urlshortner/pkg/urlshortener"
	"net/http"
)

type Handler struct {
	urlService urlshortener.Service
	router     *chi.Mux
}

func NewHandler(us urlshortener.Service) *Handler {
	r := chi.NewRouter()

	// default middlewares provided by chi
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	h := &Handler{us, r}
	r.Post("/", h.GetShortUrl)

	return h
}

func (h *Handler) GetShortUrl(w http.ResponseWriter, r *http.Request) {
	var urlRequest urlshortener.Request
	var urlResponse urlshortener.Response
	decoder := json.NewDecoder(r.Body)
	encoder := json.NewEncoder(w)

	if err := decoder.Decode(&urlRequest); err != nil {
		urlResponse.Error = "unable to decode the request body"
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(urlResponse)
	}
	shortUrl, err := h.urlService.ShortenURL(urlRequest.Url, urlRequest.UserId)
	if err != nil {
		urlResponse.Error = "url service failed to generate short url"
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(urlResponse)
	}

	urlResponse.ShortUrl = shortUrl
	w.WriteHeader(http.StatusOK)
	encoder.Encode(urlResponse)
}
