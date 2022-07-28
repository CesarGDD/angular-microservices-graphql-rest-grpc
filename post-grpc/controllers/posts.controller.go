package controllers

import (
	"cesargdd/posts-grpc/postspb"
	"cesargdd/posts-grpc/service"
	"context"
)

type Server struct {
	postspb.PostsServiceServer
}

func (*Server) CreatePost(ctx context.Context, req *postspb.CreatePostRequest) (*postspb.CreatePostResponse, error) {
	return service.CreatePost(ctx, req)
}
func (*Server) GetPost(ctx context.Context, req *postspb.GetPostRequest) (*postspb.GetPostResponse, error) {
	return service.GetPost(ctx, req)
}
func (*Server) UpdatePost(ctx context.Context, req *postspb.UpdatePostRequest) (*postspb.UpdatePostResponse, error) {
	return service.UpdatePost(ctx, req)
}
func (*Server) DeletePost(ctx context.Context, req *postspb.DeletePostRequest) (*postspb.DeletePostResponse, error) {
	return service.DeletePost(ctx, req)
}
func (*Server) GetPosts(ctx context.Context, req *postspb.GetPostsRequest) (*postspb.GetPostsResponse, error) {
	return service.GetPosts(ctx, req)
}
