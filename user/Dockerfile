FROM golang

RUN mkdir -p /go/src

ADD go.mod /go/src/

ADD ./user /go/src/user

ADD ./database /go/src/database

ADD ./protobuf /go/src/protobuf

WORKDIR /go/src

ENV GO111MODULE=on

ENV DB_HOST=db

RUN go install all

ENTRYPOINT go run ./user/main/server.go -serverPort 8080

EXPOSE 6002
