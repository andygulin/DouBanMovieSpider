package service

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	cSubject = "subject"
	cComment = "comment"
	cReview  = "review"
	cPhoto   = "photo"
)

var mongoDb *mongo.Client
var DB string

func NewMongoDb() (*mongo.Client, error) {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	path = path[:index]
	viper.SetConfigFile(path + "/conf/conf.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	host := viper.GetString("mongo.host")
	port := viper.GetInt("mongo.port")
	DB = viper.GetString("mongo.db")
	uri := fmt.Sprintf("mongodb://%s:%d", host, port)

	if mongoDb == nil {
		clientOptions := options.Client().ApplyURI(uri)
		mongoDb, err = mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			return nil, err
		}
	}
	return mongoDb, err
}
