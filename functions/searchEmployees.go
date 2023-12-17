package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"cloud.google.com/go/logging"
	"github.com/Task3/models"
)

func SearchEmployees(w http.ResponseWriter, r *http.Request) {

	mu.Lock()
	defer mu.Unlock()

	defer logClient.Close()

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

	collection := client.Collection("ems")

	// Get all documents in the collection
	docs, err := collection.Documents(ctx).GetAll()
	if err != nil {
		logger.Log(logging.Entry{
			Payload:  fmt.Sprintf("Error fetching all documents: %v", err),
			Severity: logging.Error,
		})
		w.Header().Set("Content-Type", "text/plain")
		http.Error(w, "Failed to retrieve documents", http.StatusInternalServerError)
		return
	}

	var employees []models.Employee

	query := r.URL.Query()

	// Get the search criteria from query parameters
	firstName := query.Get("firstName")
	lastName := query.Get("lastName")
	email := query.Get("email")
	role := query.Get("role")

	// Iterate through the documents and retrieve their data
	for _, doc := range docs {
		var emp models.Employee
		if err := doc.DataTo(&emp); err != nil {

			logger.Log(logging.Entry{
				Payload:  fmt.Sprintf("Error parsing employee data: %v", err),
				Severity: logging.Error,
			})
			continue
		}

		if (firstName == "" || emp.FirstName == firstName) &&
			(lastName == "" || emp.LastName == lastName) &&
			(email == "" || emp.Email == email) &&
			(role == "" || emp.Role == role) {
			employees = append(employees, emp)
		}

	}

	if len(employees) == 0 {
		logger.Log(logging.Entry{
			Payload:  ("No employee found with the provided search criteria"),
			Severity: logging.Info,
		})
	} else {
		logger.Log(logging.Entry{
			Payload:  ("Employees fetched with the provided search criteria"),
			Severity: logging.Info,
		})
	}

	// Serialize the results to JSON and send the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}
