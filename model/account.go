package model

//Account struct, the "class" representing a user account
type Account struct {
	UserID  string `json:"user_id"`
	Balance int64  `json:"balance"`
}
