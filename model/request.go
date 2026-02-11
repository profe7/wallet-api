package model

//The skeleton of a withdraw request
type WithdrawRequest struct {
	UserID string `json:"user_id"`
	Amount int64  `json:"amount"`
}

//The skeleton of a balance request
type BalanceRequest struct {
	UserID string `json:"user_id"`
}
