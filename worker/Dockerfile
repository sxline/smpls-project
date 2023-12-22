FROM golang:1.21-alpine AS build

# Копируем все файлы из текущего контекста сборки в каталог /go/src/api
COPY . /go/src/api

# Копируем data.json внутрь образа
COPY data.json /go/src/api/cmd/data.json
COPY data_test.json /go/src/api/cmd/data_test.json

WORKDIR /go/src/api/cmd

RUN go build -o app main.go

CMD ["./app"]