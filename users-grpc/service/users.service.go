package service

import (
	"cesargdd/users-grpc/db"
	"cesargdd/users-grpc/userspb"
	"context"
	"fmt"

	"github.com/couchbase/gocb/v2"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var col, cluster = db.Con()

func CreateUser(ctx context.Context, req *userspb.CreateUserRequest) (*userspb.CreateUserResponse, error) {
	id := uuid.New().String()
	res, err := cluster.Query("SELECT username FROM `users`.users.users WHERE username = $1", &gocb.QueryOptions{
		PositionalParameters: []interface{}{req.Username},
	})
	if err != nil {
		return nil, err
	}
	var username interface{}
	user := &userspb.User{
		Id:       id,
		Username: req.Username,
		Password: req.Password,
	}
	res.One(&username)
	if username == nil {
		_, err = col.Insert(id, user, nil)
		if err != nil {
			fmt.Println(err)
			return nil, status.Errorf(codes.Unknown, "Error creating user")
		}
	} else {
		return nil, status.Errorf(codes.Unknown, "Username already exits")
	}
	return &userspb.CreateUserResponse{
		User: user,
	}, nil
}

func GetUserByUsername(ctx context.Context, req *userspb.GetUserByUsernameRequest) (*userspb.GetUserByUsernameResponse, error) {
	res, err := cluster.Query("SELECT id FROM `users`.users.users WHERE username = $1", &gocb.QueryOptions{
		PositionalParameters: []interface{}{req.Username},
	})
	if err != nil {
		fmt.Println(err)
		return nil, status.Errorf(codes.NotFound, "User not found")
	}
	data := make(map[string]string)
	user := &userspb.User{}
	res.One(&data)
	id := data["id"]
	if data["id"] == "" {
		return nil, status.Errorf(codes.NotFound, "User not found")
	}
	row, err := col.Get(id, nil)
	if err != nil {
		fmt.Println(err)
		return nil, status.Errorf(codes.NotFound, "User not found")
	}
	row.Content(&user)
	return &userspb.GetUserByUsernameResponse{
		User: user,
	}, nil
}

func GetUser(ctx context.Context, req *userspb.GetUserRequest) (*userspb.GetUserResponse, error) {
	res, err := col.Get(req.Id, nil)
	if err != nil {
		fmt.Println(err)
		return nil, status.Errorf(codes.NotFound, "Error getting user")
	}
	user := &userspb.User{}
	res.Content(&user)
	return &userspb.GetUserResponse{
		User: user,
	}, nil
}
func DeleteUser(ctx context.Context, req *userspb.DeleteUserRequest) (*userspb.DeleteUserResponse, error) {
	res, err := col.Get(req.Id, nil)
	if err != nil {
		fmt.Println(err)
		return nil, status.Errorf(codes.NotFound, "Error getting user")
	}
	user := &userspb.User{}
	res.Content(&user)
	_, err = col.Remove(user.Id, nil)
	if err != nil {
		fmt.Println(err)
		return nil, status.Errorf(codes.Unknown, "Error deleting user")
	}
	return &userspb.DeleteUserResponse{
		User: user,
	}, nil
}
