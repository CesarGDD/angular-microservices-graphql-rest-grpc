package controllers

import (
	"cesargdd/users-grpc/service"
	"cesargdd/users-grpc/userspb"
	"context"
)

type Server struct {
	userspb.UsersServiceServer
}

func (*Server) CreateUser(ctx context.Context, req *userspb.CreateUserRequest) (*userspb.CreateUserResponse, error) {
	return service.CreateUser(ctx, req)
}
func (*Server) GetUser(ctx context.Context, req *userspb.GetUserRequest) (*userspb.GetUserResponse, error) {
	return service.GetUser(ctx, req)
}
func (*Server) GetUserByUsername(ctx context.Context, req *userspb.GetUserByUsernameRequest) (*userspb.GetUserByUsernameResponse, error) {
	return service.GetUserByUsername(ctx, req)
}
func (*Server) DeleteUser(ctx context.Context, req *userspb.DeleteUserRequest) (*userspb.DeleteUserResponse, error) {
	return service.DeleteUser(ctx, req)
}
