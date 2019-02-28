package main

import (
	"context"
	"fmt"
	"log"

	"github.com/snarad/golang-examples/grpc-go-course/blog/blogpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Blog Client")
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("unable to connect")
	}

	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	fmt.Println("Creating the blog")
	req := &blogpb.CreateBlogRequest{
		Blog: &blogpb.Blog{
			AuthorId: "Shivalik",
			Title:    "First Blog",
			Content:  "Blog Content",
		},
	}

	createBlogResp, err := c.CreateBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while creating the blog: %v", err)
	}
	fmt.Printf("Blog Response: %v\n", createBlogResp.Blog)
}
