FROM golang

RUN mkdir -p /go/src

ADD go.mod /go/src/

ADD ./clients /go/src/clients

ADD ./peripheral /go/src/peripheral

ADD ./database /go/src/database

ADD ./protobuf /go/src/protobuf

WORKDIR /go/src

ENV GO111MODULE=on

RUN go install all

ENTRYPOINT go run ./peripheral/main/server.go -serverPort 8080

EXPOSE 6001
