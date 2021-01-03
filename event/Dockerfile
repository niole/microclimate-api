FROM golang

RUN mkdir -p /go/src

ADD go.mod /go/src/

ADD ./shared /go/src/shared

ADD ./event /go/src/event

ADD ./protobuf /go/src/protobuf

WORKDIR /go/src

ENV GO111MODULE=on

RUN go install all

ENV ENABLE_REFLECTION=$ENABLE_REFLECTION

ENTRYPOINT go run ./event/main/server.go \
-serverPort=8080 \
-withReflection=${ENABLE_REFLECTION:-false}

EXPOSE 6004
