FROM golang:1.20-alpine

WORKDIR /app/clientFinal
COPY . .

RUN go mod download

CMD ["go", "run", "./cmd/inventory/main.go"]