package models

type Trade struct {
	ID      string  `json:"id"`
	UserID  string  `json:"user_id"`
	Type    string  `json:"type"` // buy or sell
	AssetID string  `json:"crypto"`
	Price   float64 `json:"price"`
	Amount  float64 `json:"amount"`
	Status  string  `json:"status"` // open, locked, completed
}
