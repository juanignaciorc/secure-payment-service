package transfer_processor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"
)

type PaymentProcessor struct {
	mu        sync.Mutex
	transfers map[string]struct{} // Para evitar duplicados
}

func NewPaymentProcessor() *PaymentProcessor {
	return &PaymentProcessor{
		transfers: make(map[string]struct{}),
	}
}

// Simular procesamiento de transferencia
func (p *PaymentProcessor) ProcessTransfer(w http.ResponseWriter, r *http.Request) {
	var transfer struct {
		TransferID string  `json:"transfer_id"`
		Amount     float64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&transfer); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	p.mu.Lock()
	if _, exists := p.transfers[transfer.TransferID]; exists {
		p.mu.Unlock()
		http.Error(w, "Transfer already being processed", http.StatusConflict)
		return
	}
	p.transfers[transfer.TransferID] = struct{}{}
	p.mu.Unlock()

	// Simular tiempo de procesamiento (entre 1 y 5 segundos)
	processingTime := time.Duration(rand.Intn(4)+1) * time.Second

	// Simular éxito/fallo (80% de éxito)
	success := rand.Float64() < 0.8

	// Enviar notificación después del tiempo de procesamiento
	go func() {
		time.Sleep(processingTime)
		p.SendNotification(transfer.TransferID, success)
	}()

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{
		"message":                   "Transfer processing started",
		"estimated_processing_time": fmt.Sprintf("%d seconds", processingTime.Seconds()),
	})
}

// Enviar notificación al webhook
func (p *PaymentProcessor) SendNotification(transferID string, success bool) {
	status := "COMPLETED"
	if !success {
		status = "FAILED"
	}

	payload := map[string]interface{}{
		"transfer_id": transferID,
		"status":      status,
		"message":     "Payment processed successfully",
	}

	// Convertir payload a JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("Error marshaling payload: %v\n", err)
		return
	}

	// Obtener URL del webhook de la variable de entorno
	webhookURL := os.Getenv("WEBHOOK_URL")
	if webhookURL == "" {
		fmt.Println("WEBHOOK_URL environment variable not set")
		return
	}

	// Hacer la llamada HTTP al webhook
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Printf("Error sending webhook: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Webhook failed with status: %d\n", resp.StatusCode)
	} else {
		fmt.Printf("Successfully sent notification to webhook: %v\n", payload)
	}
}
