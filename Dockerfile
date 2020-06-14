FROM golang:1.13.7-alpine3.11 AS build

ENV GOPROXY=https://goproxy.io

RUN mkdir -p /go/src/github.com/linkingthing/pg-agent
COPY . /go/src/github.com/linkingthing/pg-agent

WORKDIR /go/src/github.com/linkingthing/pg-agent
RUN CGO_ENABLED=0 GOOS=linux go build cmd/pg-agent/pg-agent.go

FROM alpine:3.9.4
COPY --from=build /go/src/github.com/linkingthing/pg-agent/pg-agent /
ENTRYPOINT ["/pg-agent"]
