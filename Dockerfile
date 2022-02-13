# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /url-shortener ./cmd/server/main.go

EXPOSE 8080

CMD [ "/url-shortener" ]