# install

- install go
- Clone project to $GOPATH directory.
- 
```sh
export SRC_DIR=./src DST_DIR=.
protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto
```
