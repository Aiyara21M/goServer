package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoConnectorInterface
type MongoConnectorInterface interface {
	GetClient() *mongo.Client
	Close() error
}

type MongoDBConnector struct {
	client *mongo.Client
}

func NewMongoDBConnector(uri string) (MongoConnectorInterface, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("❌ failed to connect to MongoDB: %w", err)
	}

	fmt.Println("✅ Connected to MongoDB:", uri)
	return &MongoDBConnector{client: client}, nil
}

func (m *MongoDBConnector) GetClient() *mongo.Client {
	return m.client
}

func (m *MongoDBConnector) Close() error {
	return m.client.Disconnect(context.Background())
}

func ConnectAllMongoDBs() (map[string]MongoConnectorInterface, error) {

	connectors := make(map[string]MongoConnectorInterface)
	uris := map[string]string{
		"mongo1": os.Getenv("MONGO_URI_DB1"),
		"mongo2": os.Getenv("MONGO_URI_DB2"),
	}

	for name, uri := range uris {
		if uri == "" {
			fmt.Printf("⚠️ URI for %s is empty. Skipping...\n", name)
			continue
		}

		conn, err := NewMongoDBConnector(uri)
		if err != nil {
			fmt.Printf("❌ Failed to connect to %s: %v\n", name, err)
			continue
		}
		connectors[name] = conn
	}

	return connectors, nil
}
