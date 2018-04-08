# 1) BUILD INGESTION API
FROM golang:alpine AS build-go
RUN apk --no-cache add git

WORKDIR /go/src/github.com/keelerh/omniscience

RUN go get -u github.com/golang/dep/...
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure -v --vendor-only

COPY ./pkg ./pkg
COPY ./protos ./protos
COPY ./cmd/api ./cmd/api

RUN cd cmd/api && go build -o api && cp api /tmp/


# 2) INSTALL DOCKERIZE
FROM alpine AS install-dockerize
RUN apk update && apk add wget
RUN wget https://github.com/jwilder/dockerize/releases/download/v0.6.1/dockerize-linux-amd64-v0.6.1.tar.gz
RUN tar -C /usr/local/bin -xzvf dockerize-linux-amd64-v0.6.1.tar.gz


# 3) BUILD FINAL IMAGE
FROM alpine
RUN apk --no-cache add ca-certificates

WORKDIR /app/server/

COPY --from=build-go /tmp/api /app/server/
COPY --from=install-dockerize /usr/local/bin/dockerize /usr/local/bin/dockerize

EXPOSE 50051
ENTRYPOINT ["dockerize", "-wait", "http://elasticsearch:9200", "-timeout", "30s", "./api"]
