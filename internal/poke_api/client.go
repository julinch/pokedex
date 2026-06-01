package pokeapi

import (
	"net/http"
	"time"

	"github.com/julinch/pokedex/internal/pokecache"
)

type Client struct {
	HttpClient http.Client
	Cache      pokecache.Cache
}

func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		Cache: pokecache.NewCache(cacheInterval),
		HttpClient: http.Client{
			Timeout: timeout,
		},
	}
}
