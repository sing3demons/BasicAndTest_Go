package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://sing3demons:KP05xxvi40@cluster0-ghtvv.mongodb.net/test?authSource=admin&replicaSet=Cluster0-shard-0&readPreference=primary&appname=MongoDB%20Compass&ssl=true"))

	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	dbs, err := client.ListDatabaseNames(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	for _, dbName := range dbs {
		fmt.Print(dbName)
	}
}
