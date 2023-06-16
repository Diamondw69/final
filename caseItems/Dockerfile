FROM golang:1.20-alpine

WORKDIR /app/caseItems
COPY . .

RUN go mod download

CMD ["go", "run", "./cmd/caseItems/main.go"]