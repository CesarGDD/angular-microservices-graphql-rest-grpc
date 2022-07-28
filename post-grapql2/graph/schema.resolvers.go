package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cesargdd/graphql-posts2/db"
	"cesargdd/graphql-posts2/graph/generated"
	"cesargdd/graphql-posts2/graph/model"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
)

var col, cluster = db.Con()

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	fmt.Println("Create Post")
	id := uuid.New().String()
	post := &model.Post{
		ID:   id,
		User: input.User,
		Content: &model.Content{
			Title: input.Content.Title,
			Body:  input.Content.Body,
		},
		Date: time.Now().String(),
	}
	_, err := col.Insert(id, post, nil)
	if err != nil {
		return nil, err
	}
	return post, nil
}

// UpdatePost is the resolver for the updatePost field.
func (r *mutationResolver) UpdatePost(ctx context.Context, input model.UpdatePost) (*model.Post, error) {
	fmt.Println("Update")
	res, err := col.Get(input.ID, nil)
	if err != nil {
		return nil, err
	}
	post := &model.Post{}
	res.Content(post)
	post.Content.Title = input.Content.Title
	post.Content.Body = input.Content.Body
	_, err = col.Upsert(input.ID, post, nil)
	if err != nil {
		return nil, err
	}

	return post, nil
}

// DeletePost is the resolver for the deletePost field.
func (r *mutationResolver) DeletePost(ctx context.Context, id string) (*model.Post, error) {
	fmt.Println("Delete Post")
	res, err := col.Get(id, nil)
	if err != nil {
		return nil, err
	}
	post := &model.Post{}
	res.Content(post)
	_, err = col.Remove(post.ID, nil)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func postsConverter(event interface{}) *model.Post {
	c := &model.Post{}
	mapstructure.Decode(event, &c)
	return c
}

// Posts is the resolver for the Posts field.
func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	fmt.Println("Posts")
	res, err := cluster.Query("SELECT x.* FROM `posts`.posts.posts x", nil)
	if err != nil {
		return nil, err
	}
	var posts []*model.Post
	var post interface{}
	for res.Next() {
		res.Row(&post)
		newPost := postsConverter(post)
		posts = append(posts, newPost)
	}
	return posts, nil
}

// Post is the resolver for the Post field.
func (r *queryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
	fmt.Println("Get Post")
	res, err := col.Get(id, nil)
	if err != nil {
		return nil, err
	}
	post := &model.Post{}
	res.Content(post)
	return post, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
