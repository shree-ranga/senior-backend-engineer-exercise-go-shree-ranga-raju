package handlers

import (
	"context"
	"encoding/json"
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
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Employment records created successfully!"))
	}
}
