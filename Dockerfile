# Stage 1: Build using a trusted Go image
FROM golang:1.23 AS builder

# Set environment variables for security and build efficiency
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Set the Current Working Directory
WORKDIR /app

# Copy the entire project into the container
COPY . .

# Download the dependencies
RUN go mod download

# Build the Go binary with optimizations for smaller size and ensure binary is executable
RUN go build -o loadgen -ldflags="-s -w" ./cmd/loadgen/ && chmod +x ./loadgen 

# Stage 2: Create a lightweight final image
FROM gcr.io/distroless/base-debian10

# Create a non-root user for the application
USER nonroot:nonroot

# Set the working directory
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/loadgen .

# Command to run the load generator, allowing runtime arguments
ENTRYPOINT ["./loadgen"]
