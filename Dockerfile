# Use an official Golang image as the base image
FROM golang:latest

# Set the working directory to /app
WORKDIR /inventory

# Copy the Go module file
COPY go.mod ./

# Install dependencies
RUN go mod download

# Copy the Go program
COPY . .

# Build the Go program
RUN go build -o main main.go

# Run the Go program when the container starts
CMD ["go","run","main.go"]