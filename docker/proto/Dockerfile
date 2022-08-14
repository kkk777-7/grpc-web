FROM golang:1.18.0

ARG PROTO_VERSION=3.21.5
ARG PROTO_VERSION_DOWNLOAD=21.5
ARG GRPCWEB_VERSION=1.3.1

WORKDIR /proto

RUN apt-get -qq update && DEBIAN_FRONTEND=noninteractive apt-get -qq install -y \
  unzip

RUN curl -sSL https://github.com/protocolbuffers/protobuf/releases/download/v${PROTO_VERSION_DOWNLOAD}/protoc-${PROTO_VERSION_DOWNLOAD}-linux-aarch_64.zip -o protoc.zip && \
  unzip -qq protoc.zip && \
  cp ./bin/protoc /usr/local/bin/protoc && \
  cp -r ./include /usr/local

RUN curl -sSL https://github.com/grpc/grpc-web/releases/download/${GRPCWEB_VERSION}/protoc-gen-grpc-web-${GRPCWEB_VERSION}-linux-x86_64 -o /usr/local/bin/protoc-gen-grpc-web && \
  chmod +x /usr/local/bin/protoc-gen-grpc-web

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest