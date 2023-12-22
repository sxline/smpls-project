FROM golang:1.21-alpine AS build

COPY . /go/src/api

WORKDIR /go/src/api/cmd

RUN go build -o app main.go

EXPOSE 8000

CMD ["./app"]