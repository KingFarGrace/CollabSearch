package conf

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestGetEngine(t *testing.T) {
	InitPersistenceLayer()
	engine := GetMySQLEngine()
	exist, err := engine.IsTableExist("user")
	if err != nil {
		t.Error("ERROR: ", err)
	}
	fmt.Println("Is table Exist? ", exist)
}

func TestGetRedisClient(t *testing.T) {
	InitCachingLayer()
	ctx := context.Background()
	client := GetRedisClient()
	client.Set(ctx, "name", "King", 5*time.Second)
	val, err := client.Get(ctx, "name").Result()
	if err != nil {
		t.Error("ERROR: ", err)
	}
	fmt.Printf("key: %s value: %v", "name", val)
}
