# Step 1: Build the Go application
FROM golang:1.22.2-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o url-shortener .

# Step 2: Run the Go application
FROM alpine:3.18

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/url-shortener /app/url-shortener

# Command to run the executable
CMD ["/app/url-shortener"]

# Expose port 3000 to the outside world
EXPOSE 3000
