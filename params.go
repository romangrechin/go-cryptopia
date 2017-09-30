package cryptopia

type TradeHistoryParams struct {
	Market      string `json:",omitempty"`
	TradePairId int    `json:",omitempty"`
	Count       uint32 `json:",omitempty"`
}
