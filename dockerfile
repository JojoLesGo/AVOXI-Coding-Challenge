FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache git && \
    go get github.com/oschwald/geoip2-golang

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on

COPY . .
# Run container
FROM alpine:latest

RUN apk --no-cache add ca-certificates

#RUN mkdir /app
#WORKDIR /app
#COPY --from=builder /app/blacklistcheck .

#CMD ["./blacklistcheck"]