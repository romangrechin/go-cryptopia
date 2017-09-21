package cryptopia

type Trade struct {
	TradeID     int
	TradePairID int
	Market      string
	Type        string
	Rate        float64
	Amount      float64
	Total       float64
	Fee         float64
	TimeStamp   Timestamp
}
