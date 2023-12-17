package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"cloud.google.com/go/logging"
	"github.com/Task3/models"
)

func GetEmployee(w http.ResponseWriter, r *http.Request) {

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

	// Fetch employee ID from the query parameters
	empID := r.URL.Query().Get("empid")

	employeeId, err := strconv.Atoi(empID)

	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		http.Error(w, "Invalid Employee ID. Employee ID must be a digit.", http.StatusBadRequest)

		logger.Log(logging.Entry{
			Payload:  ("Invalid id was employeeId was provided on GetEmployee function call."),
			Severity: logging.Error,
		})

		return
	}

	iter := client.Collection("ems").Where("ID", "==", employeeId).Documents(ctx)

	defer iter.Stop()

	doc, err := iter.Next()
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		logger.Log(logging.Entry{
			Payload:  fmt.Sprintf("Employee not found in db: %v", err),
			Severity: logging.Error,
		})
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	var employee models.Employee
	err = doc.DataTo(&employee)

	if err != nil {
		w.Header().Set("Content-Type", "text/plain")

		logger.Log(logging.Entry{
			Payload:  fmt.Sprintf("Employee not found in db: %v", err),
			Severity: logging.Error,
		})
		http.Error(w, "Failed to retrieve employee", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employee)

	logger.Log(logging.Entry{
		Payload:  fmt.Sprintf("Employee fetched with id %v", employeeId),
		Severity: logging.Info,
	})

}
