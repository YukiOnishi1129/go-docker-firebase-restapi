FROM golang:1.21.4-alpine3.18

RUN apk update && apk add git
RUN mkdir /app

WORKDIR /app

COPY . /app

RUN go mod tidy

CMD ["go", "run", "./cmd/main.go"]