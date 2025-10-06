
FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o user-service ./cmd/notification-service/main.go

EXPOSE 8080

CMD ["./notification-service"]