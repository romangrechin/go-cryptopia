package cryptopia

const (
	CancelTradeTypeAll       = "All"
	CancelTradeTypeTrade     = "Trade"
	CancelTradeTypeTradePair = "TradePair"
)

type BalanceParams struct {
	Currency   string `json:",omitempty"`
	CurrencyId int    `json:",omitempty"`
}

type TradeHistoryParams struct {
	Market      string `json:",omitempty"`
	TradePairId int    `json:",omitempty"`
	Count       uint32 `json:",omitempty"`
}

type TransactionsParams struct {
	Type  string `json:",omitempty"`
	Count uint32 `json:",omitempty"`
}

type SubmitTradeParams struct {
	TradePairId int     `json:",omitempty"`
	Market      string  `json:",omitempty"`
	Type        string  `json:",omitempty"`
	Rate        float64 `json:",omitempty"`
	Amount      float64 `json:",omitempty"`
}

type CancelTradeParams struct {
	Type        string `json:",omitempty"`
	OrderId     int    `json:",omitempty"`
	TradePairId int    `json:",omitempty"`
}
