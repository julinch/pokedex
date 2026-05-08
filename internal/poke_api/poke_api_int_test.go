package pokeapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	pokecache "pokedex/internal/pokecache"
)

func newCache() *pokecache.Cache {
	return pokecache.NewCache(5 * time.Minute)
}

func TestGetPage_FetchesFromNetwork(t *testing.T) {
	expected := Page{
		Next:     *strPtr("https://pokeapi.co/api/v2/location-area/?offset=20"),
		Previous: nil,
		Results:  []Result{{Name: "canalave-city-area"}},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(expected)
	}))
	defer server.Close()

	cache := newCache()
	page, err := GetPage(server.URL, cache)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(page.Results) != 1 || page.Results[0].Name != "canalave-city-area" {
		t.Errorf("unexpected results: %+v", page.Results)
	}
}

func TestGetPage_UsesCache(t *testing.T) {
	callCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++
		json.NewEncoder(w).Encode(Page{
			Results: []Result{{Name: "cached-area"}},
		})
	}))
	defer server.Close()

	cache := newCache()

	// First call
	_, err := GetPage(server.URL, cache)
	if err != nil {
		t.Fatalf("first call failed: %v", err)
	}

	// Second call
	page, err := GetPage(server.URL, cache)
	if err != nil {
		t.Fatalf("second call failed: %v", err)
	}
	if callCount != 1 {
		t.Errorf("expected 1 HTTP call, got %d", callCount)
	}
	if page.Results[0].Name != "cached-area" {
		t.Errorf("unexpected cached result: %+v", page.Results)
	}
}

func TestGetPage_HandlesInvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("not-valid-json"))
	}))
	defer server.Close()

	cache := newCache()
	_, err := GetPage(server.URL, cache)
	if err == nil {
		t.Error("expected an error for invalid JSON, got nil")
	}
}

func strPtr(s string) *string { return &s }
