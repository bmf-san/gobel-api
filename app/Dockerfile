FROM golang:1.20.0-alpine as builder

WORKDIR /go/gobel-api/app

COPY . .

RUN apk update && \
    apk upgrade && \
    apk add --no-cache libc-dev gcc git openssh openssl bash

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=${TARGETARCH}

RUN go build -o app

FROM --platform=$BUILDPLATFORM  alpine

COPY --from=builder /go/gobel-api/app ./

CMD ["./app"]
