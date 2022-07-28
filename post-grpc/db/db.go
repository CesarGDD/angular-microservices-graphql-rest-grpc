package db

import (
	"log"
	"os"
	"time"

	"github.com/couchbase/gocb/v2"
)

func Con() (*gocb.Collection, *gocb.Cluster) {
	// Update this to your cluster details
	bucketName := "posts"
	username := "admin"
	password := "123456"

	// gocb.SetLogger(gocb.VerboseStdioLogger())

	// Initialize the Connection
	cluster, err := gocb.Connect("couchbases://"+os.Getenv("COUCHBASE"), gocb.ClusterOptions{
		Authenticator: gocb.PasswordAuthenticator{
			Username: username,
			Password: password,
		},
		SecurityConfig: gocb.SecurityConfig{
			// WARNING: Do not set this to true in production, only use this for testing!
			TLSSkipVerify: true,
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	bucket := cluster.Bucket(bucketName)

	err = bucket.WaitUntilReady(5*time.Second, nil)
	if err != nil {
		log.Fatal(err)
	}
	// bucket.Collections().CreateScope("posts", nil)
	bucket.Collections().CreateCollection(gocb.CollectionSpec{
		Name:      "posts",
		ScopeName: "posts",
		MaxExpiry: 0,
	}, nil)
	col := bucket.Scope("posts").Collection("posts")
	return col, cluster
}
