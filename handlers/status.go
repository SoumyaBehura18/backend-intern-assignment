package handlers

import (
	"encoding/json"
	"net/http"
)

func GetJobStatusHandler(w http.ResponseWriter, r *http.Request) {
	jobID := r.URL.Query().Get("jobid")
	if jobID == "" {
		http.Error(w, `{"error":"jobid is required"}`, http.StatusBadRequest)
		return
	}

	value, exists := Jobs.Load(jobID)
	if !exists {
		http.Error(w, `{"error":"Job ID not found"}`, http.StatusBadRequest)
		return
	}

	job := value.(models.Job)
	json.NewEncoder(w).Encode(job)
}
