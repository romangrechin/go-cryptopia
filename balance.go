package cryptopia

type Balance struct {
	CurrencyId      int
	Symbol          string
	Total           float64
	Available       float64
	Unconfirmed     float64
	HeldForTrades   float64
	PendingWithdraw float64
	Address         string
	Status          string
	StatusMessage   string
	BaseAddress     string
}
