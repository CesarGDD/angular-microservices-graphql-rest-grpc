package services

import (
	"cesargdd/graph-api/graph/model"
	"cesargdd/graph-api/pb/postspb"
	"cesargdd/graph-api/servers"
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"google.golang.org/grpc/status"
)

var svc = servers.PostSvc()

func CreatePost(ctx context.Context, input model.NewPost) (*postspb.Post, error) {
	fmt.Println("Create Post graph")
	res, err := svc.CreatePost(ctx, &postspb.CreatePostRequest{
		User: input.User,
		Content: &postspb.Content{
			Tittle: input.Content.Tittle,
			Body:   input.Content.Body,
			Image:  input.Content.Image,
		},
	})
	if err != nil {
		fmt.Println(err)
		e, ok := status.FromError(err)
		if ok {
			graphql.AddError(ctx, gqlerror.Errorf(e.Message()))
		} else {
			graphql.AddError(ctx, gqlerror.Errorf("No grpc error", err))
		}
	}
	return res.Post, nil
}

func UpdatePost(ctx context.Context, input model.UpdatePost) (*postspb.Post, error) {
	fmt.Println("Update Post graph")
	res, err := svc.UpdatePost(ctx, &postspb.UpdatePostRequest{
		Id: input.ID,
		Content: &postspb.Content{
			Tittle: input.Content.Tittle,
			Body:   input.Content.Body,
			Image:  input.Content.Image,
		},
	})
	if err != nil {
		fmt.Println(err)
		e, ok := status.FromError(err)
		if ok {
			graphql.AddError(ctx, gqlerror.Errorf(e.Message()))
		} else {
			graphql.AddError(ctx, gqlerror.Errorf("No grpc error", err))
		}
	}
	return res.Post, nil
}

func DeletePost(ctx context.Context, id string) (*postspb.Post, error) {
	fmt.Println("Delete Post graph")
	res, err := svc.DeletePost(ctx, &postspb.DeletePostRequest{Id: id})
	if err != nil {
		fmt.Println(err)
		e, ok := status.FromError(err)
		if ok {
			graphql.AddError(ctx, gqlerror.Errorf(e.Message()))
		} else {
			graphql.AddError(ctx, gqlerror.Errorf("No grpc error", err))
		}
	}
	return res.Post, nil
}

func Posts(ctx context.Context) ([]*postspb.Post, error) {
	fmt.Println("Posts graph")
	res, err := svc.GetPosts(ctx, &postspb.GetPostsRequest{})
	if err != nil {
		fmt.Println(err)
		e, ok := status.FromError(err)
		if ok {
			graphql.AddError(ctx, gqlerror.Errorf(e.Message()))
		} else {
			graphql.AddError(ctx, gqlerror.Errorf("No grpc error", err))
		}
	}
	return res.Post, nil
}

func Post(ctx context.Context, id string) (*postspb.Post, error) {
	fmt.Println("Get Post graph")
	res, err := svc.GetPost(ctx, &postspb.GetPostRequest{Id: id})
	if err != nil {
		fmt.Println(err)
		e, ok := status.FromError(err)
		if ok {
			graphql.AddError(ctx, gqlerror.Errorf(e.Message()))
		} else {
			graphql.AddError(ctx, gqlerror.Errorf("No grpc error", err))
		}
	}
	return res.Post, nil
}
