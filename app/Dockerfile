FROM golang:1.13.0-alpine

RUN apk update && \
    apk upgrade && \
    apk add --no-cache libc-dev gcc git openssh openssl bash && \
    GO111MODULE=off go get github.com/oxequa/realize && \
    go get -u golang.org/x/lint/golint

WORKDIR /go/src/github.com/bmf-san/gobel-api/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -o app

CMD realize start --run

