FROM golang:alpine AS builder

COPY go/ /go/
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64

RUN go build -a -ldflags="-s -w" -o /go/bin/updateIssue cmd/update-issue/main.go

FROM alpine:edge AS resource

RUN apk --no-cache add \
        curl \
        jq \
        bash \
;

COPY --from=builder /go/bin/updateIssue /usr/local/bin/
COPY resources/ /opt/resource

FROM resource