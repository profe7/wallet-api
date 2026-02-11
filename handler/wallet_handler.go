package handler

import (
	"encoding/json"
	"net/http"

	"wallet-api/middleware/httperrors"
	"wallet-api/model"
	"wallet-api/service"
	"wallet-api/utils"
)

// Future reference notes, this is like a controller in Spring
type WalletHandler struct {
	service *service.WalletService
}

// Returns a new WalletHandler, wired to the given service
func NewWalletHandler(svc *service.WalletService) *WalletHandler {
	return &WalletHandler{service: svc}
}

// Check balance, accepts get with query param user_id, returns a JSON response
func (h *WalletHandler) CheckBalance(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return httperrors.MethodNotAllowedError("Check Balance Handler", "only GET is allowed")
	}

	userID := r.URL.Query().Get("user_id")
	if userID == "" {
		return httperrors.BadRequestError("Check Balance Handler", "query param user_id is missing")
	}

	result, err := h.service.CheckBalance(userID)
	if err != nil {
		return err
	}

	utils.WriteJSON(w, http.StatusOK, model.APIResponse{
		Success: true,
		Data:    result,
	})
	return nil
}

// Withdraw, accepts post with JSON body, returns a JSON response
func (h *WalletHandler) Withdraw(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodPost {
		return httperrors.MethodNotAllowedError("Withdraw Handler", "only POST is allowed")
	}

	var req model.WithdrawRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return httperrors.BadRequestError("Withdraw Handler", "invalid JSON body: "+err.Error())
	}

	result, err := h.service.Withdraw(req)
	if err != nil {
		return err
	}

	utils.WriteJSON(w, http.StatusOK, model.APIResponse{
		Success: true,
		Message: "Withdrawal successful",
		Data:    result,
	})
	return nil
}
