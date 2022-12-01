FROM golang:1.17-alpine as builder
ARG VERSION=0.0.2

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app

ENV GOPROXY="https://goproxy.cn,direct"

COPY . .
RUN go build -ldflags "-s -w -X main.version=$VERSION" -o amrom-scheduler

FROM gcr.io/google_containers/ubuntu-slim:0.14
COPY --from=builder /app/amrom-scheduler /usr/bin/amrom-scheduler
ENTRYPOINT ["amrom-scheduler"]
