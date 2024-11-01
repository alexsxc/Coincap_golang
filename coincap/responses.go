package coincap

import "fmt"

type assetsResponse struct {
	Asset     []Asset `json:"data"`
	Timestamp int64   `json:"timestamp"`
}

type assetResponse struct {
	Asset     Asset `json:"data"`
	Timestamp int64 `json:"timestamp"`
}

type Asset struct {
	ID               string `json:"id"`
	Rank             string `json:"rank"`
	Symbol           string `json:"symbol"`
	Name             string `json:"name"`
	Supply           string `json:"supply"`
	MaxSupply        string `json:"maxSupply"`
	MarketCapUSD     string `json:"marketCapUsd"`
	VolumeUSD24h     string `json:"volumeUsd24Hr"`
	PriceUSD         string `json:"priceUsd"`
	ChangePercent24h string `json:"changePercent24Hr"`
	VWrap24h         string `json:"vwap24Hr"`
	Explorer         string `json:"explorer"`
}

func (d Asset) Info() string {
	return fmt.Sprintf("[ID] %s | [RANK] %s | [SYMBOL] %s | [NAME] %s | [PRICE] %s | [EXPLORER] %s", d.ID, d.Rank, d.Symbol, d.Name, d.PriceUSD, d.Explorer)
}
