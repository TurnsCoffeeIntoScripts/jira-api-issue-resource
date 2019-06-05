#FROM golang:alpine AS builder
FROM ubuntu:18.04

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64
ENV GOMOD /app/go.mod

RUN apt-get update \
    && apt-get install make

COPY . /app
RUN make /app

#RUN go build -a -ldflags="-s -w" -o bin/jiraApiResource cmd/jira-api/main.go
#RUN make

FROM alpine:edge AS resource

RUN apk --no-cache add \
        curl \
        jq \
        bash \
;

COPY --from=builder bin/jiraApiResource /usr/local/bin/
COPY resources/ /opt/resource

FROM resource
