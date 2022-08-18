ARG GOVERSION=1.19
FROM golang:$GOVERSION-alpine as builder
ARG GOVERSION=1.19
ARG REVISION
ARG BUILD_DATE
ARG VERSION

RUN apk add build-base

ENV GO111MODULE="on"        \
    CGO_ENABLED=1           \
    GOOS=linux              \
    GOARCH=amd64

WORKDIR /app

COPY . .

# Build
RUN  go build    \
    -ldflags="-X 'github.com/py4mac/fizzbuzz/pkg/constants.Version=${VERSION}' -X 'github.com/py4mac/fizzbuzz/pkg/constants.Revision=${REVISION}' -X 'github.com/py4mac/fizzbuzz/pkg/constants.Built=${BUILD_DATE}'" \
    -a -o  fizzbuzz /app/cmd/main.go
RUN chown 1001:1001 fizzbuzz


FROM alpine
WORKDIR /app
COPY --from=builder /app/fizzbuzz .
USER 1001:1001

RUN ls /app
ENTRYPOINT ["/app/fizzbuzz"]