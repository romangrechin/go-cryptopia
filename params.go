package cryptopia

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
