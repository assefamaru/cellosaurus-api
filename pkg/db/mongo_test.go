package db

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMongo(t *testing.T) {
	cases := []struct {
		config Mongo
		dsn    string
	}{
		{
			config: Mongo{host: "gotham", port: "comic"},
			dsn:    "mongodb://gotham:comic",
		},
		{
			config: Mongo{host: "metropolis", port: "comic"},
			dsn:    "mongodb://metropolis:comic",
		},
	}

	for _, c := range cases {
		mongo := NewMongo(c.config.host, c.config.port)
		assert.Equal(t, c.dsn, mongo.DSN())
	}
}

func TestNewMongoFromEnv(t *testing.T) {
	cases := []struct {
		config Mongo
		dsn    string
	}{
		{
			config: Mongo{host: "gotham", port: "comic"},
			dsn:    "mongodb://gotham:comic",
		},
		{
			config: Mongo{host: "metropolis", port: "comic"},
			dsn:    "mongodb://metropolis:comic",
		},
	}

	for _, c := range cases {
		assert.Nil(t, os.Setenv(mongoServiceHostEnv, c.config.host))
		assert.Nil(t, os.Setenv(mongoServicePortEnv, c.config.port))

		mongo, err := NewMongoFromEnv()
		assert.Nil(t, err)
		assert.Equal(t, c.dsn, mongo.DSN())
	}
}
