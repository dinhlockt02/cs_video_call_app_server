package authstore

import "go.mongodb.org/mongo-driver/mongo"

type MongoStore struct {
	database *mongo.Database
}

func NewMongoStore(database *mongo.Database) *MongoStore {
	return &MongoStore{database: database}
}
