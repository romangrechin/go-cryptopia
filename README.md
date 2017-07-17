go-cryptopia
==========

go-cryptopia is an implementation of the Cryptopia API (public and private) in Golang.

Based off of https://github.com/toorop/go-bittrex/

This library is more of a framework for some bots I use so it is expected that a lot of things don't work but pull requests are excepted.

## Import
	import "github.com/jyap808/go-cryptopia"
	
## Usage
~~~ go
package main

import (
    "fmt"
    "github.com/jyap808/go-cryptopia"
)

const (
    API_KEY    = "YOUR_API_KEY"
    API_SECRET = "YOUR_API_SECRET"
)

func main() {
    // Cryptopia client
    cryptopia := cryptopia.New(API_KEY, API_SECRET)

    // Get tickers
    tickers, err := cryptopia.GetTickers()
    fmt.Println(err, tickers)
}
~~~	

See ["Examples" folder for more... examples](https://github.com/jyap808/go-cryptopia/blob/master/examples/cryptopia.go)

## Stay tuned
[Follow me on Twitter](https://twitter.com/jyap)

