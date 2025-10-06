FROM golang:1.20-slim

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/notification-service .

EXPOSE 8080

CMD ["/app/notification-service"]