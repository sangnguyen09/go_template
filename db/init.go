package db

import (
	"context"
	"fmt"
	"time"

	"github.com/letanthang/my_framework/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	dbName = "go3008"
)

var Client *mongo.Client

func init() {
	uri := config.Config.Mongo.URI
	if config.Config.Mongo.URI == "" {
		uri = fmt.Sprintf("mongodb://%s:%s@%s:27017", config.Config.Mongo.User, config.Config.Mongo.Password, config.Config.Mongo.Host)
	}
	fmt.Println("uri1", uri)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	Client = client
}
