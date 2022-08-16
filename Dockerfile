ARG GOVERSION=1.19
FROM golang:$GOVERSION-alpine

RUN apk add build-base

ENV GO111MODULE="on"        \
    CGO_ENABLED=1           \
    GOOS=linux              \
    GOARCH=amd64

WORKDIR /app

COPY . .

# Build
RUN  go build    \
    -ldflags="-X 'github.com/py4mac/fizzbuzz/pkg/config.Version=${VERSION}' -X 'github.com/py4mac/fizzbuzz/pkg/config.Revision=${VCS_REF}' -X 'github.com/py4mac/fizzbuzz/pkg/config.Built=${BUILD_DATE}'" \
    -a -o  fizzbuzz main.go
RUN chown 1001:1001 fizzbuzz

ENTRYPOINT ["/app/fizzbuzz"]