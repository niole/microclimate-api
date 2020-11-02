FROM golang

RUN mkdir -p /go/src

ADD go.mod /go/src/

ADD ./deployment /go/src/deployment

ADD ./database /go/src/database

ADD ./protobuf /go/src/protobuf

WORKDIR /go/src

ENV GO111MODULE=on

ENV DB_HOST=db

RUN go install all

ENTRYPOINT go run ./deployment/main/server.go

EXPOSE 6003
