package data

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectToDB(url string) (*mongo.Client, error) {
	// create connection options
	clientOptions := options.Client().ApplyURI(url)
	clientOptions.SetAuth(options.Credential{
		Username: "artisan",
		Password: "artisan",
	})

	// connect
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Error connecting", err)
		return nil, err
	}

	Client = c

	return c, nil
}

func CloseConnection() {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	defer func() {
		if err := Client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}
