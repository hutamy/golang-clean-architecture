package cache

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

type DBManager struct {
	client *redis.Client
}

// NewDBManager initializes a Redis client
func NewDBManager(address, password string, db int) *DBManager {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})

	// Test connection
	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Connected to Redis")

	return &DBManager{client: client}
}

// GetClient returns the Redis client
func (dbm *DBManager) GetClient() *redis.Client {
	return dbm.client
}

// Close closes the Redis connection
func (dbm *DBManager) Close() {
	if err := dbm.client.Close(); err != nil {
		log.Printf("Error closing Redis connection: %v", err)
	} else {
		log.Println("Redis connection closed")
	}
}
