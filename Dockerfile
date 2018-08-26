FROM golang:1.10.3-alpine3.8

WORKDIR /go/src/github.com/hanks/terraform-variables-generator

RUN apk add --no-cache git gcc && \
    go get -u github.com/derekparker/delve/cmd/dlv && \
    go get github.com/golang/lint/golint && \
    mkdir -p ./dist/bin

CMD ["sh"]
