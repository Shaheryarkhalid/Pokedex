package pokeapi

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
	cache      *cache
}

func NewClient(interval, timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: NewCache(interval),
	}

}
