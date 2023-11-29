package service

import "testing"

var infoHandle Handle
var fileHandle Handle
var storeHandle Handle

var storeQuery Query

const size = 10
const pageNo = 1
const pageSize = 10

func init() {
	infoHandle = new(InfoHandle)
	fileHandle = new(FileHandle)
	storeHandle = new(StoreHandle)

	storeQuery = &StoreQuery{SubjectId: "26373447", PageNo: pageNo, PageSize: pageSize}
}

func TestInfoHandle_HandleSubject(t *testing.T) {
	result, err := infoHandle.HandleSubject(SubjectResponse{})
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(result)
}

func TestInfoHandle_HandleComment(t *testing.T) {
	var responses = make([]CommentResponse, size)
	for i := 0; i < size; i++ {
		responses = append(responses, CommentResponse{})
	}
	result, err := infoHandle.HandleComment(responses)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(result)
}

func TestInfoHandle_HandleReview(t *testing.T) {
	var responses = make([]ReviewResponse, size)
	for i := 0; i < size; i++ {
		responses = append(responses, ReviewResponse{})
	}
	result, err := infoHandle.HandleReview(responses)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(result)
}

func TestInfoHandle_HandlePhoto(t *testing.T) {
	var responses = make([]PhotoResponse, size)
	for i := 0; i < size; i++ {
		responses = append(responses, PhotoResponse{})
	}
	result, err := infoHandle.HandlePhoto(responses)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(result)
}

func TestFileHandle_HandleSubject(t *testing.T) {
	result, err := fileHandle.HandleSubject(SubjectResponse{})
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(result)
}

func TestFileHandle_HandleComment(t *testing.T) {
	var responses = make([]CommentResponse, size)
	for i := 0; i < size; i++ {
		responses = append(responses, CommentResponse{})
	}
	result, err := fileHandle.HandleComment(responses)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(result)
}

func TestFileHandle_HandleReview(t *testing.T) {
	var responses = make([]ReviewResponse, size)
	for i := 0; i < size; i++ {
		responses = append(responses, ReviewResponse{})
	}
	result, err := fileHandle.HandleReview(responses)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(result)
}

func TestFileHandle_HandlePhoto(t *testing.T) {
	var responses = make([]PhotoResponse, size)
	for i := 0; i < size; i++ {
		responses = append(responses, PhotoResponse{})
	}
	result, err := fileHandle.HandlePhoto(responses)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(result)
}

func TestStoreHandle_HandleSubject(t *testing.T) {
	result, err := storeHandle.HandleSubject(SubjectResponse{})
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(result)
}

func TestStoreHandle_HandleComment(t *testing.T) {
	var responses = make([]CommentResponse, size)
	for i := 0; i < size; i++ {
		responses[i] = CommentResponse{}
	}
	result, err := storeHandle.HandleComment(responses)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(result)
}

func TestStoreHandle_HandleReview(t *testing.T) {
	var responses = make([]ReviewResponse, size)
	for i := 0; i < size; i++ {
		responses[i] = ReviewResponse{}
	}
	result, err := storeHandle.HandleReview(responses)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(result)
}

func TestStoreHandle_HandlePhoto(t *testing.T) {
	var responses = make([]PhotoResponse, size)
	for i := 0; i < size; i++ {
		responses[i] = PhotoResponse{}
	}
	result, err := storeHandle.HandlePhoto(responses)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(result)
}

func TestStoreQuery_QuerySubject(t *testing.T) {
	result, err := storeQuery.QuerySubject()
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(result)
}

func TestStoreQuery_QueryComment(t *testing.T) {
	result, err := storeQuery.QueryComment()
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(result)
}

func TestStoreQuery_QueryReview(t *testing.T) {
	result, err := storeQuery.QueryReview()
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(result)
}

func TestStoreQuery_QueryPhoto(t *testing.T) {
	result, err := storeQuery.QueryPhoto()
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(result)
}
