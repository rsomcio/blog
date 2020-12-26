#!/bin/bash
(cd ~/go/src;
protoc --go_out=plugins=grpc:. github.com/rsomcio/blog/blogpb/blogpb.proto)
