#FROM golang:alpine AS builder
FROM ubuntu:18.04

COPY cmd/ cmd/
COPY pkg/ pkg/
COPY go.mod go.mod
COPY Makefile Makefile

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64
ENV GOMOD go.mod

#RUN go build -a -ldflags="-s -w" -o bin/jiraApiResource cmd/jira-api/main.go
RUN make

FROM alpine:edge AS resource

RUN apk --no-cache add \
        curl \
        jq \
        bash \
;

COPY --from=builder bin/jiraApiResource /usr/local/bin/
COPY resources/ /opt/resource

FROM resource