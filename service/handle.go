package service

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"path/filepath"
)

type Result string

type Handle interface {
	HandleSubject(response SubjectResponse) (Result, error)
	HandleComment(response []CommentResponse) (Result, error)
	HandleReview(response []ReviewResponse) (Result, error)
	HandlePhoto(response []PhotoResponse) (Result, error)
}

type InfoHandle struct {
}

func (obj *InfoHandle) HandleSubject(response SubjectResponse) (Result, error) {
	return ToInfoResult(response)
}

func (obj *InfoHandle) HandleComment(response []CommentResponse) (Result, error) {
	return ToInfoResult(response)
}

func (obj *InfoHandle) HandleReview(response []ReviewResponse) (Result, error) {
	return ToInfoResult(response)
}

func (obj *InfoHandle) HandlePhoto(response []PhotoResponse) (Result, error) {
	return ToInfoResult(response)
}

type FileHandle struct {
}

func (obj *FileHandle) HandleSubject(response SubjectResponse) (Result, error) {
	fileName := fmt.Sprintf("subject_%s.json", response.SubjectId)
	return ToFileResult(response, fileName)
}

func (obj *FileHandle) HandleComment(response []CommentResponse) (Result, error) {
	fileName := fmt.Sprintf("comment_%s.json", response[0].SubjectId)
	return ToFileResult(response, fileName)
}

func (obj *FileHandle) HandleReview(response []ReviewResponse) (Result, error) {
	fileName := fmt.Sprintf("review_%s.json", response[0].SubjectId)
	return ToFileResult(response, fileName)
}

func (obj *FileHandle) HandlePhoto(response []PhotoResponse) (Result, error) {
	fileName := fmt.Sprintf("photo_%s.json", response[0].SubjectId)
	return ToFileResult(response, fileName)
}

var mongoDb *mongo.Client

func NewMongoDb() (*mongo.Client, error) {
	var err error
	if mongoDb == nil {
		clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
		mongoDb, err = mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			return nil, err
		}
	}
	return mongoDb, err
}

type StoreHandle struct {
}

const (
	cSubject = "subject"
	cComment = "comment"
	cReview  = "review"
	cPhoto   = "photo"
)

func (obj *StoreHandle) HandleSubject(response SubjectResponse) (Result, error) {
	arr := []any{response}
	return ToStoreResult(arr, cSubject)
}

func (obj *StoreHandle) HandleComment(response []CommentResponse) (Result, error) {
	var arr []any
	for _, commentResponse := range response {
		arr = append(arr, commentResponse)
	}
	return ToStoreResult(arr, cComment)
}

func (obj *StoreHandle) HandleReview(response []ReviewResponse) (Result, error) {
	var arr []any
	for _, commentResponse := range response {
		arr = append(arr, commentResponse)
	}
	return ToStoreResult(arr, cReview)
}

func (obj *StoreHandle) HandlePhoto(response []PhotoResponse) (Result, error) {
	var arr []any
	for _, commentResponse := range response {
		arr = append(arr, commentResponse)
	}
	return ToStoreResult(arr, cPhoto)
}

func ToInfoResult(obj any) (Result, error) {
	b, err := json.MarshalIndent(obj, "", "	")
	if err != nil {
		return "", err
	}
	return Result(b), nil
}

func ToInfoResultArray(obj []any) (Result, error) {
	b, err := json.MarshalIndent(obj, "", "	")
	if err != nil {
		return "", err
	}
	return Result(b), nil
}

func ToFileResult(obj any, fileName string) (Result, error) {
	result, err := ToInfoResult(obj)
	if err != nil {
		return "", err
	}
	err = writeFile(fileName, string(result))
	if err != nil {
		return "", err
	}

	output, _ := filepath.Abs(fileName)
	return Result(output), nil
}

func ToStoreResult(obj []any, collection string) (Result, error) {
	client, err := NewMongoDb()
	if err != nil {
		return "", err
	}

	coll := client.Database(DB).Collection(collection)
	result, err := coll.InsertMany(context.Background(), obj)
	if err != nil {
		return "", err
	}

	return ToInfoResultArray(result.InsertedIDs)
}

func writeFile(fileName string, content string) error {
	file, err := os.Create(fileName)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	if err != nil {
		return err
	}
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}
