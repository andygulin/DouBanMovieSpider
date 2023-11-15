package spider

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSpiderRequest_SpiderContent(t *testing.T) {
	obj := Request{SubjectId: "1292226"}
	response, err := obj.SpiderSubject()
	if err != nil {
		fmt.Println(err)
	}
	content, err := json.MarshalIndent(response, "", "	")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(content))
}
