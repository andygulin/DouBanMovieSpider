package service

import (
	"encoding/json"
	"testing"
)

var obj Spider

func init() {
	obj = &Request{SubjectId: "26373447"}
}

func TestRequest_SpiderSubject(t *testing.T) {
	response, err := obj.SpiderSubject()
	if err != nil {
		t.Log(err)
	}
	content, err := json.MarshalIndent(response, "", "	")
	if err != nil {
		t.Log(err)
	}
	t.Log(string(content))
}

func TestRequest_SpiderComment(t *testing.T) {
	response, err := obj.SpiderComment()
	if err != nil {
		t.Log(err)
	}

	limit := 10
	if len(response) > limit {
		response = response[0:limit]
	}

	content, err := json.MarshalIndent(response, "", "	")
	if err != nil {
		t.Log(err)
	}
	t.Log(string(content))
}

func TestRequest_SpiderReview(t *testing.T) {
	response, err := obj.SpiderReview()
	if err != nil {
		t.Log(err)
	}

	limit := 10
	if len(response) > limit {
		response = response[0:limit]
	}

	content, err := json.MarshalIndent(response, "", "	")
	if err != nil {
		t.Log(err)
	}
	t.Log(string(content))
}

func TestRequest_SpiderPhoto(t *testing.T) {
	response, err := obj.SpiderPhoto()
	if err != nil {
		t.Log(err)
	}

	limit := 10
	if len(response) > limit {
		response = response[0:limit]
	}

	content, err := json.MarshalIndent(response, "", "	")
	if err != nil {
		t.Log(err)
	}
	t.Log(string(content))
}
