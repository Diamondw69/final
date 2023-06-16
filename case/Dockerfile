FROM golang:1.17-alpine

WORKDIR /app/case
COPY . .

RUN go mod download

CMD ["go", "run", "cmd/case/main.go"]