package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	transfer_processor "secure-payment-service/internal/transfer-processor"
)

func main() {
	port := flag.String("port", "8081", "Port to run the transfer processor on")
	flag.Parse()

	processor := transfer_processor.NewPaymentProcessor()

	// Endpoint para procesar transferencias
	http.HandleFunc("/process_transfer", processor.ProcessTransfer)

	// Endpoint para simular envío de notificaciones
	http.HandleFunc("/simulate_notification", func(w http.ResponseWriter, r *http.Request) {
		var notification struct {
			TransferID string `json:"transfer_id"`
			Status     string `json:"status"`
		}

		if err := json.NewDecoder(r.Body).Decode(&notification); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		processor.SendNotification(notification.TransferID, notification.Status == "COMPLETED")
		w.WriteHeader(http.StatusOK)
	})

	log.Printf("Starting transfer processor on port %s", *port)
	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		log.Fatal(err)
	}
}
