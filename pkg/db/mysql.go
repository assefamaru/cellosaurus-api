package db

import (
	"database/sql"
	"fmt"
	"os"

	// go mysql driver
	_ "github.com/go-sql-driver/mysql"
)

const (
	mysqlServiceUserEnv = "MYSQL_SERVICE_USER"
	mysqlServicePassEnv = "MYSQL_SERVICE_PASS"
	mysqlServiceDBEnv   = "MYSQL_SERVICE_DB"
	mysqlServiceHostEnv = "MYSQL_SERVICE_HOST"
	mysqlServicePortEnv = "MYSQL_SERVICE_PORT"
)

type MySQL struct {
	user string
	pass string
	db   string
	host string
	port string
}

// NewMySQL returns a new MySQL instance using arguments.
func NewMySQL(user, pass, dbName, host, port string) *MySQL {
	return &MySQL{
		user: user,
		pass: pass,
		db:   dbName,
		host: host,
		port: port,
	}
}

// NewMySQLFromEnv returns a new MySQL instance using
// environment variables.
func NewMySQLFromEnv() (*MySQL, error) {
	mysqlConfig := &MySQL{}
	errPrefixFormatStr := "missing environment variable: %v"
	if mysqlConfig.user = os.Getenv(mysqlServiceUserEnv); mysqlConfig.user == "" {
		return nil, fmt.Errorf(errPrefixFormatStr, mysqlServiceUserEnv)
	}
	if mysqlConfig.pass = os.Getenv(mysqlServicePassEnv); mysqlConfig.pass == "" {
		return nil, fmt.Errorf(errPrefixFormatStr, mysqlServicePassEnv)
	}
	if mysqlConfig.db = os.Getenv(mysqlServiceDBEnv); mysqlConfig.db == "" {
		return nil, fmt.Errorf(errPrefixFormatStr, mysqlServiceDBEnv)
	}
	if mysqlConfig.host = os.Getenv(mysqlServiceHostEnv); mysqlConfig.host == "" {
		return nil, fmt.Errorf(errPrefixFormatStr, mysqlServiceHostEnv)
	}
	if mysqlConfig.port = os.Getenv(mysqlServicePortEnv); mysqlConfig.port == "" {
		return nil, fmt.Errorf(errPrefixFormatStr, mysqlServicePortEnv)
	}
	return mysqlConfig, nil
}

// Connect returns a new mysql connection.
func (m *MySQL) Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", m.DSN())
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

// DSN returns the Data Source Name from MySQL suitable for sql.Open.
func (m *MySQL) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", m.user, m.pass, m.host, m.port, m.db)
}
