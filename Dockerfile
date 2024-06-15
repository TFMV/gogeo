FROM golang:1.22 AS build

ADD . /app

WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o geocoding-service ./cmd/server

# Use a minimal Docker image to reduce the attack surface.
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=build /app/geocoding-service .

CMD ["./geocoding-service"]
