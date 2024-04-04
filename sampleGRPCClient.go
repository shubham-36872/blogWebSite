package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "blogApp/protos"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewBlogServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Create Post
	createResponse, err := c.CreatePost(ctx, &pb.Post{
		Title:           "Example Post",
		Content:         "This is an example post content.",
		Author:          "John Doe",
		PublicationDate: time.Now().Format(time.RFC3339),
		Tags:            []string{"example", "golang", "grpc"},
	})
	if err != nil {
		log.Fatalf("could not create post: %v", err)
	}
	fmt.Printf("Created Post: %+v\n", createResponse)

	// Create Post2
	createResponse2, err2 := c.CreatePost(ctx, &pb.Post{
		Title:           "New Post",
		Content:         "Coming Soon.",
		Author:          "Chetan bhagat",
		PublicationDate: time.Now().Format(time.RFC3339),
		Tags:            []string{"new", "indian", "grpc"},
	})
	if err2 != nil {
		log.Fatalf("could not create post: %v", err)
	}
	fmt.Printf("Created Post: %+v\n", createResponse2)

	// Read Post1
	readResponse, err := c.ReadPost(ctx, &pb.PostID{PostId: createResponse.GetPostId()})
	if err != nil {
		log.Fatalf("could not read post: %v", err)
	}
	fmt.Printf("Read Post: %+v\n", readResponse)

	// Read Post2
	readResponse2, err := c.ReadPost(ctx, &pb.PostID{PostId: createResponse2.GetPostId()})
	if err != nil {
		log.Fatalf("could not read post: %v", err)
	}
	fmt.Printf("Read Post: %+v\n", readResponse2)

	// Update Post1
	updateResponse, err := c.UpdatePost(ctx, &pb.Post{
		PostId:          createResponse.GetPostId(),
		Title:           "Updated Example Post",
		Content:         "This is the updated content of the example post.",
		Author:          "Jane Smith",
		PublicationDate: time.Now().Format(time.RFC3339),
		Tags:            []string{"updated", "example", "golang", "grpc"},
	})
	if err != nil {
		log.Fatalf("could not update post: %v", err)
	}
	fmt.Printf("Updated Post: %+v\n", updateResponse)

	// Example: Delete Post1
	deleteResponse, err := c.DeletePost(ctx, &pb.PostID{PostId: createResponse.GetPostId()})
	if err != nil {
		log.Fatalf("could not delete post: %v", err)
	}
	fmt.Printf("Delete Post Response: %+v\n", deleteResponse)
}
