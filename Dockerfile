FROM golang:1.22.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Runtime image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY views views/ 
COPY css css/
COPY resources resources/

CMD ["./main"]
