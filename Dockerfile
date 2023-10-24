FROM golang:1.21.1

WORKDIR /usr/local/app

RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go mod tidy


