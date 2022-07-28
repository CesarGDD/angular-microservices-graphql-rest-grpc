package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cesargdd/graph-api/graph/generated"
	"cesargdd/graph-api/graph/model"
	"cesargdd/graph-api/pb/postspb"
	"cesargdd/graph-api/pb/userspb"
	"cesargdd/graph-api/services"
	"context"
)

// SignIn is the resolver for the signIn field.
func (r *mutationResolver) SignIn(ctx context.Context, input *model.NewUser) (*model.AuthResponse, error) {
	return services.SignIn(ctx, input)
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input *model.LoginInput) (*model.AuthResponse, error) {
	return services.Login(ctx, input)
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*userspb.User, error) {
	return services.DeleteUser(ctx, id)
}

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*postspb.Post, error) {
	return services.CreatePost(ctx, input)
}

// UpdatePost is the resolver for the updatePost field.
func (r *mutationResolver) UpdatePost(ctx context.Context, input model.UpdatePost) (*postspb.Post, error) {
	return services.UpdatePost(ctx, input)
}

// DeletePost is the resolver for the deletePost field.
func (r *mutationResolver) DeletePost(ctx context.Context, id string) (*postspb.Post, error) {
	return services.DeletePost(ctx, id)
}

// Posts is the resolver for the Posts field.
func (r *queryResolver) Posts(ctx context.Context) ([]*postspb.Post, error) {
	return services.Posts(ctx)
}

// Post is the resolver for the Post field.
func (r *queryResolver) Post(ctx context.Context, id string) (*postspb.Post, error) {
	return services.Post(ctx, id)
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*userspb.User, error) {
	return services.GetUser(ctx, id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
