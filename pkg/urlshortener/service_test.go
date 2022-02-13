package urlshortener

import (
	"github.com/kmchary/urlshortner/pkg/urlshortener/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestService_ShortenURL(t *testing.T) {
	mockRedisRepo := &mocks.Storage{}
	urlShortener := NewService(mockRedisRepo)

	tests := []struct {
		name             string
		url              string
		userId           string
		expectedShortUrl string
		expectedError    error
	}{
		{
			"Test1",
			"http://something.com/something/something",
			"testuser",
			"testurl",
			nil,
		},

		{
			"Test2",
			"http://something.com/something/something",
			"testuser",
			"testurl",
			FailedToStoreError,
		},
	}

	for _, test := range tests {
		if test.name == "Test2" {
			mockRedisRepo.On("Get", mock.Anything, mock.Anything).Return("")
			mockRedisRepo.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(FailedToStoreError)
		} else {
			mockRedisRepo.On("Get", mock.Anything, mock.Anything).Return("testurl")
			mockRedisRepo.On("Set", mock.Anything, mock.Anything, mock.Anything).Return(nil)
		}
		t.Run(test.name, func(t *testing.T) {
			actualShorturl, err := urlShortener.ShortenURL(test.url, test.userId)
			assert.Nil(t, err)
			assert.Equal(t, test.expectedShortUrl, actualShorturl)
		})
	}
}
