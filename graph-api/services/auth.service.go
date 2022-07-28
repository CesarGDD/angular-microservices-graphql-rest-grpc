package services

import (
	"cesargdd/graph-api/graph/model"
	"cesargdd/graph-api/pb/userspb"
	"cesargdd/graph-api/tools"
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"google.golang.org/grpc/status"
)

func SignIn(ctx context.Context, input *model.NewUser) (*model.AuthResponse, error) {
	input.Password = tools.HashPassword(input.Password)
	res, err := userSvc.CreateUser(ctx, &userspb.CreateUserRequest{
		Username: input.Username,
		Password: input.Password,
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
	token, err := JwtGenerate(ctx, res.User.Id)
	if err != nil {
		return nil, err
	}
	return &model.AuthResponse{
		Token: token,
		User:  res.User,
	}, nil
}

// Login is the resolver for the login field.
func Login(ctx context.Context, input *model.LoginInput) (*model.AuthResponse, error) {
	user, err := GetUserByUsername(ctx, input.Username)
	if err != nil {
		fmt.Println(err)
		e, ok := status.FromError(err)
		if ok {
			graphql.AddError(ctx, gqlerror.Errorf(e.Message()))
		} else {
			graphql.AddError(ctx, gqlerror.Errorf("No grpc error", err))
		}
	}
	if err := tools.ComparePassword(user.Password, input.Password); err != nil {
		return nil, err
	}
	token, err := JwtGenerate(ctx, user.Id)
	if err != nil {
		return nil, err
	}
	return &model.AuthResponse{
		Token: token,
		User:  user,
	}, nil
}
