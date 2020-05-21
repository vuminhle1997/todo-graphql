package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client *mongo.Client
}

func New() *Database {
	// MONGODB := os.Getenv("MONGODB")

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	clientOptions = clientOptions.SetMaxPoolSize(50)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Connection to DB")

	return &Database{
		Client: client,
	}
}
