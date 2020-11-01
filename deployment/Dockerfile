FROM golang

RUN mkdir -p /go/src/api/server/impl/deployment

RUN mkdir -p /go/src/api/protobuf
# Copy the local package files to the container's workspace.
ADD . /server/impl/deployment

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/golang/example/outyet

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/outyet

# Document that the service listens on port 8080.
EXPOSE 6003
