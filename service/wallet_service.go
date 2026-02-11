package service

import (
	"wallet-api/middleware/httperrors"
	"wallet-api/model"
	"wallet-api/repository"
)

// WalletService contains the business logic for wallet operations.
// Self Note : Analogous to a @Service class in Spring Boot.
type WalletService struct {
	repo *repository.AccountRepository
}

// NewWalletService creates a service wired to the given repository.
func NewWalletService(repo *repository.AccountRepository) *WalletService {
	return &WalletService{repo: repo}
}

// CheckBalance returns the current balance for a user.
func (s *WalletService) CheckBalance(userID string) (*model.BalanceResponse, error) {
	if userID == "" {
		return nil, httperrors.BadRequestError("WalletService", "user_id is required")
	}

	account, err := s.repo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	return &model.BalanceResponse{
		UserID:  account.UserID,
		Balance: account.Balance,
	}, nil
}

// Withdraw processes a withdrawal and returns the result.
func (s *WalletService) Withdraw(req model.WithdrawRequest) (*model.WithdrawResponse, error) {
	if req.UserID == "" {
		return nil, httperrors.BadRequestError("WalletService", "user_id is required")
	}
	if req.Amount <= 0 {
		return nil, httperrors.BadRequestError("WalletService", "amount must be greater than zero")
	}

	account, err := s.repo.Withdraw(req.UserID, req.Amount)
	if err != nil {
		return nil, err
	}

	return &model.WithdrawResponse{
		UserID:     account.UserID,
		Withdrawn:  req.Amount,
		NewBalance: account.Balance,
	}, nil
}
