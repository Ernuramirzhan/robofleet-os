FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o robofleet-api ./cmd/api


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/robofleet-api .

EXPOSE 8080

CMD ["./robofleet-api"]