package clients

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const connectionString = "mongodb+srv://admin:adminpass@freemongocluster.ynduakg.mongodb.net/?retryWrites=true&w=majority" // todo to config

func NewMongo(c context.Context) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(c, 10*time.Second)
	defer cancel()

	clientOpt := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(ctx, clientOpt)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongoDB due to error: %v", err)
	}

	return client, nil
}
