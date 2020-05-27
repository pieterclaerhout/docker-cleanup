FROM golang:alpine AS go-james

RUN apk update && apk add git && rm -rf /var/cache/apk/*
RUN GO111MODULE=on go get -u github.com/pieterclaerhout/go-james/cmd/go-james


FROM go-james AS mod-download

RUN mkdir -p /app

ADD go.mod /app
ADD go.sum /app

WORKDIR /app


RUN go mod download

FROM mod-download AS builder

ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 go-james build -v


FROM scratch

COPY --from=builder "/app/build/docker-cleanup" /

ENTRYPOINT ["/docker-cleanup"]