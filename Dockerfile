FROM golang:1.20.5-buster AS builder

WORKDIR /go/src/app

COPY Makefile ./

RUN make deps-openapi

COPY go.sum go.mod ./

RUN make download

COPY . .

RUN  make openapi build

FROM alpine:3.19.0

WORKDIR /opt/app

COPY --from=builder /go/src/app/bin /opt/app

RUN chown -R nobody:nogroup /opt/app
RUN chmod +x /opt/app/*
USER nobody:nogroup