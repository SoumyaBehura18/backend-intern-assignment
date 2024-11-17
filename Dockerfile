 
# Stage 1: Build the application
FROM golang:1.20 as builder

# Set the working directory in the container
WORKDIR /app

# Copy all files from the current directory to the container
COPY . .

# Download dependencies and build the Go application
RUN go mod tidy
RUN go build -o main .

# Stage 2: Create a smaller runtime image
FROM alpine:3.17

# Set the working directory in the final container
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose the port the app runs on
EXPOSE 8080

# Run the application
CMD ["./main"]
