FROM golang:1.20 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o service_name ./cmd/service_name

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/service_name .
EXPOSE 8080
CMD ["./service_name"]