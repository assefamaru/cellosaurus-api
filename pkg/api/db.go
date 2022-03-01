package api

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// go mysql driver
	_ "github.com/go-sql-driver/mysql"
)

type mysqlConfig struct {
	// MySQL user
	User string

	// MySQL password
	Pass string

	// MySQL database
	DB string

	// MySQL host
	Host string

	// MySQL port
	Port string
}

var mysqlConf mysqlConfig

// Updates mysqlConf using environment variables.
func SetupMySQLConfig() {
	// set user
	if mysqlConf.User = os.Getenv("MYSQL_SERVICE_USER"); mysqlConf.User == "" {
		log.Fatal("MYSQL_SERVICE_USER environment variable is missing")
	}

	// set password
	if mysqlConf.Pass = os.Getenv("MYSQL_SERVICE_PASS"); mysqlConf.Pass == "" {
		log.Fatal("MYSQL_SERVICE_PASS environment variable is missing")
	}

	// set database
	if mysqlConf.DB = os.Getenv("MYSQL_SERVICE_DB"); mysqlConf.DB == "" {
		log.Fatal("MYSQL_SERVICE_DB environment variable is missing")
	}

	// set host
	if mysqlConf.Host = os.Getenv("MYSQL_SERVICE_HOST"); mysqlConf.Host == "" {
		log.Fatal("MYSQL_SERVICE_HOST environment variable is missing")
	}

	// set port
	if mysqlConf.Port = os.Getenv("MYSQL_SERVICE_PORT"); mysqlConf.Port == "" {
		log.Fatal("MYSQL_SERVICE_PORT environment variable is missing")
	}
}

// Returns a connection string suitable for sql.Open.
func (c mysqlConfig) dataSourceName() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.User, c.Pass, c.Host, c.Port, c.DB)
}

// Creates a new database connection.
func Database() (*sql.DB, error) {
	db, err := sql.Open("mysql", mysqlConf.dataSourceName())
	if err != nil {
		return nil, fmt.Errorf("mysql: could not get a connection: %v", err)
	}
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("mysql: could not establish a good connection: %v", err)
	}
	return db, nil
}
