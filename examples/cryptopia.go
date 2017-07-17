package main

import (
	"fmt"

	"github.com/jyap808/go-cryptopia"
)

const (
	API_KEY    = ""
	API_SECRET = ""
)

func main() {
	// Cryptopia client
	cryptopia := cryptopia.New(API_KEY, API_SECRET)

	// GetTradePairs
	pairs, _ := cryptopia.GetTradePairs()
	fmt.Println(len(pairs))

	for i, _ := range pairs {
		if pairs[i].BaseSymbol == "BTC" {
			fmt.Println(pairs[i].Label)
		}
	}

}
