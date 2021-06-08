FROM golang:1.16.0-alpine

WORKDIR /go/gobel-api/app

COPY . .

RUN apk update && \
    apk upgrade && \
    apk add --no-cache libc-dev gcc git openssh openssl bash && \
    GO111MODULE=off go get github.com/oxequa/realize && \
    go get -u golang.org/x/lint/golint

CMD realize start --run

