package processor

import (
	"fmt"
	"image"
	_ "image/jpeg" // JPEG decoding
	_ "image/png"  // PNG decoding
	"math/rand"
	"net/http"
	"sync"
	"time"

	"backend-intern/models"
)

func ProcessJobs(jobQueue <-chan models.Job, jobStatus *sync.Map) {
	for job := range jobQueue {
		for _, visit := range job.Visits {
			for _, imageURL := range visit.ImageURLs {
				result, err := processImage(imageURL)
				if err != nil {
					job.Status = "failed"
					job.Error = append(job.Error, models.JobError{
						StoreID: visit.StoreID,
						Error:   err.Error(),
					})
					continue
				}
				fmt.Printf("Processed Image: %s, Perimeter: %f\n", imageURL, result)
			}
		}

		if job.Status != "failed" {
			job.Status = "completed"
		}
		jobStatus.Store(job.JobID, job)
	}
}

func processImage(url string) (float64, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("failed to download image: %v", err)
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("failed to decode image: %v", err)
	}

	bounds := img.Bounds()
	perimeter := 2 * float64(bounds.Dx()+bounds.Dy())

	// Simulate GPU processing
	time.Sleep(time.Duration(100+rand.Intn(300)) * time.Millisecond)

	return perimeter, nil
}
