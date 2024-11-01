package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"syndio/contracts"
	"syndio/services"
)

func CreateEmployments(service services.SyndioService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// create new context
		ctx := context.Background()

		var req []contracts.EmploymentContract
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		if err := service.CreateEmployments(ctx, req); err != nil {
			log.Printf("Error creating employment records: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Employment records created successfully!"))
	}
}
