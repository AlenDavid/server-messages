FROM golang:1.19-alpine

WORKDIR /app

ADD . .

RUN go install ./cmd/server

CMD server
