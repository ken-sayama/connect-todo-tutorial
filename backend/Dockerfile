FROM golang:1.22.0-alpine as builder

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /go/src/app
COPY . .

RUN apk update && apk add git

RUN go get
RUN go build -o /app

# for local development
#FROM golang:1.22.0-alpine AS dev
#ENV CGO_ENABLED=0
#WORKDIR /go/src/app
#COPY . .
#
#RUN apk update && apk add git
#
#RUN go install github.com/go-delve/delve/cmd/dlv@latest
#RUN go install github.com/cosmtrek/air@latest
#
#CMD ["air", "-c", ".air.toml", "server"]

FROM golang:1.22.0-alpine
ENV CGO_ENABLED=0
WORKDIR /go/src/app
COPY . .

RUN apk update && apk add git

RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/cosmtrek/air@latest

CMD ["air", "-c", ".air.toml", "server"]