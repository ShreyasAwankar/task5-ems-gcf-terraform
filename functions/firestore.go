package controllers

import (
	"context"
	"log"
	"sync"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/logging"
)

var mu sync.Mutex

var logClient *logging.Client
var ctx context.Context
var client *firestore.Client
var logger *logging.Logger

// Initializing firestore connection and cloud logging.
func init() {
	ctx = context.Background()

	logClient, err := logging.NewClient(ctx, "terraform-cloud-functions-ems")

	if err != nil {
		log.Fatalf("Failed to create logging client: %v", err)
	}

	// Setting the name of the log to write to.
	logName := "my-log"

	// Creates a logger for the specified log name.
	logger = logClient.Logger(logName)

	c, err := firestore.NewClient(ctx, "terraform-cloud-functions-ems")
	if err != nil {
		logger.Log(logging.Entry{
			Payload:  "Failed to create Firestore client",
			Severity: logging.Error,
		})
	}
	client = c
}
