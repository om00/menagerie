# Use Go Alpine for small image
FROM golang:1.24-alpine

RUN apk add --no-cache curl netcat-openbsd

WORKDIR /app

# Copy go.mod and download dependencies first for caching
COPY go.mod ./
RUN go mod download

# Copy all project files
COPY . .

# Build the Go binary
RUN go build -o app .

# Expose port (optional, Docker Compose maps it anyway)
EXPOSE 8080

# Run the compiled binary
CMD ["./app"]
