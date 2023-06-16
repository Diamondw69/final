FROM golang:1.20-alpine

WORKDIR /app/authTest
COPY . .

RUN go mod download

CMD ["go", "run", "./cmd/authService/main.go"]