package pokeapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	pokecache "github.com/julinch/pokedex/internal/pokecache"
)

func newCache() *pokecache.Cache {
	cache := pokecache.NewCache(5 * time.Minute)
	return &cache
}

// ---- GetLocationArea tests ----

func TestGetLocationArea_Success(t *testing.T) {
	expected := LocationArea{Name: "canalave-city-area"}
	body, _ := json.Marshal(expected)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}))
	defer srv.Close()

	srvCli := srv.Client()
	client := &Client{HttpClient: *srvCli}
	// point pokedexEP at the test server by passing the full URL as areaName
	area, err := client.GetLocationArea(expected.Name, newCache())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if area.Name != expected.Name {
		t.Errorf("expected name %q, got %q", expected.Name, area.Name)
	}
}

func TestGetLocationArea_EmptyName(t *testing.T) {
	client := &Client{HttpClient: http.Client{}}
	_, err := client.GetLocationArea("", newCache())
	if err == nil {
		t.Fatal("expected error for empty areaName, got nil")
	}
}

func TestGetLocationArea_CacheHit(t *testing.T) {
	expected := LocationArea{Name: "canalave-city-area"}
	body, _ := json.Marshal(expected)

	cache := newCache()
	cacheKey := pokedexLocationAreaEP + "canalave-city-area"
	cache.Add(cacheKey, body)

	// no real server needed — cache should be hit before any HTTP call
	client := &Client{HttpClient: http.Client{}, Cache: *cache}
	area, err := client.GetLocationArea("canalave-city-area", cache)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if area.Name != expected.Name {
		t.Errorf("expected name %q, got %q", expected.Name, area.Name)
	}
}

// ---- GetPage tests ----

func TestGetPage_Success(t *testing.T) {
	expected := Page{Next: "https://pokeapi.co/api/v2/location-area/?offset=20"}
	body, _ := json.Marshal(expected)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}))
	defer srv.Close()

	page, err := GetPage(srv.URL, newCache())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if page.Next != expected.Next {
		t.Errorf("expected Next %q, got %q", expected.Next, page.Next)
	}
}

func TestGetPage_CacheHit(t *testing.T) {
	expected := Page{Next: "https://pokeapi.co/api/v2/location-area/?offset=40"}
	body, _ := json.Marshal(expected)

	cache := newCache()
	cacheKey := "cached-page-url"
	cache.Add(cacheKey, body)

	page, err := GetPage(cacheKey, cache)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if page.Next != expected.Next {
		t.Errorf("expected Next %q, got %q", expected.Next, page.Next)
	}
}

func TestClient_GetPokemon(t *testing.T) {
	// sample fake pokemon JSON (must match your struct)
	mockPokemonJSON := `{
		"name": "pikachu",
		"id": 25
	}`

	cache := (pokecache.NewCache(time.Second))

	tests := []struct {
		name          string
		timeout       time.Duration
		cacheInterval time.Duration
		pokemonName   string
		cache         *pokecache.Cache
		want          Pokemon
		wantErr       bool

		// internal test server behavior
		serverHandler http.HandlerFunc
	}{
		{
			name:          "successful fetch from API",
			timeout:       time.Second,
			cacheInterval: time.Second,
			pokemonName:   "pikachu",
			cache:         &cache,

			want: Pokemon{
				Name: "pikachu",
				ID:   25,
			},
			wantErr: false,

			serverHandler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(mockPokemonJSON))
			},
		},
		{
			name:          "pokemon not found (404)",
			timeout:       time.Second,
			cacheInterval: time.Second,
			pokemonName:   "unknownmon",
			cache:         &cache,

			want:    Pokemon{},
			wantErr: true,

			serverHandler: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNotFound)
			},
		},
		{
			name:          "empty pokemon name returns error",
			timeout:       time.Second,
			cacheInterval: time.Second,
			pokemonName:   "",
			cache:         &cache,

			want:    Pokemon{},
			wantErr: true,

			serverHandler: nil, // won't be used
		},
		{
			name:          "cache hit returns cached data (no server call)",
			timeout:       time.Second,
			cacheInterval: time.Second,
			pokemonName:   "pikachu",
			cache:         &cache,

			want: Pokemon{
				Name: "pikachu",
				ID:   25,
			},
			wantErr: false,

			serverHandler: func(w http.ResponseWriter, r *http.Request) {
				t.Fatal("server should not be called on cache hit")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// var serverURL string
			if tt.serverHandler != nil {
				server := httptest.NewServer(tt.serverHandler)
				defer server.Close()

				// override base URL used by client (IMPORTANT)
				// serverURL = server.URL + "/"
			}

			client := NewClient(tt.timeout, tt.cacheInterval)

			// ⚠️ assuming you can override base URL like this:
			// client.BaseURL = serverURL

			// preload cache for cache test
			if tt.name == "cache hit returns cached data (no server call)" {
				data, _ := json.Marshal(tt.want)
				tt.cache.Add(pokedexPokemonEP+"pikachu", data)
			}

			got, err := client.GetPokemon(tt.pokemonName, tt.cache)

			if err != nil {
				if !tt.wantErr {
					t.Errorf("unexpected error: %v", err)
				}
				return
			}

			if tt.wantErr {
				t.Fatal("expected error but got nil")
			}

			if got.Name != tt.want.Name || got.ID != tt.want.ID {
				t.Errorf("got %+v, want %+v", got, tt.want)
			}
		})
	}
}
