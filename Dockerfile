FROM golang:latest

COPY . /golang

WORKDIR /golang

RUN go build cmd/main.go

CMD ["./main"]