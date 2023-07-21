# 1.Build stage

# Use the official Go image as the base image
FROM golang:1.20-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the entire project into the container's working directory
COPY . .

# Install Go dependencies
RUN go mod download

# Build the Go application
RUN go build -o main cmd/main.go

# Download and extract Golang Migrate tool for database migrations
RUN apk add --no-cache curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz

# 2.Run stage
FROM alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the compiled Go binary from the build stage
COPY --from=builder /app/main .

# Copy the Golang Migrate tool from the build stage
COPY --from=builder /app/migrate ./migrate

# Copy the configuration file
COPY app.env .

# Copy start-up script
COPY start.sh .

# Copy wait-for script for handling database dependencies
COPY wait-for.sh .

# Copy SQL migration files
COPY pkg/db/orm/migration ./migration

# Expose the port your application listens on
EXPOSE 8080

# Run the application when the container starts
CMD ["/app/main"]

# Set the entrypoint to the start-up script
ENTRYPOINT ["/app/start.sh"]
