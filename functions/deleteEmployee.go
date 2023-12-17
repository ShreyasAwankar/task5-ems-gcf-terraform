package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"cloud.google.com/go/logging"
)

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	defer logClient.Close()

	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "DELETE")
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
		http.Error(w, "Invalid Employee ID", http.StatusBadRequest)

		logger.Log(logging.Entry{
			Payload:  fmt.Sprintf("Invalid ID provided for employee: %v", err),
			Severity: logging.Error,
		})
		return
	}

	// Find the document to delete based on the employee ID
	query := client.Collection("ems").Where("ID", "==", employeeId)
	docs, err := query.Documents(ctx).GetAll()
	if err != nil {

		logger.Log(logging.Entry{
			Payload:  fmt.Sprintf("Error finding document: %v", err),
			Severity: logging.Error,
		})
		w.Header().Set("Content-Type", "text/plain")
		http.Error(w, "Failed to retrieve employee", http.StatusInternalServerError)
		return
	}

	if len(docs) != 1 {

		logger.Log(logging.Entry{
			Payload:  fmt.Sprintf("Employee fetched with id %v", employeeId),
			Severity: logging.Error,
		})
		w.Header().Set("Content-Type", "text/plain")
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	// Get the reference to the document to be deleted
	docRef := docs[0].Ref

	// Perform the deletion
	_, deleteErr := docRef.Delete(ctx)
	if deleteErr != nil {
		// logger.Errorf("Failed to delete employee: %v", deleteErr)
		logger.Log(logging.Entry{
			Payload:  fmt.Sprintf("Failed to delete employee: %v", deleteErr),
			Severity: logging.Error,
		})
		w.Header().Set("Content-Type", "text/plain")
		http.Error(w, "Failed to delete employee", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	// logger.Infof("Employee with employee ID %d deleted successfully", employeeId)
	logger.Log(logging.Entry{
		Payload:  fmt.Sprintf("Employee with employee ID %d deleted successfully", employeeId),
		Severity: logging.Info,
	})
}
