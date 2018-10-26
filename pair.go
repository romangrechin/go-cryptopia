package cryptopia

type Pair struct {
	Id               int64   `json:"Id"`
	Label            string  `json:"Label"`
	Currency         string  `json:"Currency"`
	Symbol           string  `json:"Symbol"`
	BaseCurrency     string  `json:"BaseCurrency"`
	BaseSymbol       string  `json:"BaseSymbol"`
	Status           string  `json:"Status"`
	StatusMessage    string  `json:"StatusMessage"`
	TradeFee         float64 `json:"TradeFee"`
	MinimumTrade     float64 `json:"MinimumTrade"`
	MaximumTrade     float64 `json:"MaximumTrade"`
	MinimumBaseTrade float64 `json:"MinimumBaseTrade"`
	MaximumBaseTrade float64 `json:"MaximumBaseTrade"`
	MinimumPrice     float64 `json:"MinimumPrice"`
	MaximumPrice     float64 `json:"MaximumPrice"`
}

// {"Id":105,"Label":"FTC/BTC","Currency":"Feathercoin","Symbol":"FTC","BaseCurrency":"Bitcoin","BaseSymbol":"BTC","Status":"OK","StatusMessage":null,"TradeFee":0.20000000,"MinimumTrade":0.00000001,"MaximumTrade":100000000.0,"MinimumBaseTrade":0.00000500,"MaximumBaseTrade":100000000.0,"MinimumPrice":0.00000001,"MaximumPrice":100000000.0},
