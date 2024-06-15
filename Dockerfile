# Use the official Golang image to create a build artifact.
FROM golang:1.22 AS build

# Copy the local package files to the container's workspace.
ADD . /app

# Set the working directory.
WORKDIR /app

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o geocoding-service ./cmd/server

# Use a minimal Docker image to reduce the attack surface.
FROM alpine:latest

RUN apk --no-cache add ca-certificates

# Set the working directory.
WORKDIR /root/

# Copy the build artifact from the build stage.
COPY --from=build /app/geocoding-service .

# Command to run the executable
CMD ["./geocoding-service"]
