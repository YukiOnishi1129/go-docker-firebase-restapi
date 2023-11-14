FROM golang:1.21.4-alpine3.18

ARG GO111MODULE=on

RUN apk update && apk add git
RUN mkdir /app
# ワーキングディレクトリの設定
WORKDIR /app

COPY /app /app

RUN go mod tidy && \
    go install github.com/cosmtrek/air@v1.46.0