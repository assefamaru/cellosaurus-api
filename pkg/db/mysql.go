package db

import (
	// go mysql driver
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysqlServiceUserEnv = "MYSQL_SERVICE_USER"
	mysqlServicePassEnv = "MYSQL_SERVICE_PASS"
	mysqlServiceDBEnv   = "MYSQL_SERVICE_DB"
	mysqlServiceHostEnv = "MYSQL_SERVICE_HOST"
	mysqlServicePortEnv = "MYSQL_SERVICE_PORT"
)

// MySQLClient is a client interface for
// managing a pool of zero or more
// underlying mysql connections.
type MySQLClient struct {
	conn   *sql.DB
	config *mySQLConfig
}

type mySQLConfig struct {
	user string
	pass string
	db   string
	host string
	port string
}

// MySQLClient creates a new MySQL.
func NewMySQLClient(ctx context.Context) (*MySQLClient, error) {
	config, err := newMySQLConfigFromEnv()
	if err != nil {
		return nil, err
	}
	conn, err := sql.Open("mysql", mysqlDSN(config))
	if err != nil {
		return nil, err
	}
	if err := conn.PingContext(ctx); err != nil {
		conn.Close()
		return nil, err
	}
	return &MySQLClient{
		conn:   conn,
		config: config,
	}, nil
}

// Close closes the underlying mysql connection.
func (c *MySQLClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

// Conn is an accessor method.
func (c *MySQLClient) Conn() *sql.DB {
	return c.conn
}

// newMySQLConfigFromEnv creates a new mySQLConfig
// using environment variables. An error will be
// thrown if any of the environment variables are
// missing.
func newMySQLConfigFromEnv() (*mySQLConfig, error) {
	config := &mySQLConfig{
		user: os.Getenv(mysqlServiceUserEnv),
		pass: os.Getenv(mysqlServicePassEnv),
		db:   os.Getenv(mysqlServiceDBEnv),
		host: os.Getenv(mysqlServiceHostEnv),
		port: os.Getenv(mysqlServicePortEnv),
	}
	if config.user == "" {
		return nil, fmt.Errorf("%w: %s", errMissingEnv, mysqlServiceUserEnv)
	}
	if config.pass == "" {
		return nil, fmt.Errorf("%w: %s", errMissingEnv, mysqlServicePassEnv)
	}
	if config.db == "" {
		return nil, fmt.Errorf("%w: %s", errMissingEnv, mysqlServiceDBEnv)
	}
	if config.host == "" {
		return nil, fmt.Errorf("%w: %s", errMissingEnv, mysqlServiceHostEnv)
	}
	if config.port == "" {
		return nil, fmt.Errorf("%w: %s", errMissingEnv, mysqlServicePortEnv)
	}
	return config, nil
}

// mysqlDSN returns the MySQL Data Source Name
// suitable for sql.Open.
func mysqlDSN(config *mySQLConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.user, config.pass, config.host, config.port, config.db)
}
