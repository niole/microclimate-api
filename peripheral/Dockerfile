FROM golang

RUN mkdir -p /go/src

ADD go.mod /go/src/

ADD ./shared /go/src/shared

ADD ./peripheral /go/src/peripheral

ADD ./protobuf /go/src/protobuf

WORKDIR /go/src

ENV GO111MODULE=on

ENV ENABLE_REFLECTION=$ENABLE_REFLECTION

RUN go install all

ENTRYPOINT go run ./peripheral/main/server.go \
-serverPort 8080 \
-withReflection=${ENABLE_REFLECTION:-false}

EXPOSE 6001
