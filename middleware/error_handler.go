package middleware

import (
	"log"
	"net/http"

	"wallet-api/middleware/httperrors"
	"wallet-api/model"
	"wallet-api/utils"
)

type HandlerWithError func(w http.ResponseWriter, r *http.Request) error

func ErrorHandler(h HandlerWithError) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err == nil {
			return
		}

		if appErr, ok := err.(*httperrors.APIError); ok {
			log.Printf("ERROR %s: %s", appErr.Source, appErr.Message)
			utils.WriteJSON(w, appErr.Code, model.APIResponse{
				Success: false,
				Message: appErr.Error(),
			})
			return
		}

		log.Printf("UNEXPECTED ERROR: %v", err)
		utils.WriteJSON(w, http.StatusInternalServerError, model.APIResponse{
			Success: false,
			Message: "Internal Server Error",
		})
	}
}
