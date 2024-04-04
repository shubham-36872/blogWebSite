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

/*type PostDetails struct {
	PostId          int64    `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	Title           string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content         string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Author          string   `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
	PublicationDate string   `protobuf:"bytes,5,opt,name=publication_date,json=publicationDate,proto3" json:"publication_date,omitempty"`
	Tags            []string `protobuf:"bytes,6,rep,name=tags,proto3" json:"tags,omitempty"`
}*/

func (s *server) CreatePost(ctx context.Context, in *pb.Post) (*pb.Post, error) {
	postID := rand.Intn(101)
	in.PostId = int64(postID)
	s.posts[int64(postID)] = in
	return in, nil
}

func (s *server) ReadPost(ctx context.Context, in *pb.PostID) (*pb.Post, error) {
	post, ok := s.posts[in.GetPostId()]
	if !ok {
		return nil, errors.New("post not found")
	}
	return post, nil
}

func (s *server) UpdatePost(ctx context.Context, in *pb.Post) (*pb.Post, error) {
	_, ok := s.posts[in.GetPostId()]
	if !ok {
		return nil, errors.New("post not found")
	}
	s.posts[in.GetPostId()] = in
	return in, nil
}

func (s *server) DeletePost(ctx context.Context, in *pb.PostID) (*pb.DeleteResponse, error) {
	_, ok := s.posts[in.GetPostId()]
	if !ok {
		return &pb.DeleteResponse{Success: false}, errors.New("post not found")
	}
	delete(s.posts, in.GetPostId())
	return &pb.DeleteResponse{Success: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9090")
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
