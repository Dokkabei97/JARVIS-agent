FROM golang:1.20

WORKDIR /home/app

COPY . .

RUN go mod tidy

WORKDIR /home/app/cmd

RUN go build -o main .

CMD ["/home/app/cmd/main"]