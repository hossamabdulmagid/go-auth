package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://sawsawdb:13102001a@cluster0.qhjct.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)

	}

	defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	databases, err := client.ListDatabaseNames(ctx, bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(databases)
}
