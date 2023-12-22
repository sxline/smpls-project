FROM golang:1.21-alpine AS build

COPY . /go/src/grpc-service

WORKDIR /go/src/grpc-service/cmd

RUN go build -o app main.go

EXPOSE 50051

CMD ["./app"]