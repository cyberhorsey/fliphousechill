package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/cyberhorsey/fliphousechill/cache"
	"github.com/cyberhorsey/fliphousechill/marketcaps"
)

func withCORS(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		handler(w, r)
	}
}

var defaultCache *cache.Cache

func updateCache(ctx context.Context, cache *cache.Cache) {
	now := time.Now().UTC()
	for _, mc := range marketcaps.MarketCaps {
		go func(mc marketcaps.MarketCapReq) {
			if item, found := cache.GetRaw(mc.Label); found {
				if now.Sub(item.CachedAt) < mc.RefreshRate {
					slog.Info("Using cached market cap",
						"label", mc.Label,
						"marketCap", item.Item.MarketCap,
					)
					return
				}
			}

			resp, err := mc.FetchFunc()
			if err != nil {
				slog.Error("Failed to fetch market cap",
					"label", mc.Label,
					"error", err,
				)
				return
			}

			cache.Set(mc.Label, *resp)
			slog.Info("Fetched new market cap",
				"label", resp.Label,
				"marketCap", resp.MarketCap,
			)
		}(mc)
	}
}

// in your worker
func cacheLoop(ctx context.Context, cache *cache.Cache) {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			updateCache(ctx, cache)
		}
	}
}

func main() {
	defaultCache = cache.New()
	updateCache(context.Background(), defaultCache)

	go cacheLoop(context.Background(), defaultCache)

	port := os.Getenv("PORT")
	if port == "" {
		slog.Info("PORT environment variable not set, defaulting to 8080")
		port = "8080"
	}
	http.HandleFunc("/cache", withCORS(func(w http.ResponseWriter, r *http.Request) {
		// send cache as JSON
		w.Header().Set("Content-Type", "application/json")
		items := defaultCache.All()
		if err := json.NewEncoder(w).Encode(items); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}))

	// here we wrap the optimism op-node interface expectation of a "/healthz" endpoint
	// and make it call the "/upcheck" endpoint of web3signer, proxying the request.
	http.HandleFunc("/healthz", withCORS(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	}))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
