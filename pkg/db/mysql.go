package db

import (
	"database/sql"
	"fmt"
	"os"

	// go mysql driver
	_ "github.com/go-sql-driver/mysql"
)

const (
	mysqlUserEnv = "MYSQL_SERVICE_USER"
	mysqlPassEnv = "MYSQL_SERVICE_PASS"
	mysqlDBEnv   = "MYSQL_SERVICE_DB"
	mysqlHostEnv = "MYSQL_SERVICE_HOST"
	mysqlPortEnv = "MYSQL_SERVICE_PORT"
)

type mysqlConfig struct {
	user string
	pass string
	db   string
	host string
	port string
}

// ConnectMySQL returns a new mysql connection.
func ConnectMySQL() (*sql.DB, error) {
	config, err := mysqlConfigFromEnv()
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("mysql", formatDSN(config))
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

// mysqlConfigFromEnv returns a new mysqlConfig from environment
// variables and defaults.
func mysqlConfigFromEnv() (*mysqlConfig, error) {
	var config *mysqlConfig
	if config.user = os.Getenv(mysqlUserEnv); config.user == "" {
		return nil, fmt.Errorf("missing environment variable: %v", mysqlUserEnv)
	}
	if config.pass = os.Getenv(mysqlPassEnv); config.pass == "" {
		return nil, fmt.Errorf("missing environment variable: %v", mysqlPassEnv)
	}
	if config.db = os.Getenv(mysqlDBEnv); config.db == "" {
		return nil, fmt.Errorf("missing environment variable: %v", mysqlDBEnv)
	}
	if config.host = os.Getenv(mysqlHostEnv); config.host == "" {
		config.host = "localhost"
	}
	if config.port = os.Getenv(mysqlPortEnv); config.port == "" {
		config.port = "3306"
	}
	return config, nil
}

// formatDSN returns a datasource name suitable for sql.Open.
func formatDSN(c *mysqlConfig) string {
	if c == nil {
		return ""
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.user, c.pass, c.host, c.port, c.db)
}
