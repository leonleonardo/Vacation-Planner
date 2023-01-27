FROM golang:1.19-alpine AS base
WORKDIR /go/src/app

ENV GO111MODULE=on

# Development build with hot reloding using Air 
FROM base AS dev
WORKDIR /go/src/app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./

RUN air -v

RUN go mod download

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]