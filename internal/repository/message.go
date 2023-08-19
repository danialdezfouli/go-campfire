package repository

import (
	"campfire/internal/database"
	"campfire/internal/domain"
	"context"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageRepositoryMongo struct {
}

func NewMessageRepositoryMongodb() MessageRepositoryMongo {
	return MessageRepositoryMongo{}
}

func (r MessageRepositoryMongo) CreateUser(ctx context.Context, user *domain.User) error {
	result, err := database.GetMongodb().Collection(domain.UserTableName).InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	insertedID := result.InsertedID.(primitive.ObjectID).Hex()

	userIdInt, err := strconv.Atoi(insertedID)
	if err != nil {
		return err
	}

	userId := domain.UserId(userIdInt)

	user.Id = userId
	return nil
}

func (r MessageRepositoryMongo) GetUser(ctx context.Context, userId int) (*domain.User, error) {
	return &domain.User{
		Name: "Danial",
	}, nil
}
