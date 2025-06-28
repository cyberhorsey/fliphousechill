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
				Icon:      "",
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
	{
		Label:       "Mar A Lago Trump Visit Cost",
		RefreshRate: 5 * time.Minute,
		FetchFunc: func() (*MarketCapResp, error) {
			return &MarketCapResp{
				Label:     "Mar A Lago Trump Visit Cost",
				Icon:      "",
				MarketCap: 3400000,
			}, nil
		},
	},
	{
		Label:       "Pebble Time Kickstartre Amount Raised",
		RefreshRate: 5 * time.Minute,
		FetchFunc: func() (*MarketCapResp, error) {
			return &MarketCapResp{
				Label:     "Pebble Time Kickstarter Amount Raised",
				Icon:      "",
				MarketCap: 20338986,
			}, nil
		},
	},
	{
		Label:       "NGMI Pandas NFT Collection",
		RefreshRate: 5 * time.Minute,
		FetchFunc: func() (*MarketCapResp, error) {
			return &MarketCapResp{
				Label:     "NGMI Pandas NFT Collection",
				Icon:      "",
				MarketCap: 15100,
			}, nil
		},
	},
	{
		Label:       "Hoover Dam Construction Cost",
		RefreshRate: 5 * time.Minute,
		FetchFunc: func() (*MarketCapResp, error) {
			return &MarketCapResp{
				Label:     "Hoover Dam Construction Cost",
				Icon:      "",
				MarketCap: 49000000, // $49 M
			}, nil
		},
	},
	{
		Label:       "Cost to Build 1 Mile of Urban Road in US",
		RefreshRate: 5 * time.Minute,
		FetchFunc: func() (*MarketCapResp, error) {
			return &MarketCapResp{
				Label:     "Cost to Build 1 Mile of Urban Road in US",
				Icon:      "",
				MarketCap: 30000000, // $30 M
			}, nil
		},
	},
	{
		Label:       "Stephen Curry 2022 NBA Salary",
		RefreshRate: 5 * time.Minute,
		FetchFunc: func() (*MarketCapResp, error) {
			return &MarketCapResp{
				Label:     "Stephen Curry 2022 NBA Salary",
				Icon:      "",
				MarketCap: 48000000, // $48 M
			}, nil
		},
	},
	{
		Label:       "Boobs in California",
		RefreshRate: 5 * time.Minute,
		FetchFunc: func() (*MarketCapResp, error) {
			return &MarketCapResp{
				Label:     "Boobs in California",
				Icon:      "",
				MarketCap: 38000000, // $48 M
			}, nil
		},
	},
	{
		Label:       "District 9 Production Budget",
		RefreshRate: 5 * time.Minute,
		FetchFunc: func() (*MarketCapResp, error) {
			return &MarketCapResp{
				Label:     "District 9 Production Budget",
				Icon:      "",
				MarketCap: 30000000, // $30 M
			}, nil
		},
	},
	{
		Label:       "NYC Public Art Installation Grant",
		RefreshRate: 5 * time.Minute,
		FetchFunc: func() (*MarketCapResp, error) {
			return &MarketCapResp{
				Label:     "NYC Public Art Installation Grant",
				Icon:      "",
				MarketCap: 12000000, // $12 M
			}, nil
		},
	},
	{
		Label:       "Broadway Musical Production Cost",
		RefreshRate: 5 * time.Minute,
		FetchFunc: func() (*MarketCapResp, error) {
			return &MarketCapResp{
				Label:     "Broadway Musical Production Cost",
				Icon:      "",
				MarketCap: 17000000, // $20 M
			}, nil
		},
	},
	{
		Label:       "Fortnite World Cup 2019 Prize Pool",
		RefreshRate: 5 * time.Minute,
		FetchFunc: func() (*MarketCapResp, error) {
			return &MarketCapResp{
				Label:     "Fortnite World Cup 2019 Prize Pool",
				Icon:      "",
				MarketCap: 30000000, // $30 M
			}, nil
		},
	},
	{
		Label:       "US Small Police Dept Annual Budget",
		RefreshRate: 5 * time.Minute,
		FetchFunc: func() (*MarketCapResp, error) {
			return &MarketCapResp{
				Label:     "US Small Police Dept Annual Budget",
				Icon:      "",
				MarketCap: 25000000, // $25 M
			}, nil
		},
	},
	{
		Label:       "White House Lawn Renovation Cost",
		RefreshRate: 5 * time.Minute,
		FetchFunc: func() (*MarketCapResp, error) {
			return &MarketCapResp{
				Label:     "White House Lawn Renovation Cost",
				Icon:      "",
				MarketCap: 12000000, // $12 M
			}, nil
		},
	},
	{
		Label:       "Fleshlight yearly sales",
		RefreshRate: 5 * time.Minute,
		FetchFunc: func() (*MarketCapResp, error) {
			return &MarketCapResp{
				Label:     "Fleshlight yearly sales",
				Icon:      "",
				MarketCap: 50000000,
			}, nil
		},
	},
	{
		Label:       "Crazy Frog Yearly Ringstone Sales Peak",
		RefreshRate: 5 * time.Minute,
		FetchFunc: func() (*MarketCapResp, error) {
			return &MarketCapResp{
				Label:     "Crazy Frog Yearly Ringstone Sales Peak",
				Icon:      "",
				MarketCap: 50000000,
			}, nil
		},
	},
	{
		Label:       "Sale Price of CryptoPunk #7523",
		RefreshRate: 5 * time.Minute,
		FetchFunc: func() (*MarketCapResp, error) {
			return &MarketCapResp{
				Label:     "Sale Price of CryptoPunk #7523",
				Icon:      "",
				MarketCap: 11754000, // $11.754 M
			}, nil
		},
	},
	{
		Label:       "Cannes Film Festival 2023 Organizing Budget",
		RefreshRate: 5 * time.Minute,
		FetchFunc: func() (*MarketCapResp, error) {
			return &MarketCapResp{
				Label:     "Cannes Film Festival 2023 Organizing Budget",
				Icon:      "",
				MarketCap: 15100000, // $15 M
			}, nil
		},
	},
	{
		Label:       "Annual Operating Budget of San Diego Zoo",
		RefreshRate: 5 * time.Minute,
		FetchFunc: func() (*MarketCapResp, error) {
			return &MarketCapResp{
				Label:     "Annual Operating Budget of San Diego Zoo",
				Icon:      "",
				MarketCap: 40000000, // $40 M
			}, nil
		},
	},
	{
		Label:       "Red Bull Stratos Jump Project Cost",
		RefreshRate: 5 * time.Minute,
		FetchFunc: func() (*MarketCapResp, error) {
			return &MarketCapResp{
				Label:     "Red Bull Stratos Jump Project Cost",
				Icon:      "",
				MarketCap: 50000000, // $50 M
			}, nil
		},
	},
	{
		Label:       "McRib Annual Sales in US",
		RefreshRate: 5 * time.Minute,
		FetchFunc: func() (*MarketCapResp, error) {
			return &MarketCapResp{
				Label:     "McRib Annual Sales in US",
				Icon:      "",
				MarketCap: 18600000,
			}, nil
		},
	},
	{
		Label:       "Doritos Locos Tacos Launch Day Sales",
		RefreshRate: 5 * time.Minute,
		FetchFunc: func() (*MarketCapResp, error) {
			return &MarketCapResp{
				Label:     "Doritos Locos Tacos Launch Day Sales",
				Icon:      "",
				MarketCap: 15700000,
			}, nil
		},
	},
	{
		Label:       "Fyre Festival Actual Expenses",
		RefreshRate: 5 * time.Minute,
		FetchFunc: func() (*MarketCapResp, error) {
			return &MarketCapResp{
				Label:     "Fyre Festival Actual Expenses",
				Icon:      "",
				MarketCap: 24300000, // $25 M
			}, nil
		},
	},
	{
		Label:       "Unicorn Frappuccino Day-1 Sales",
		RefreshRate: 5 * time.Minute,
		FetchFunc: func() (*MarketCapResp, error) {
			return &MarketCapResp{
				Label:     "Unicorn Frappuccino Day-1 Sales",
				Icon:      "",
				MarketCap: 17000000, // $17 M
			}, nil
		},
	},
	{
		Label:       "Chernobyl Miniseries Budget",
		RefreshRate: 5 * time.Minute,
		FetchFunc: func() (*MarketCapResp, error) {
			return &MarketCapResp{
				Label:     "Chernobyl Miniseries Budget",
				Icon:      "",
				MarketCap: 18000000, // $18 M
			}, nil
		},
	},
	{
		Label:       "Budget for Moon Cheese Research",
		RefreshRate: 5 * time.Minute,
		FetchFunc: func() (*MarketCapResp, error) {
			return &MarketCapResp{
				Label:     "Budget for Moon Cheese Research",
				Icon:      "",
				MarketCap: 18000000, // $18 M
			}, nil
		},
	},
	{
		Label:       "Vancouver Birthday Clown Subsidies",
		RefreshRate: 5 * time.Minute,
		FetchFunc: func() (*MarketCapResp, error) {
			return &MarketCapResp{
				Label:     "Vancouver Birthday Clown Subsidies",
				Icon:      "",
				MarketCap: 17500000, // $17.5 M
			}, nil
		},
	},
	{
		Label:       "Unicorn Statue Construction in Vegas",
		RefreshRate: 5 * time.Minute,
		FetchFunc: func() (*MarketCapResp, error) {
			return &MarketCapResp{
				Label:     "Unicorn Statue Construction in Vegas",
				Icon:      "",
				MarketCap: 17800000, // $17.8 M
			}, nil
		},
	},
}
