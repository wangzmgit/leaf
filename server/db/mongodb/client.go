package mongodb

import "go.mongodb.org/mongo-driver/mongo"

type MongoClient struct {
	db *mongo.Database
}

func (m *MongoClient) Like() *mongo.Collection {
	return m.db.Collection("like")
}

func (m *MongoClient) Collect() *mongo.Collection {
	return m.db.Collection("collect")
}

func (m *MongoClient) Comment() *mongo.Collection {
	return m.db.Collection("comment")
}
