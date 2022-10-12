FROM golang:1.18

WORKDIR /app

ADD . .

RUN go install ./cmd/server

CMD server
