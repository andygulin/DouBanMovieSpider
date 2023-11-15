package store

import (
	"context"
	"fmt"
	"testing"
)

func TestNewMongoDb(t *testing.T) {
	client, err := NewMongoDb()
	if err != nil {
		fmt.Println(err.Error())
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
