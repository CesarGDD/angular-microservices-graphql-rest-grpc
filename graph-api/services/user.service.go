package services

import (
	"cesargdd/graph-api/graph/model"
	"cesargdd/graph-api/pb/userspb"
	"cesargdd/graph-api/servers"
	"cesargdd/graph-api/tools"
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"google.golang.org/grpc/status"
)

var userSvc = servers.UsersSvc()

func CreateUser(ctx context.Context, input model.NewUser) (*userspb.User, error) {
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
	return res.User, nil
}

func GetUser(ctx context.Context, id string) (*userspb.User, error) {
	res, err := userSvc.GetUser(ctx, &userspb.GetUserRequest{Id: id})
	if err != nil {
		fmt.Println(err)
		e, ok := status.FromError(err)
		if ok {
			graphql.AddError(ctx, gqlerror.Errorf(e.Message()))
		} else {
			graphql.AddError(ctx, gqlerror.Errorf("No grpc error", err))
		}
	}
	return res.User, nil
}

func GetUserByUsername(ctx context.Context, username string) (*userspb.User, error) {
	res, err := userSvc.GetUserByUsername(ctx, &userspb.GetUserByUsernameRequest{
		Username: username,
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
	return res.User, nil
}

func DeleteUser(ctx context.Context, id string) (*userspb.User, error) {
	res, err := userSvc.DeleteUser(ctx, &userspb.DeleteUserRequest{Id: id})
	if err != nil {
		fmt.Println(err)
		e, ok := status.FromError(err)
		if ok {
			graphql.AddError(ctx, gqlerror.Errorf(e.Message()))
		} else {
			graphql.AddError(ctx, gqlerror.Errorf("No grpc error", err))
		}
	}
	return res.User, nil
}
