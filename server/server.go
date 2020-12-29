package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/rsomcio/blog/blogpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Blog(ctx context.Context, req *blogpb.BlogRequest) (*blogpb.BlogResponse, error) {
	fmt.Printf("Blog function was invoked with %v\n", req)
	authorId := req.GetBlog().GetAuthorId()
	result := "Hello " + authorId
	res := &blogpb.BlogResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello World!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen")
	}

	s := grpc.NewServer()
	blogpb.RegisterBlogServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
