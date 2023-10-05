FROM golang:1.19.2-alpine3.16 as setup

WORKDIR /go/src/github.com/htchan/ContainerResourceCollector

COPY go.mod go.sum .

RUN --mount=type=cache,target=$GOPATH/pkg \
    go mod download -x

COPY ./internal ./internal


FROM setup as builder

WORKDIR /go/src/github.com/htchan/ContainerResourceCollector

RUN apk add gcc musl-dev libc-dev

COPY ./cmd/main ./cmd/main

RUN --mount=type=cache,target=/root/.cache/go-build \
    go build -v ./cmd/main


FROM alpine

COPY --from=builder /go/src/github.com/htchan/ContainerResourceCollector/main .

CMD './main'