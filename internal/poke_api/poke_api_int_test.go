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
	cacheKey := pokedexEP + "canalave-city-area"
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
