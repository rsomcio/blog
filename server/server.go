package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/rsomcio/blog/blogpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Blog(ctx context.Context, req *blogpb.BlogRequest) (*blogpb.BlogResponse, error) {
	fmt.Printf("Blog function was invoked with %v\n", req)
	authorID := req.GetBlog().GetAuthorId()
	result := "Hello " + authorID
	res := &blogpb.BlogResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	// if we crash the co code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Blog Server Started.")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen")
	}

	s := grpc.NewServer()
	blogpb.RegisterBlogServiceServer(s, &server{})

	go func() {
		fmt.Println("Staring Server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Wait for Control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until signal is received
	<-ch
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("Closing the listener")
	lis.Close()
	fmt.Println("End of Program")
}
