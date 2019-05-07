package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBClient struct {
	Client *mongo.Client
}

var (
	dbCtx  context.Context
	cancel context.Context
)

func InitDB() (*DBClient, error) {
	mongoURI := os.Getenv("MONGO_URI_1")

	clientOptions := options.Client().ApplyURI(mongoURI).SetAppName("Test_GOLANG")
	dbCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoClient, err := mongo.Connect(dbCtx, clientOptions)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")

	return &DBClient{Client: mongoClient}, nil

}
