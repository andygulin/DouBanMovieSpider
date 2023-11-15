package store

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoDb *mongo.Client

func NewMongoDb() (*mongo.Client, error) {
	var err error
	if mongoDb == nil {
		clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017/douban")
		mongoDb, err = mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			return nil, err
		}
	}
	return mongoDb, err
}
