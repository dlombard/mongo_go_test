package main

import (
	"context"
	"fmt"
	"time"

	"github.com/dlombard/test_go/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var coll *mongo.Collection

func poll() {

	fmt.Println("Polling")
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		res := coll.FindOne(ctx, bson.M{"type": "x"})

		fmt.Println(res.DecodeBytes())
		time.Sleep(5 * time.Second)

	}
}
func main() {
	dbClient, err := db.InitDB()

	if err != nil {
		fmt.Println(err.Error())
	}

	coll = dbClient.Client.Database("test").Collection("coll")

	if err != nil {
		fmt.Println(err.Error())
	}

	poll()
}
