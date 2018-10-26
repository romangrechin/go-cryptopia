// Package Cryptopia is an implementation of the Cryptopia API in Golang.
package cryptopia

import (
	"encoding/json"
	"errors"
)

const (
	API_BASE                   = "https://www.cryptopia.co.nz/api/" // Cryptopia API endpoint
	API_VERSION                = "v1.1"                             // Cryptopia API version
	DEFAULT_HTTPCLIENT_TIMEOUT = 30                                 // HTTP client timeout
)

// New return a instanciate cryptopia struct
func New(apiKey, apiSecret string) *Cryptopia {
	client := NewClient(apiKey, apiSecret)
	return &Cryptopia{client}
}

// handleErr gets JSON response from Cryptopia API en deal with error
func handleErr(r jsonResponse) error {
	if !r.Success {
		return errors.New(r.Message)
	}
	return nil
}

// cryptopia represent a cryptopia client
type Cryptopia struct {
	client *client
}

// set enable/disable http request/response dump
func (c *Cryptopia) SetDebug(enable bool) {
	c.client.debug = enable
}

// GetTradePairs Returns all trade pair data
func (b *Cryptopia) GetTradePairs() (pairs []Pair, err error) {
	r, err := b.client.do("GET", "GetTradePairs", "", false)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(response.Result, &pairs)
	return
}

// GetBalance Returns all balances or a specific currency balance
func (b *Cryptopia) GetBalance(params BalanceParams) (balances []Balance, err error) {
	payload, err := json.Marshal(params)
	if err != nil {
		return
	}
	r, err := b.client.do("POST", "GetBalance", string(payload), true)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(response.Result, &balances)
	return
}

// GetTradeHistory Returns all trade history data
func (b *Cryptopia) GetTradeHistory(params TradeHistoryParams) (trades []Trade, err error) {
	payload, err := json.Marshal(params)
	if err != nil {
		return
	}
	r, err := b.client.do("POST", "GetTradeHistory", string(payload), true)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(response.Result, &trades)
	return
}

// GetTransactions Returns a list of transactions
func (b *Cryptopia) GetTransactions(params TransactionsParams) (transactions []Transaction, err error) {
	payload, err := json.Marshal(params)
	if err != nil {
		return
	}
	r, err := b.client.do("POST", "GetTransactions", string(payload), true)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(response.Result, &transactions)
	return
}

// GetMarkets Returns a list of markets
func (b *Cryptopia) GetMarkets() (markets []Market, err error) {
	r, err := b.client.do("GET", "GetMarkets", "", false)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(response.Result, &markets)
	return
}

// GetDepositAddress Returns specific currency deposit address
func (b *Cryptopia) GetDepositAddress(params BalanceParams) (address DepositAddress, err error) {
	payload, err := json.Marshal(params)
	if err != nil {
		return
	}
	r, err := b.client.do("POST", "GetDepositAddress", string(payload), true)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}

	err = json.Unmarshal(response.Result, &address)
	return
}

// GetOpenOrders Returns a list of open orders
func (b *Cryptopia) GetOpenOrders(params TradeHistoryParams) (orders []Order, err error) {
	payload, err := json.Marshal(params)
	if err != nil {
		return
	}
	r, err := b.client.do("POST", "GetOpenOrders", string(payload), true)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}

	err = json.Unmarshal(response.Result, &orders)
	return
}

// SubmitTrade
func (b *Cryptopia) SubmitTrade(params SubmitTradeParams) (orders []SubmitOrder, err error) {
	payload, err := json.Marshal(params)
	if err != nil {
		return
	}
	r, err := b.client.do("POST", "SubmitTrade", string(payload), true)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}

	err = json.Unmarshal(response.Result, &orders)
	return
}

// CancelTrade
func (b *Cryptopia) CancelTrade(params CancelTradeParams) (orders []SubmitOrder, err error) {
	payload, err := json.Marshal(params)
	if err != nil {
		return
	}
	r, err := b.client.do("POST", "CancelTrade", string(payload), true)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}

	err = json.Unmarshal(response.Result, &orders)
	return
}
