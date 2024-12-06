package models

type User struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	KYCVerified  bool   `json:"kyc_verified"`
}
