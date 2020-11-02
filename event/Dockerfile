FROM golang

RUN mkdir -p /go/src

ADD go.mod /go/src/

ADD ./event /go/src/event

ADD ./database /go/src/database

ADD ./protobuf /go/src/protobuf

WORKDIR /go/src

ENV GO111MODULE=on

RUN go install all

ENTRYPOINT go run ./event/main/server.go

EXPOSE 6002
