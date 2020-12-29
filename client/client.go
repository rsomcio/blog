package main

import (
	"context"
	"fmt"
	"log"

	"github.com/rsomcio/blog/blogpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	doUnary(c)
}

func doUnary(c blogpb.BlogServiceClient) {
	fmt.Println("Starting to do a Undary RPC")
	req := &blogpb.BlogRequest{
		Blog: &blogpb.Blog{
			Id:       "5",
			AuthorId: "rsomcio",
			Title:    "When doves cry",
			Content:  "A Prince song.",
		},
	}
	res, err := c.Blog(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Blog RPC: %v", err)
	}
	log.Printf("Response from Blog: %v", res.Result)

}
