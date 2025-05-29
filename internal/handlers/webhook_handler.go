package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
	"secure-payment-service/internal/models"
	"secure-payment-service/internal/services"
)

type WebhookHandler struct {
	transferService *services.TransferService
	logger         *logrus.Logger
}

func NewWebhookHandler(transferService *services.TransferService, logger *logrus.Logger) *WebhookHandler {
	return &WebhookHandler{
		transferService: transferService,
		logger:         logger,
	}
}

// WebhookPayload representa el payload que recibiremos del webhook
// Este es un ejemplo simple, podrías ajustarlo según el formato real que necesites
// para tu gateway de pagos

type WebhookPayload struct {
	TransferID string                `json:"transfer_id"`
	Status     models.TransferStatus `json:"status"`
	Message    string               `json:"message,omitempty"`
}

func (h *WebhookHandler) HandleWebhook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	payload := &WebhookPayload{}
	if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
		h.logger.WithError(err).Error("failed to decode webhook payload")
		handleError(w, http.StatusBadRequest, err)
		return
	}

	h.logger.Printf("received webhook notification: transfer_id=%s status=%s", payload.TransferID, payload.Status)

	// Actualizar el estado de la transferencia
	err := h.transferService.UpdateTransferStatus(ctx, payload.TransferID, payload.Status)
	if err != nil {
		h.logger.WithError(err).Error("failed to update transfer status")
		handleError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Transfer status updated successfully",
	})
}
