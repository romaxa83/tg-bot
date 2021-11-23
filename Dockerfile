#---Build stage---
FROM golang:1.16 AS builder
COPY . /go/src/tg-bot
WORKDIR /go/src/tg-bot/cmd/tg-bot-server

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags='-w -s' -o /go/bin/service

#---Final stage---
FROM alpine:latest
COPY --from=builder /go/bin/service /go/bin/service
CMD /go/bin/service --port 8080 --host '0.0.0.0'