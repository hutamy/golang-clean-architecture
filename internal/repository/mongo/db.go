package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBManager struct {
	client   *mongo.Client
	database *mongo.Database
}

// NewDBManager initializes a new MongoDB client and connects to the specified database
func NewDBManager(uri, databaseName string) *DBManager {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB")

	return &DBManager{
		client:   client,
		database: client.Database(databaseName),
	}
}

// GetCollection returns a MongoDB collection
func (dbm *DBManager) GetCollection(name string) *mongo.Collection {
	return dbm.database.Collection(name)
}

// Close closes the MongoDB connection
func (dbm *DBManager) Close() {
	if err := dbm.client.Disconnect(context.Background()); err != nil {
		log.Printf("Error disconnecting MongoDB: %v", err)
	}
}
