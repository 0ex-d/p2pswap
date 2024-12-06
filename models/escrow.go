package models

type Escrow struct {
	TradeID string  `json:"trade_id"`
	Amount  float64 `json:"amount"`
	Crypto  string  `json:"crypto"`
	Locked  bool    `json:"locked"`
}
