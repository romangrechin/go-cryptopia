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
