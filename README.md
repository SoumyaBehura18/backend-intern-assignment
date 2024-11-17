# Backend Intern Assignment

## Description
A service to process thousands of images collected from stores.

### Features
1. Submit jobs with image URLs and store IDs.
2. Process jobs concurrently with perimeter calculation.
3. Retrieve job statuses.

---

## Setup Instructions

### Without Docker
1. Clone the repository.
2. Install dependencies: `go mod tidy`.
3. Run the application: `go run main.go`.

### With Docker
1. Build the image: `docker build -t backend-assignment .`.
2. Run the container: `docker run -p 8080:8080 backend-assignment`.

---

## Testing
Use Postman or `curl` to test endpoints:
1. Submit Job: `POST /api/submit`
2. Get Job Status: `GET /api/status?jobid=<job_id>`

---

## Future Improvements
- Add database for persistent storage.
- Implement retry mechanisms for failed downloads.
- Add metrics and monitoring.
