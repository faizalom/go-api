# Build Stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go-api ./cmd/server

# Runtime Stage
FROM alpine:latest

WORKDIR /root/

# Copy the pre-built binary from the builder stage
COPY --from=builder /go-api .

# Copy the config file
COPY configs/config.docker.yaml .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./go-api"]
