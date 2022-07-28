package servers

import (
	"cesargdd/graph-api/pb/postspb"
	"cesargdd/graph-api/pb/userspb"
	"log"
	"os"

	"google.golang.org/grpc"
)

func PostSvc() postspb.PostsServiceClient {
	opts := grpc.WithInsecure()
	ccu, err := grpc.Dial(os.Getenv("POSTS_URL"), opts)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	postSvc := postspb.NewPostsServiceClient(ccu)
	return postSvc
}
func UsersSvc() userspb.UsersServiceClient {
	opts := grpc.WithInsecure()
	ccu, err := grpc.Dial(os.Getenv("USERS_URL"), opts)
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	usersSvc := userspb.NewUsersServiceClient(ccu)
	return usersSvc
}
