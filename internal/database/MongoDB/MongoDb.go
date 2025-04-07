package MongoDB

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
)

var (
	mongoClient *mongo.Client
	once        sync.Once
)

// GetMongoClient возвращает singleton MongoDB-клиент
func GetMongoClient() *mongo.Client {
	once.Do(func() {
		ctx := context.TODO()
		client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:55555"))
		if err != nil {
			log.Fatalf("Ошибка подключения к MongoDB: %v", err)
		}
		mongoClient = client
	})
	return mongoClient
}
