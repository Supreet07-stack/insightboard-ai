FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY cmd ./cmd
COPY internal ./internal

RUN go build -o gateway ./cmd/server

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /app/gateway ./gateway

EXPOSE 8080

CMD ["./gateway"]