package cryptopia

type Transaction struct {
	Id            int
	Currency      string
	TxId          string
	Type          string
	Amount        float64
	Fee           float64
	Status        string
	Confirmations int
	TimeStamp     Timestamp
	Address       string
}
