FROM golang:1.19.2-alpine3.16 as builder

WORKDIR /usr/src/app

COPY main.go main.go

RUN go build main.go


FROM alpine

COPY --from=builder /usr/src/app/main .

CMD './main'