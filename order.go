package cryptopia

const (
	OrderTypeBuy  = "Buy"
	OrderTypeSell = "Sell"
)

type Order struct {
	OrderId     int
	TradePairId int
	Market      string
	Type        string
	Rate        float64
	Amount      float64
	Total       float64
	Remaining   float64
	TimeStamp   Timestamp
}

type SubmitOrder struct {
	OrderId      int
	FilledOrders []int
}
