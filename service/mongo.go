package service

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const (
	cSubject = "subject"
	cComment = "comment"
	cReview  = "review"
	cPhoto   = "photo"
)

var (
	mongoDb *mongo.Client
	DB      string
	once    sync.Once
)

func NewMongoDb() (*mongo.Client, error) {
	var err error
	once.Do(func() {
		viper.SetConfigFile("../conf/conf.yaml")
		if err := viper.ReadInConfig(); err != nil {
			return
		}

		host := viper.GetString("mongo.host")
		port := viper.GetInt("mongo.port")
		DB = viper.GetString("mongo.db")
		uri := fmt.Sprintf("mongodb://%s:%d", host, port)

		clientOptions := options.Client().ApplyURI(uri)
		mongoDb, err = mongo.Connect(clientOptions)
	})
	return mongoDb, err
}
