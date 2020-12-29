#!/bin/bash
(cd ~/go/src;
protoc github.com/rsomcio/blog/blogpb/blogpb.proto --go_out=plugins=grpc:.)
