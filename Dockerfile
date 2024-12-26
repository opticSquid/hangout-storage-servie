FROM golang:1.23-bookworm AS builder

# Set the working directory
WORKDIR /app

# Copy the Go source code to the working directory
COPY . .

# Build the Go application
RUN go build -o main .

# Create a new stage for the final image
FROM ubuntu:noble

# Install necessary packages
RUN apt-get update && apt-get install -y ffmpeg imagemagick gpac

# Copy the built binary from the previous stage
COPY --from=builder /app/main /

# Set the working directory
WORKDIR /

# Define the command to run the application
CMD ["./main"]
