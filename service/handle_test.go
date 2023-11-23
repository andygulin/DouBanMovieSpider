package service

import "testing"

var infoHandle Handle
var fileHandle Handle
var storeHandle Handle

const size = 10

func init() {
	infoHandle = new(InfoHandle)
	fileHandle = new(FileHandle)
	storeHandle = new(StoreHandle)
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
