package models

type JobRequest struct {
	Count  int     `json:"count"`
	Visits []Visit `json:"visits"`
}

type Visit struct {
	StoreID   string   `json:"store_id"`
	ImageURLs []string `json:"image_url"`
	VisitTime string   `json:"visit_time"`
}

type Job struct {
	JobID  string    `json:"job_id"`
	Status string    `json:"status"`
	Visits []Visit   `json:"visits"`
	Error  []JobError `json:"error,omitempty"`
}

type JobError struct {
	StoreID string `json:"store_id"`
	Error   string `json:"error"`
}

