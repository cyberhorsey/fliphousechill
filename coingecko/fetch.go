package coingecko

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func Fetch(coin string) (float64, error) {
	url := "https://api.coingecko.com/api/v3/simple/price?ids=" + coin + "&vs_currencies=usd&include_market_cap=true"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch data from CoinGecko: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error response from CoinGecko: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	var result map[string]map[string]float64
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("Failed to parse JSON response: %v", err)
	}

	mc, ok := result[coin]["usd_market_cap"]
	if !ok {
		log.Fatalf("Mc for coin %s not found in response", coin)
	}

	fmt.Printf("Current marketcap of %s in USD: %.2f\n", coin, mc)

	return mc, nil
}
