package main

import (
	pb "blogApp/protos"
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.BlogServiceServer
	posts map[int64]*pb.Post
	//pb.mustEmbedUnimplementedBlogServiceServer
}

// Creates new Blog
func (s *server) CreatePost(ctx context.Context, in *pb.Post) (*pb.Post, error) {
	postID := rand.Intn(101)
	in.PostId = int64(postID)
	s.posts[int64(postID)] = in
	return in, nil
}

// For Reading Current Blog using PostID
func (s *server) ReadPost(ctx context.Context, in *pb.PostID) (*pb.Post, error) {
	post, ok := s.posts[in.GetPostId()]
	if !ok {
		return nil, errors.New("post not found")
	}
	return post, nil
}

// Updates already existing Post's Details if PostID in request exists
func (s *server) UpdatePost(ctx context.Context, in *pb.Post) (*pb.Post, error) {
	_, ok := s.posts[in.GetPostId()]
	if !ok {
		return nil, errors.New("post not found")
	}
	s.posts[in.GetPostId()] = in
	return in, nil
}

// For Deleting Current Blog using PostID
func (s *server) DeletePost(ctx context.Context, in *pb.PostID) (*pb.DeleteResponse, error) {
	_, ok := s.posts[in.GetPostId()]
	if !ok {
		return &pb.DeleteResponse{Success: false}, errors.New("post not found")
	}
	delete(s.posts, in.GetPostId())
	return &pb.DeleteResponse{Success: true}, nil
}

const (
	listenerPort = ":9090"
)

// start listening on a port
// Hosts a GRPC API
func main() {
	lis, err := net.Listen("tcp", listenerPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterBlogServiceServer(s, &server{posts: make(map[int64]*pb.Post)})

	fmt.Println("Server listening on port 9090")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
