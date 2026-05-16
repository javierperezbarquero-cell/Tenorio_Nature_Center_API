FROM golang:1.26-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build -o main .

EXPOSE 8080

CMD ["./main"]