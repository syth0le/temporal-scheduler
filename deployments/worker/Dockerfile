FROM golang:1.23.4-alpine as builder

WORKDIR /usr/src/app

COPY . .
RUN go mod download && go mod verify

RUN CGO_ENABLED=0 go build -o /usr/local/bin/worker ./cmd/worker/main.go

EXPOSE 8080