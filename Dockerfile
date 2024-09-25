## dockerfile to run go app
FROM golang:1.23.1-alpine

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main .

CMD ["./main"]
