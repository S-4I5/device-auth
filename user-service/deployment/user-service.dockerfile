FROM golang:1.22.5-alpine3.19 AS builder

RUN apk  update && \
    apk  add --no-cache gcc g++ libc-dev

WORKDIR /build

ADD ../go.mod .

COPY .. .

RUN go build -o main ./cmd/main.go

EXPOSE 8080
EXPOSE 50051
EXPOSE 50052

CMD ["./main"]