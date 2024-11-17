package handlers

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"backend-intern/models"
)

var (
	JobQueue = make(chan models.Job, 100) // Job queue
	Jobs     = sync.Map{}                 // Track jobs in memory
)

func SubmitJobHandler(w http.ResponseWriter, r *http.Request) {
	var payload models.JobRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, `{"error":"Invalid request payload"}`, http.StatusBadRequest)
		return
	}

	// Validate input
	if len(payload.Visits) != payload.Count {
		http.Error(w, `{"error":"Count mismatch with visits"}`, http.StatusBadRequest)
		return
	}

	jobID := uuid.New().String()
	job := models.Job{
		JobID:  jobID,
		Status: "ongoing",
		Visits: payload.Visits,
	}

	Jobs.Store(jobID, job) // Store job
	JobQueue <- job        // Enqueue job

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"job_id": jobID})
}
