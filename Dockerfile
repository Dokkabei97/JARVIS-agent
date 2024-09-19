FROM golang:1.20 AS builder

WORKDIR /home/app

COPY . .

RUN go mod tidy

WORKDIR /home/app/cmd

RUN go build -o main .

FROM alpine:latest

WORKDIR /home/app

COPY --from=builder /home/app/cmd/main .

CMD ["./main"]