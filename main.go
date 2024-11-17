package main

import (
	"log"
	"net/http"

	"backend-intern/handlers"
	"backend-intern/models"
	"backend-intern/processor"
)

func main() {
	// Load store master data
	err := models.LoadStoreData("store_master.csv")
	if err != nil {
		log.Fatalf("Failed to load store master data: %v", err)
	}

	// Start worker pool
	go processor.ProcessJobs(handlers.JobQueue, &handlers.Jobs)

	// Set up routes
	http.HandleFunc("/api/submit", handlers.SubmitJobHandler)
	http.HandleFunc("/api/status", handlers.GetJobStatusHandler)

	// Start the server
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
