FROM golang

RUN mkdir -p /go/src

ADD go.mod /go/src/

ADD ./clients /go/src/clients

ADD ./event /go/src/event

ADD ./database /go/src/database

ADD ./protobuf /go/src/protobuf

WORKDIR /go/src

ENV GO111MODULE=on

ENV DB_HOST=db

ENV PERIPHERAL_SERVER_ADDR="peripheral:8080"

RUN go install all

ENTRYPOINT go run ./event/main/server.go -serverPort 8080

EXPOSE 6004
