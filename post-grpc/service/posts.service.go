package service

import (
	"cesargdd/posts-grpc/db"
	"cesargdd/posts-grpc/helpers"
	"cesargdd/posts-grpc/postspb"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var col, cluster = db.Con()

func CreatePost(ctx context.Context, req *postspb.CreatePostRequest) (*postspb.CreatePostResponse, error) {
	fmt.Println("Create POst gRPC")
	id := uuid.New().String()
	post := &postspb.Post{
		Id:   id,
		User: req.User,
		Content: &postspb.Content{
			Tittle: req.Content.Tittle,
			Body:   req.Content.Body,
			Image:  req.Content.Image,
		},
		Date: time.Now().String(),
	}
	_, err := col.Insert(id, post, nil)
	if err != nil {
		fmt.Println(err)
		return nil, status.Errorf(codes.Unknown, "Error creating post")
	}
	return &postspb.CreatePostResponse{
		Post: post,
	}, nil
}
func GetPost(ctx context.Context, req *postspb.GetPostRequest) (*postspb.GetPostResponse, error) {
	res, err := col.Get(req.Id, nil)
	if err != nil {
		fmt.Println(err)
		return nil, status.Errorf(codes.NotFound, "Post not found")
	}
	post := &postspb.Post{}
	res.Content(&post)
	return &postspb.GetPostResponse{
		Post: post,
	}, nil
}
func UpdatePost(ctx context.Context, req *postspb.UpdatePostRequest) (*postspb.UpdatePostResponse, error) {
	res, err := col.Get(req.Id, nil)
	if err != nil {
		fmt.Println(err)
		return nil, status.Errorf(codes.NotFound, "Post not found")
	}
	post := &postspb.Post{}
	res.Content(&post)
	post.Content.Body = req.Content.Body
	post.Content.Image = req.Content.Image
	post.Content.Tittle = req.Content.Tittle
	_, err = col.Upsert(post.Id, post, nil)
	if err != nil {
		fmt.Println(err)
		return nil, status.Errorf(codes.Unavailable, "Error updating post")
	}
	return &postspb.UpdatePostResponse{
		Post: post,
	}, nil
}
func DeletePost(ctx context.Context, req *postspb.DeletePostRequest) (*postspb.DeletePostResponse, error) {
	res, err := col.Get(req.Id, nil)
	if err != nil {
		fmt.Println(err)
		return nil, status.Errorf(codes.NotFound, "Post not found")
	}
	post := &postspb.Post{}
	res.Content(&post)
	_, err = col.Remove(post.Id, nil)
	if err != nil {
		fmt.Println(err)
		return nil, status.Errorf(codes.Unknown, "Could not delete post")
	}
	return &postspb.DeletePostResponse{
		Post: post,
	}, nil
}
func GetPosts(ctx context.Context, req *postspb.GetPostsRequest) (*postspb.GetPostsResponse, error) {
	fmt.Println("Posts")
	res, err := cluster.Query("SELECT x.* FROM `posts`.posts.posts x", nil)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Error getting posts found")
	}
	var posts []*postspb.Post
	var post interface{}
	for res.Next() {
		res.Row(&post)
		newPost := helpers.PostsConverter(post)
		posts = append(posts, newPost)
	}
	return &postspb.GetPostsResponse{
		Post: posts,
	}, nil
}
