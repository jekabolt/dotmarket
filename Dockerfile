FROM golang:1.12-alpine AS builder

ENV GO111MODULE=on

RUN apk add --no-cache git libgit2-dev alpine-sdk

WORKDIR /go/src/github.com/jekabolt/dotmarket/

# https://divan.github.io/posts/go_get_private/
COPY .gitconfig /root/.gitconfig
COPY go.mod .
COPY go.sum .
# install dependencies
RUN go mod download

COPY ./cmd/ ./cmd/

COPY ./routers/ ./routers/
RUN CGO_ENABLED=1 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o /go/bin/dotmarket ./cmd/
RUN ls /go/bin

EXPOSE 8080

CMD ["obd-server"]