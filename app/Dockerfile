FROM golang:1.17.1-alpine as builder

WORKDIR /go/gobel-api/app

COPY . .

RUN apk update && \
    apk upgrade && \
    apk add --no-cache libc-dev gcc git openssh openssl bash

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -o app

FROM --platform=linux/amd64 alpine

COPY --from=builder /go/gobel-api/app ./

CMD ["./app"]
