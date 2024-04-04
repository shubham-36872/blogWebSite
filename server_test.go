package main_test

import (
	"context"
	"testing"

	pb "blogApp/protos"
	"main"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestCRUDOperations(t *testing.T) {

	server := main.NewServer()
	go func() {
		if err := server.Serve(":9090"); err != nil {
			t.Fatalf("failed to start server: %v", err)
		}
	}()
	defer server.Stop()

	// Create gRPC client connection
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial server: %v", err)
	}
	defer conn.Close()
	client := pb.NewBlogServiceClient(conn)

	// Testing CreatePost
	post := &pb.Post{
		Title:           "Test Post",
		Content:         "Test Content",
		Author:          "Test Author",
		PublicationDate: "2024-04-05",
		Tags:            []string{"tag1", "tag2"},
	}
	createResponse, err := client.CreatePost(context.Background(), post)
	assert.NoError(t, err)
	assert.NotNil(t, createResponse)
	assert.NotEqual(t, int64(0), createResponse.GetPostId())
}
