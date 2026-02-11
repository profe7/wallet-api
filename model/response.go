package model

//The skeleton of a standard API response
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

//The response for balance inquiry
type BalanceResponse struct {
	UserID  string `json:"user_id"`
	Balance int64  `json:"balance"`
}

//The response for a withdrawal operation
type WithdrawResponse struct {
	UserID     string `json:"user_id"`
	Withdrawn  int64  `json:"withdrawn"`
	NewBalance int64  `json:"new_balance"`
}
