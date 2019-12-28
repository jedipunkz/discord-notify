FROM golang:1.13-alpine3.10
MAINTAINER jedipunkz

WORKDIR /go/src/

ADD . /go/src/
ADD ./.discord-notify.yaml /root/

RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/discord-notify

ENTRYPOINT ["/go/bin/discord-notify"]
