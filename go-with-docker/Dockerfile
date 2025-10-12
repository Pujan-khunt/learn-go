# Base image for build stage.
FROM golang:1.25 as build

# Set the working directory for the build stage.
WORKDIR /app

# Copy this file first to leverage `Docker layer cache` to reduce build time.
COPY go.mod .
RUN go mod download

# Copy (`*.go`) source code files.
COPY . .

# Build the application, disabling CGO to create a static binary.
RUN CGO_ENABLED=0 go build -v -o /app/go-api

# Using a simple image for the final stage.
FROM alpine:latest

# Create a non-root user to run the application.
RUN addgroup -S appgroup && adduser -S appuser -G appgroup # All in one line to reduce layers

# Switch to the newly created user.
USER appuser

# Set working directory for the final stage.
WORKDIR /home/appuser

# Copy the built binary from the build stage.
COPY --from=build /app/go-api .

# Expose the port that the HTTP server runs on.
EXPOSE 8080

# Run the Go binary
CMD ["./go-api"]
