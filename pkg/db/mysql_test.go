package db

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMySQL(t *testing.T) {
	cases := []struct {
		config MySQL
		dsn    string
	}{
		{
			config: MySQL{user: "brucewayne", pass: "batman", db: "justiceleague", host: "gotham", port: "comic"},
			dsn:    "brucewayne:batman@tcp(gotham:comic)/justiceleague",
		},
		{
			config: MySQL{user: "clarkkent", pass: "superman", db: "justiceleague", host: "metropolis", port: "comic"},
			dsn:    "clarkkent:superman@tcp(metropolis:comic)/justiceleague",
		},
	}

	for _, c := range cases {
		mysql := NewMySQL(c.config.user, c.config.pass, c.config.db, c.config.host, c.config.port)
		assert.Equal(t, c.dsn, mysql.DSN())
	}
}

func TestNewMySQLFromEnv(t *testing.T) {
	cases := []struct {
		config MySQL
		dsn    string
	}{
		{
			config: MySQL{user: "brucewayne", pass: "batman", db: "justiceleague", host: "gotham", port: "comic"},
			dsn:    "brucewayne:batman@tcp(gotham:comic)/justiceleague",
		},
		{
			config: MySQL{user: "clarkkent", pass: "superman", db: "justiceleague", host: "metropolis", port: "comic"},
			dsn:    "clarkkent:superman@tcp(metropolis:comic)/justiceleague",
		},
	}

	for _, c := range cases {
		assert.Nil(t, os.Setenv(mysqlServiceUserEnv, c.config.user))
		assert.Nil(t, os.Setenv(mysqlServicePassEnv, c.config.pass))
		assert.Nil(t, os.Setenv(mysqlServiceDBEnv, c.config.db))
		assert.Nil(t, os.Setenv(mysqlServiceHostEnv, c.config.host))
		assert.Nil(t, os.Setenv(mysqlServicePortEnv, c.config.port))

		mysql, err := NewMySQLFromEnv()
		assert.Nil(t, err)
		assert.Equal(t, c.dsn, mysql.DSN())
	}
}
