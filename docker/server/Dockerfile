FROM golang:1.18.0-alpine

ENV GO111MODULE=on

WORKDIR /go/src/grpc-web

RUN apk add --no-cache --update \
  git

RUN go install github.com/pilu/fresh@latest
CMD ["fresh"]