FROM golang:1.17-alpine

WORKDIR /app/authTest
COPY . .

RUN go mod download

CMD ["go", "run", "cmd/authService/main.go"]