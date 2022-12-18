package mongodb

import "go.mongodb.org/mongo-driver/mongo"

type MongoClient struct {
	db *mongo.Database
}

func (m *MongoClient) Like() *mongo.Collection {
	return m.db.Collection("like")
}
