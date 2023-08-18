package repository

import (
	"campfire/internal/database"
	"campfire/internal/domain"
	"context"
	"fmt"
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

	user.Id = fmt.Sprintf("%s", result.InsertedID)

	return err
}

func (r MessageRepositoryMongo) GetUser(ctx context.Context, userId int) (*domain.User, error) {
	return &domain.User{
		Name: "Danial",
	}, nil
}
