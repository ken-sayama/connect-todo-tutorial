FROM golang:1.22.0-alpine

RUN apk update && apk add git bash protobuf-dev

WORKDIR /app

RUN go install github.com/bufbuild/buf/cmd/buf@latest
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest

# ドキュメント生成用
# RUN go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@latest