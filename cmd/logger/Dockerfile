FROM golang:1.13.4-alpine3.10

WORKDIR $GOPATH/src/go-logger

COPY  ./pkg/ ./pkg/

COPY ./cmd/logger ./cmd/logger

COPY ./go.mod ./go.mod

ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn

RUN CGO_ENABLED=0 go build -mod=vendor -ldflags '-w -s' -o logger ./cmd/logger

FROM alpine:3.10.3

WORKDIR /logger/app

COPY --from=0 /go/src/go-logger/logger ./

WORKDIR /logger

ENTRYPOINT ["./app/logger"]

