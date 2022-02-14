package rest

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kmchary/urlshortner/pkg/urlshortener"
	"net/http"
)

type Handler struct {
	UrlService urlshortener.Service
	Router     *chi.Mux
}

func NewHandler(us urlshortener.Service) *Handler {
	return &Handler{us, chi.NewRouter()}
}

func (h *Handler) InitRoutes() {

	// default middlewares provided by chi
	h.Router.Use(middleware.RequestID)
	h.Router.Use(middleware.RealIP)
	h.Router.Use(middleware.Logger)
	h.Router.Use(middleware.Recoverer)
	h.Router.Post("/api/base62random", h.GenerateShortUrlUsingBase62RandomChars)
	h.Router.Get("/{short_url}", h.GetShortUrl)
}

func (h *Handler) GenerateShortUrlUsingBase62RandomChars(w http.ResponseWriter, r *http.Request) {
	var urlRequest urlshortener.Request
	var urlResponse urlshortener.Response
	decoder := json.NewDecoder(r.Body)
	encoder := json.NewEncoder(w)

	if err := decoder.Decode(&urlRequest); err != nil {
		urlResponse.Error = "unable to decode the request body"
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(urlResponse)
	}
	shortUrl, err := h.UrlService.GenerateURLUsingBase62RandomChars(urlRequest.Url, urlRequest.UserId)
	if err != nil {
		urlResponse.Error = "url service failed to generate short url"
		w.WriteHeader(http.StatusInternalServerError)
		encoder.Encode(urlResponse)
	}

	urlResponse.ShortUrl = shortUrl
	w.WriteHeader(http.StatusOK)
	encoder.Encode(urlResponse)
}

func (h *Handler) GetShortUrl(w http.ResponseWriter, r *http.Request) {
	shortUrl := chi.URLParam(r, "short_url")
	var urlResponse urlshortener.Response
	encoder := json.NewEncoder(w)

	if shortUrl == "" {
		urlResponse.Error = "Invalid request."
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(urlResponse)
	}
	url := h.UrlService.GetActualUrl(shortUrl)
	if url == "" {
		urlResponse.Error = "Invalid Short URL code."
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(urlResponse)
	}

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
