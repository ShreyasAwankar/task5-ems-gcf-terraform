package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"cloud.google.com/go/logging"
	"github.com/Task3/models"
	"google.golang.org/api/iterator"
)

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Set CORS headers for the main request.
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var employees []models.Employee

	iter := client.Collection("ems").Documents(ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			logger.Log(logging.Entry{
				Payload:  fmt.Sprintf("Failed to iterate through employees: %v", err),
				Severity: logging.Error,
			})
			http.Error(w, "Failed to retrieve employees", http.StatusInternalServerError)
			return
		}

		var employee models.Employee

		err = doc.DataTo(&employee)
		if err != nil {
			// logger.Errorf("Failed to parse employee data: %v", err)
			logger.Log(logging.Entry{
				Payload:  fmt.Sprintf("Failed to parse employee data: %v", err),
				Severity: logging.Error,
			})
			w.Header().Set("Content-Type", "text/plain")
			http.Error(w, "Failed to retrieve employees", http.StatusInternalServerError)
			return
		}

		employees = append(employees, employee)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employees)

	logger.Log(logging.Entry{
		Payload:  "Employees fetched successfully",
		Severity: logging.Info,
	})
}
