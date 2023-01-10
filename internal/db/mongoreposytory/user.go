package mongoreposytory

import (
	"chat-app-golang/internal/clients"
	"chat-app-golang/internal/db"
	"chat-app-golang/internal/domain"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type User struct {
}

func (u User) Get(ctx context.Context, user domain.User) (*domain.User, error) {
	c, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	mongoClient, err := clients.NewMongo(ctx)
	if err != nil {
		return nil, err // todo Errors
	}

	var result domain.User
	if err = mongoClient.Database("chat-app-golang").Collection("user").FindOne(c, &user).Decode(&result); errors.Is(err, mongo.ErrNoDocuments) {
		return nil, db.NoResultErr
	}

	return &result, err
}

func (u User) Insert(ctx context.Context, user domain.User) error {
	c, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	mongo, err := clients.NewMongo(c)
	if err != nil {
		return err
	}

	result, err := mongo.Database("chat-app-golang").Collection("user").InsertOne(c, user)
	if err != nil {
		return err
	}

	fmt.Println(result)
	return nil
}

func (u User) Update(ctx context.Context, user domain.User) {
	//TODO implement me
	panic("implement me")
}

func (u User) Archive(ctx context.Context, user domain.User) {
	//TODO implement me
	panic("implement me")
}

func (u User) Delete(ctx context.Context, user domain.User) {
	//TODO implement me
	panic("implement me")
}

func (u User) GetUsersByCredentials(ctx context.Context, user domain.User) (*[]domain.User, error) {
	c, cancel := context.WithTimeout(ctx, 500000*time.Second)
	defer cancel()

	mongoClient, err := clients.NewMongo(ctx)
	if err != nil {
		return nil, err // todo Errors
	}

	filter := bson.M{
		"key1": 1,
		"$or": []interface{}{
			bson.M{"email": *user.Email},
			bson.M{"username": *user.Username},
		},
	}

	var result []domain.User
	res, _ := mongoClient.Database("chat-app-golang").Collection("user").Find(c, filter)
	if err = res.Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
