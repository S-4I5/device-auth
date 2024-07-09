FROM golang:1.22.5-alpine3.19 AS builder

RUN apk  update && \
    apk  add --no-cache gcc g++ libc-dev

WORKDIR /build

COPY ./device-service ./device-service
COPY ./user-service ./user-service

WORKDIR /build/device-service

RUN go mod download

RUN go build -o main ./cmd/main.go

EXPOSE 8081

CMD ["./main"]