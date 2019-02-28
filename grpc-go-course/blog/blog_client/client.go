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

	// create blog
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
	blogID := createBlogResp.GetBlog().GetId()

	// read blog
	fmt.Println("Reading the blog")
	req1 := &blogpb.ReadBlogRequest{BlogId: "5c77989a34dbba2581ce80a4"}
	_, err2 := c.ReadBlog(context.Background(), req1)
	if err2 != nil {
		fmt.Printf("Error happened while reading: %v\n", err2)
	}

	req2 := &blogpb.ReadBlogRequest{BlogId: blogID}
	readBlogRes, err3 := c.ReadBlog(context.Background(), req2)
	if err3 != nil {
		fmt.Printf("Error happened while reading: %v\n", err3)
	}
	fmt.Printf("ReadBlog was called: %v\n", readBlogRes)

	// update blog
	fmt.Println("Updating the blog")
	req3 := &blogpb.UpdateBlogRequest{
		Blog: &blogpb.Blog{
			Id:       blogID,
			AuthorId: "Shivalik Narad",
			Title:    "First Blog changed to Second Blog",
			Content:  "Blog Content and some more blog content",
		},
	}

	updateRes, updateErr := c.UpdateBlog(context.Background(), req3)
	if updateErr != nil {
		fmt.Printf("Error while updating the blog: %v\n", updateErr)
	}
	fmt.Printf("Blog was updated: %v\n", updateRes)

	// delete blog
	req4 := &blogpb.DeleteBlogRequest{BlogId: blogID}
	deleteRes, deleteErr := c.DeleteBlog(context.Background(), req4)
	if deleteErr != nil {
		fmt.Printf("Error while updating the blog: %v\n", deleteErr)
	}
	fmt.Printf("Blog was deleted: %v\n", deleteRes)
}
