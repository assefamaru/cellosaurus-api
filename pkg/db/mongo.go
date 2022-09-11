package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	mongoServiceHostEnv = "MONGODB_SERVICE_HOST"
	mongoServicePortEnv = "MONGODB_SERVICE_PORT"
)

type Mongo struct {
	host string
	port string
}

// NewMongo returns a new Mongo instance using arguments.
func NewMongo(host, port string) *Mongo {
	return &Mongo{
		host: host,
		port: port,
	}
}

// NewMongoFromEnv returns a new Mongo instance using
// environment variables.
func NewMongoFromEnv() (*Mongo, error) {
	mongoConfig := &Mongo{}
	errPrefixFormatStr := "missing environment variable: %v"
	if mongoConfig.host = os.Getenv(mongoServiceHostEnv); mongoConfig.host == "" {
		return nil, fmt.Errorf(errPrefixFormatStr, mongoServiceHostEnv)
	}
	if mongoConfig.port = os.Getenv(mongoServicePortEnv); mongoConfig.port == "" {
		return nil, fmt.Errorf(errPrefixFormatStr, mongoServicePortEnv)
	}
	return mongoConfig, nil
}

// Connect returns a new mongo connection.
func (m *Mongo) Connect() (*mongo.Client, context.Context, context.CancelFunc, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(m.DSN()))
	if err != nil {
		return nil, nil, nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	if err := client.Connect(ctx); err != nil {
		return nil, ctx, cancel, err
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, ctx, cancel, err
	}
	return client, ctx, cancel, nil
}

// DSN returns the Mongo Data Source Name connection URI.
func (m *Mongo) DSN() string {
	return fmt.Sprintf("mongodb://%s:%s", m.host, m.port)
}
