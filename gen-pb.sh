#/bin/bash

protoc -I. -I${GOPATH}/src --go_out=plugins=grpc:./pb/ --validate_out="lang=go:./pb/" main.proto