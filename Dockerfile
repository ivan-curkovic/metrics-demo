## dockerfile to run go app
FROM golang:1.22.1-alpine

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main .

CMD ["./main"]
