package yobit

import (
	"encoding/json"
	"os"
	"testing"
)

func Test_PublicAPI(t *testing.T) {
	api := &PublicAPI{api}

	info, err := api.Info()
	if err != nil {
		t.Error(err)
	}

	depth, err := api.Depth("eth_rur")
	if err != nil {
		t.Error(err)
	}

	trades, err := api.Trades("eth_rur")
	if err != nil {
		t.Error(err)
	}

	ticker, err := api.Ticker("eth_rur")
	if err != nil {
		t.Error(err)
	}

	_ = info
	_ = depth
	_ = trades
	_ = ticker

	json.NewEncoder(os.Stdout).Encode(ticker)
}
