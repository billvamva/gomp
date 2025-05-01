# Stage 1: Build the Go application
FROM golang:1.22 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download all the dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code to the working directory
COPY . .

# Navigate to the directory containing main.go
WORKDIR /app/cmd/portfolio

# Build the Go application with static linking
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/portfolio .

# Stage 2: Create a minimal container with the Go application
FROM gcr.io/distroless/base-debian11 AS build-release-stage

# Set the working directory inside the container
WORKDIR /

# Copy the binary from the builder stage
COPY --from=builder /app/portfolio .


# Copy necessary static files
COPY --from=builder /app/web /web
COPY --from=builder /app/internal/templates /internal/templates
COPY --from=builder /app/.env .env

# Expose the port that your application will run on
EXPOSE 8080

# Command to run the executable
ENTRYPOINT  ["./portfolio"]
