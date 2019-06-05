FROM golang:alpine AS builder

COPY cmd/ cmd/
COPY pkg/ pkg/

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64

RUN go build -a -ldflags="-s -w" -o bin/jiraApiResource cmd/jira-api/main.go

FROM alpine:edge AS resource

RUN apk --no-cache add \
        curl \
        jq \
        bash \
;

COPY --from=builder bin/jiraApiResource /usr/local/bin/
COPY resources/ /opt/resource

FROM resource