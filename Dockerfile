FROM golang:1.14 as builder

WORKDIR /go/src/app

COPY . .

RUN echo "Fetch dependencies..." && \
    go get -d -v ./... && \
    echo "Building app..." && \
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/opiapi && \
    echo "Done"

FROM alpine:latest
WORKDIR  /home/opiapi/

RUN apk --update --no-cache add bash

COPY --from=builder /go/src/app/build/opiapi .
COPY --from=builder /go/src/app/data data/


CMD ["/home/opiapi/opiapi"]