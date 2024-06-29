FROM golang:1.21.0-alpine as builder

WORKDIR /go/gobel-api/app

ARG TARGETPLATFORM
ARG TARGETOS
ARG TARGETARCH

COPY . .

RUN apk update && \
    apk upgrade && \
    apk add --no-cache libc-dev gcc git openssh openssl bash

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o app

FROM --platform=$TARGETPLATFORM alpine

COPY --from=builder /go/gobel-api/app ./

CMD ["./app"]
