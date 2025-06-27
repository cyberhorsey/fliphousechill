package marketcaps

import (
	"time"

	"github.com/cyberhorsey/fliphousechill/coingecko"
)

type MarketCapResp struct {
	Label     string  `json:"label"`
	Icon      string  `json:"icon"`
	MarketCap float64 `json:"market_cap"`
}

type MarketCapReq struct {
	Label       string
	RefreshRate time.Duration
	FetchFunc   func() (*MarketCapResp, error)
}

var MarketCaps = []MarketCapReq{
	{
		Label:       "Ethereum Classic",
		RefreshRate: 5 * time.Minute,
		FetchFunc: func() (*MarketCapResp, error) {
			mc, err := coingecko.Fetch("ethereum-classic")
			if err != nil {
				return nil, err
			}

			return &MarketCapResp{
				Label:     "Ethereum Classic",
				Icon:      "https://assets.coingecko.com/coins/images/453/large/ethereum-classic-logo.png",
				MarketCap: mc,
			}, nil
		},
	},
	{
		Label:       "Chillhouse",
		RefreshRate: 5 * time.Minute,
		FetchFunc: func() (*MarketCapResp, error) {
			mc, err := coingecko.Fetch("chill-house")
			if err != nil {
				return nil, err
			}

			return &MarketCapResp{
				Label:     "Chillhouse",
				Icon:      "https://assets.coingecko.com/coins/images/453/large/ethereum-classic-logo.png",
				MarketCap: mc,
			}, nil
		},
	},
	{
		Label:       "The Emoji Movie Opening Weekend Box Office",
		RefreshRate: 5 * time.Minute,
		FetchFunc: func() (*MarketCapResp, error) {
			return &MarketCapResp{
				Label:     "The Emoji Movie Opening Weekend Box Office",
				Icon:      "",
				MarketCap: 24531923,
			}, nil
		},
	},
}
