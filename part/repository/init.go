package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type dbs struct {
	MongoDB *mongo.Client
	Context *context.Context
}

var DBS dbs

func Init() {
	mongoConnection()
}

func mongoConnection() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongodb:27017/parts"))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	DBS.MongoDB = client
	DBS.Context = &ctx
	fmt.Println("Successfully connected and pinged.")
}
