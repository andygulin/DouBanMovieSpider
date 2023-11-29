package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
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

type StoreHandle struct {
}

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

type Query interface {
	QuerySubject() (Result, error)
	QueryComment() (Result, error)
	QueryReview() (Result, error)
	QueryPhoto() (Result, error)
}

type StoreQuery struct {
	SubjectId string `json:"subject_id" bson:"subject_id"`
	PageNo    int64  `json:"page_no" bson:"page_no"`
	PageSize  int64  `json:"page_size" bson:"page_size"`
}

func (obj *StoreQuery) QuerySubject() (Result, error) {
	client, err := NewMongoDb()
	if err != nil {
		return "", err
	}

	coll := client.Database(DB).Collection(cSubject)

	var ret SubjectResponse
	err = coll.FindOne(context.TODO(), bson.M{"subject_id": obj.SubjectId}).Decode(&ret)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return Result(fmt.Sprintf("Not Found Document %s", obj.SubjectId)), err
		}
	}
	return ToInfoResult(ret)
}

func (obj *StoreQuery) QueryComment() (Result, error) {
	client, err := NewMongoDb()
	if err != nil {
		return "", err
	}

	coll := client.Database(DB).Collection(cComment)

	var ret []CommentResponse

	pageNo := obj.PageNo
	pageSize := obj.PageSize
	skip := (pageNo - 1) * pageSize
	opts := options.Find().SetSkip(skip).SetLimit(pageSize)
	cursor, err := coll.Find(context.TODO(), bson.M{"subject_id": obj.SubjectId}, opts)
	if err != nil {
		return "", err
	}
	if err = cursor.All(context.TODO(), &ret); err != nil {
		return "", err
	}

	if ret == nil {
		return "Empty Result", nil
	}
	return ToInfoResult(ret)
}

func (obj *StoreQuery) QueryReview() (Result, error) {
	client, err := NewMongoDb()
	if err != nil {
		return "", err
	}

	coll := client.Database(DB).Collection(cReview)

	var ret []ReviewResponse

	pageNo := obj.PageNo
	pageSize := obj.PageSize
	skip := (pageNo - 1) * pageSize
	opts := options.Find().SetSkip(skip).SetLimit(pageSize)
	cursor, err := coll.Find(context.TODO(), bson.M{"subject_id": obj.SubjectId}, opts)
	if err != nil {
		return "", err
	}
	if err = cursor.All(context.TODO(), &ret); err != nil {
		return "", err
	}

	if ret == nil {
		return "Empty Result", nil
	}
	return ToInfoResult(ret)
}

func (obj *StoreQuery) QueryPhoto() (Result, error) {
	client, err := NewMongoDb()
	if err != nil {
		return "", err
	}

	coll := client.Database(DB).Collection(cPhoto)

	var ret []PhotoResponse

	pageNo := obj.PageNo
	pageSize := obj.PageSize
	skip := (pageNo - 1) * pageSize
	opts := options.Find().SetSkip(skip).SetLimit(pageSize)
	cursor, err := coll.Find(context.TODO(), bson.M{"subject_id": obj.SubjectId}, opts)
	if err != nil {
		return "", err
	}
	if err = cursor.All(context.TODO(), &ret); err != nil {
		return "", err
	}

	if ret == nil {
		return "Empty Result", nil
	}
	return ToInfoResult(ret)
}
