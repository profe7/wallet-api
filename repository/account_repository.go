package repository

import (
	"database/sql"
	"errors"

	"wallet-api/middleware/httperrors"
	"wallet-api/model"
)

// Common errors that can be returned by the repository methods
var (
	ErrAccountNotFound   = httperrors.NotFoundError("Account Repository", "account not found")
	ErrInsufficientFunds = httperrors.BadRequestError("Account Repository", "insufficient funds")
)

// This is just like the repository layer in Spring, handling direct DB operations
type AccountRepository struct {
	db *sql.DB
}

// Returns a new AccountRepository, wired to the given database
func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

// Queries the DB for an account by user ID, returns the account or an error
func (r *AccountRepository) GetByUserID(userID string) (*model.Account, error) {
	row := r.db.QueryRow("SELECT user_id, balance FROM accounts WHERE user_id = ?", userID)

	var account model.Account
	err := row.Scan(&account.UserID, &account.Balance)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrAccountNotFound
		}
		return nil, httperrors.InternalServerError("Account Repository", err.Error())
	}
	return &account, nil
}

// Withdraw decreases the balance of the account by the specified amount.
// Returns the updated account or an error.
func (r *AccountRepository) Withdraw(userID string, amount int64) (*model.Account, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, httperrors.InternalServerError("Account Repository", err.Error())
	}
	//Rollback the transaction if not committed
	defer tx.Rollback()

	//This mimics the behaviour of @Transactional in Spring, guarding atomicity
	var balance int64
	err = tx.QueryRow("SELECT balance FROM accounts WHERE user_id = ?", userID).Scan(&balance)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrAccountNotFound
		}
		return nil, httperrors.InternalServerError("Account Repository", err.Error())
	}

	if balance < amount {
		return nil, ErrInsufficientFunds
	}

	newBalance := balance - amount
	_, err = tx.Exec("UPDATE accounts SET balance = ? WHERE user_id = ?", newBalance, userID)
	if err != nil {
		return nil, httperrors.InternalServerError("Account Repository", err.Error())
	}

	if err = tx.Commit(); err != nil {
		return nil, httperrors.InternalServerError("AccountRepository", err.Error())
	}

	return &model.Account{UserID: userID, Balance: newBalance}, nil
}
