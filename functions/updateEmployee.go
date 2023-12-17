package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"cloud.google.com/go/logging"
	"github.com/Task3/models"
	"github.com/Task3/validations"
)

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	defer logClient.Close()

	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "PUT")
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
		http.Error(w, "Invalid Employee ID. Employee Id must be a digit", http.StatusBadRequest)

		logger.Log(logging.Entry{
			Payload:  fmt.Sprintf("Invalid ID provided for employee: %v", err),
			Severity: logging.Error,
		})
		return
	}

	var employee models.Employee
	err = json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		w.Header().Set("Content-Type", "text/plain")

		logger.Log(logging.Entry{
			Payload:  fmt.Sprintf("Invalid JSON input for employee: %v", err),
			Severity: logging.Error,
		})
		return
	}

	employee.ID = employeeId

	err = validations.V.Struct(employee)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		http.Error(w, "Invalid input for employee details\n-First name must contain alphabets or spaces only.\n-First name must contain alphabets or spaces only.\n-Email id must be valid eg. abc@example.com.\n-Password must be atleast 6 charecters long.\n-Phone no. should be valid.\n-Role must be either of - 'admin', 'developer', 'manager', 'tester'. (case sensetive)\n-Salary must be a number.\n-Birthdate should be in yyyy-mm-dd format.", http.StatusUnprocessableEntity)

		logger.Log(logging.Entry{
			Payload:  fmt.Sprintf("Invalid employee data input: %v", err),
			Severity: logging.Error,
		})
		return
	}

	// Find the document based on a field other than the document ID
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

	if len(docs) == 0 {

		logger.Log(logging.Entry{
			Payload:  ("Employee not found"),
			Severity: logging.Info,
		})
		w.Header().Set("Content-Type", "text/plain")
		http.Error(w, "Employee not found ", http.StatusNotFound)
		return
	}

	docRef := docs[0].Ref

	// Perform the update
	_, updateErr := docRef.Set(ctx, employee)
	if updateErr != nil {

		logger.Log(logging.Entry{
			Payload:  fmt.Sprintf("Failed to update employee: %v", updateErr),
			Severity: logging.Error,
		})
		w.Header().Set("Content-Type", "text/plain")
		http.Error(w, "Failed to update employee", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employee)

	logger.Log(logging.Entry{
		Payload:  fmt.Sprintf("Employee with employee ID %d updated successfully", employeeId),
		Severity: logging.Info,
	})
}
