package db

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	mongoServiceHostEnv = "MONGO_SERVICE_HOST"
	mongoServicePortEnv = "MONGO_SERVICE_PORT"
)

// MongoClient is a client interface
// for managing connections to a
// MongoDB deployment.
type MongoClient struct {
	conn   *mongo.Client
	config *mongoConfig
}

type mongoConfig struct {
	host string
	port string
}

// NewMongoClient creates a new MongoClient.
func NewMongoClient(ctx context.Context) (*MongoClient, error) {
	config, err := newMongoConfigFromEnv()
	if err != nil {
		return nil, err
	}
	conn, err := mongo.NewClient(options.Client().ApplyURI(mongoDSN(config)))
	if err != nil {
		return nil, err
	}
	if err := conn.Connect(ctx); err != nil {
		return nil, err
	}
	if err := conn.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	return &MongoClient{
		conn:   conn,
		config: config,
	}, nil
}

// Close disconnects the underlying client connection pool.
func (c *MongoClient) Close(ctx context.Context) error {
	if c.conn != nil {
		return c.conn.Disconnect(ctx)
	}
	return nil
}

// Conn is an accessor method.
func (c *MongoClient) Conn() *mongo.Client {
	return c.conn
}

// newMongoConfigFromEnv creates a new mongoConfig
// using environment variables. An error will be
// thrown if any of the environment variables are
// missing.
func newMongoConfigFromEnv() (*mongoConfig, error) {
	config := &mongoConfig{
		host: os.Getenv(mongoServiceHostEnv),
		port: os.Getenv(mongoServicePortEnv),
	}
	if config.host == "" {
		return nil, fmt.Errorf("%w: %s", errMissingEnv, mongoServiceHostEnv)
	}
	if config.port == "" {
		return nil, fmt.Errorf("%w: %s", errMissingEnv, mongoServicePortEnv)
	}
	return config, nil
}

// mongoDSN returns the Mongo Data Source Name connection URI.
func mongoDSN(config *mongoConfig) string {
	return fmt.Sprintf("mongodb://%s:%s", config.host, config.port)
}
