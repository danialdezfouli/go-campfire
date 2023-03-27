package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var m *mongo.Client

func GetMongodb() *mongo.Database {
	return m.Database("campfire")
}

func CreateMongodbConnection() error {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable.")
	}

	var err error
	m, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	return nil
}

func CloseMongodbConnection() {
	if err := m.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
