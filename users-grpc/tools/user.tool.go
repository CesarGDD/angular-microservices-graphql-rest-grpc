package tools

import (
	"cesargdd/users-grpc/userspb"

	"github.com/mitchellh/mapstructure"
)

func UserConverter(event interface{}) *userspb.User {
	c := &userspb.User{}
	mapstructure.Decode(event, &c)
	return c
}
