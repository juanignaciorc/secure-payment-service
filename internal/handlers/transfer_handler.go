package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"secure-payment-service/internal/models"
	"secure-payment-service/internal/services"
)

type TransferHandler struct {
	transferService *services.TransferService
}

func NewTransferHandler(transferService *services.TransferService) *TransferHandler {
	return &TransferHandler{
		transferService: transferService,
	}
}

func (h *TransferHandler) CreateTransfer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var transfer models.Transfer

	if err := json.NewDecoder(r.Body).Decode(&transfer); err != nil {
		handleError(w, http.StatusBadRequest, err)
		return
	}

	err := h.transferService.CreateTransfer(ctx, &transfer)
	if err != nil {
		handleError(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(transfer)
}

func (h *TransferHandler) GetTransfer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	transferID := r.URL.Query().Get("id")
	if transferID == "" {
		handleError(w, http.StatusBadRequest, fmt.Errorf("transfer ID is required"))
		return
	}

	transfer, err := h.transferService.GetTransfer(ctx, transferID)
	if err != nil {
		handleError(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transfer)
}

func (h *TransferHandler) UpdateTransferStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	transferID := r.URL.Query().Get("id")
	status := r.URL.Query().Get("status")

	if transferID == "" || status == "" {
		handleError(w, http.StatusBadRequest, fmt.Errorf("transfer ID and status are required"))
		return
	}

	err := h.transferService.UpdateTransferStatus(ctx, transferID, models.TransferStatus(status))
	if err != nil {
		handleError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *TransferHandler) GetAccountBalance(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	accountID := r.URL.Query().Get("account_id")
	if accountID == "" {
		handleError(w, http.StatusBadRequest, fmt.Errorf("account ID is required"))
		return
	}

	balance, err := h.transferService.GetAccountBalance(ctx, accountID)
	if err != nil {
		handleError(w, http.StatusInternalServerError, err)
		return
	}

	response := map[string]float64{"balance": balance}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func handleError(w http.ResponseWriter, statusCode int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
}
