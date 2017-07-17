package cryptopia

type Pair struct {
	Id           int    `json:"Id"`
	Label        string `json:"Label"`
	Currency     string `json:"Currency"`
	Symbol       string `json:"Symbol"`
	BaseCurrency string `json:"BaseCurrency"`
	BaseSymbol   string `json:"BaseSymbol"`
	Status       string `json:"Status"`
}

// {"Id":105,"Label":"FTC/BTC","Currency":"Feathercoin","Symbol":"FTC","BaseCurrency":"Bitcoin","BaseSymbol":"BTC","Status":"OK","StatusMessage":null,"TradeFee":0.20000000,"MinimumTrade":0.00000001,"MaximumTrade":100000000.0,"MinimumBaseTrade":0.00000500,"MaximumBaseTrade":100000000.0,"MinimumPrice":0.00000001,"MaximumPrice":100000000.0},
