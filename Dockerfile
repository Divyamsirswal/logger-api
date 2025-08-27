FROM golang:1.25.0 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o logger-api .

FROM debian:bookworm-slim 
WORKDIR /app
COPY --from=builder /app/logger-api .
EXPOSE 8080
CMD ["./logger-api"]
