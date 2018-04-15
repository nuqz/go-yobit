package yobit

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	api = baseAPI("https://yobit.net/api/3/")
)

type baseAPI string

func (a baseAPI) method(name string, dest interface{}) error {
	resp, err := http.Get(fmt.Sprintf("%s%s", a, name))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(dest)
}

type PublicAPI struct {
	baseAPI
}

type Pair struct {
	DecimalPlaces uint    `json:"decimal_places"`
	MinPrice      float64 `json:"min_price"`
	MaxPrice      float64 `json:"max_price"`
	MinAmount     float64 `json:"min_amount"`
	Hidden        uint    `json:"hidden"`
	Fee           float64 `json:"fee"`
}

type InfoResponse struct {
	ServerTime int64            `json:"server_time"`
	Pairs      map[string]*Pair `json:"pairs"`
}

func (a *PublicAPI) Info() (*InfoResponse, error) {
	resp := new(InfoResponse)
	if err := a.method("info", resp); err != nil {
		return nil, err
	}

	return resp, nil
}

type TickerResponse struct {
	High    float64 `json:"high"`
	Low     float64 `json:"low"`
	Avg     float64 `json:"avg"`
	Vol     float64 `json:"vol"`
	VolCur  float64 `json:"vol_cur"`
	Last    float64 `json:"last"`
	Buy     float64 `json:"buy"`
	Sell    float64 `json:"sell"`
	Updated int     `json:"updated"`
}

func (a *PublicAPI) Ticker(name string) (*TickerResponse, error) {
	resp := map[string]*TickerResponse{name: new(TickerResponse)}
	if err := a.method(fmt.Sprintf("ticker/%s", name), &resp); err != nil {
		return nil, err
	}

	return resp[name], nil
}

type DepthResponse struct {
	Asks [][]float64 `json:"asks"`
	Bids [][]float64 `json:"bids"`
}

func (a *PublicAPI) Depth(name string) (*DepthResponse, error) {
	resp := map[string]*DepthResponse{name: new(DepthResponse)}
	if err := a.method(fmt.Sprintf("ticker/%s", name), &resp); err != nil {
		return nil, err
	}

	return resp[name], nil
}

type TradesResponse struct {
	Type      string  `json:"type"`
	Price     float64 `json:"price"`
	Amount    float64 `json:"amount"`
	Tid       int     `json:"tid"`
	Timestamp int     `json:"timestamp"`
}

func (a *PublicAPI) Trades(name string) (*TradesResponse, error) {
	resp := map[string]*TradesResponse{name: new(TradesResponse)}
	if err := a.method(fmt.Sprintf("ticker/%s", name), &resp); err != nil {
		return nil, err
	}

	return resp[name], nil
}
