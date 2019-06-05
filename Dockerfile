FROM golang:1.12 AS builder
#FROM ubuntu:18.04 AS builder

# Copy everything from the jira-api-ressource module to /app in the image
COPY . /app

# Set the working directory to /app where our code and scripts are
WORKDIR /app

# Environment variables
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOARCH amd64
ENV GOMOD /app/go.mod

#RUN go build \
#    -tags release \
#    -ldflags '-X jiraApiResource/cmd.Version=v0.0.1-rc.2-4-gb20652f-dirty -X jiraApiResource/cmd.BuildData=2019-06-05' -o ./bin/jiraApiResource cmd/jira-api/main.go

#RUN go build -a -ldflags="-s -w" -o bin/jiraApiResource cmd/jira-api/main.go
RUN make

FROM alpine:edge AS resource

RUN apk --no-cache add \
        curl \
        jq \
        bash \
;

COPY --from=builder /app/bin/jiraApiResource /usr/local/bin/
COPY resources/ /opt/resource

FROM resource
