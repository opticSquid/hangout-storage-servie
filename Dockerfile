FROM golang:1.23-bookworm AS builder

WORKDIR /app

# Copy only necessary files
COPY go.mod go.sum ./
RUN go mod download
COPY . . 

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o main .

# Create a new stage for the final image
FROM ubuntu:noble

# Install necessary packages with optimizations
RUN apt-get update && apt-get install -y --no-install-recommends ffmpeg imagemagick && apt-get clean && rm -rf /var/lib/apt/lists/*

# Copy the built binary from the previous stage
COPY --from=builder /app/main /
WORKDIR /
CMD ["./main"]