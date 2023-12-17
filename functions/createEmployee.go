package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/logging"
	"github.com/Task3/models"
	"github.com/Task3/validations"
)

func CreateEmployee(w http.ResponseWriter, r *http.Request) {

	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	// Set CORS headers for the main request..
	w.Header().Set("Access-Control-Allow-Origin", "*")

	mu.Lock()
	defer mu.Unlock()

	var employee models.Employee
	err := json.NewDecoder(r.Body).Decode(&employee)

	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		http.Error(w, "Invalid JSON", http.StatusBadRequest)

		logger.Log(logging.Entry{
			Payload:  "Invalid JSON input for type Employee during controllers.CreateEmployee function call.",
			Severity: logging.Error,
		})
		w.Header().Set("Content-Type", "application/json")
		return
	}

	err1 := validations.V.Struct(employee)

	if err1 != nil {
		w.Header().Set("Content-Type", "text/plain")
		http.Error(w, "Invalid input for employee deatails\n-First name must contain alphabets or spaces only.\n-First name must contain alphabets or spaces only.\n-Email id must be valid eg. abc@example.com.\n-Password must be atleast 6 charecters long.\n-Phone no. should be valid.\n-Role must be either of - 'admin', 'developer', 'manager', 'tester'. (case sensetive)\n-Salary must be a number.\n-Birthdate should be in yyyy-mm-dd format.", http.StatusUnprocessableEntity)

		logger.Log(logging.Entry{
			Payload:  fmt.Sprintf("Invalid employee data input : occured while validating employee fields during controllers.CreateEmployee function call.\n%v", err1),
			Severity: logging.Error,
		})

		w.Header().Set("Content-Type", "application/json")
		return
	}

	// Query to get the maximum employee ID
	iter := client.Collection("ems").OrderBy("ID", firestore.Desc).Limit(1).Documents(ctx)

	var lastEmployee models.Employee

	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}
		doc.DataTo(&lastEmployee)
		// break
	}

	// Generate new employee ID
	EmpId := lastEmployee.ID + 1
	employee.ID = EmpId

	// Save employee data to Firestore
	_, _, err = client.Collection("ems").Add(ctx, employee)
	if err != nil {
		logger.Log(logging.Entry{
			Payload:  "Failed to create employee",
			Severity: logging.Error,
		})
		w.Header().Set("Content-Type", "text/plain")
		http.Error(w, "Failed to create employee", http.StatusInternalServerError)
		return
	}

	// Handle preflight request
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(employee)

	logger.Log(logging.Entry{
		Payload:  fmt.Sprintf("Employee created successfully with emp id %v", EmpId),
		Severity: logging.Info,
	})

}
