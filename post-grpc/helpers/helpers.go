package helpers

import (
	"cesargdd/posts-grpc/postspb"

	"github.com/mitchellh/mapstructure"
)

func PostsConverter(event interface{}) *postspb.Post {
	c := &postspb.Post{}
	mapstructure.Decode(event, &c)
	return c
}
