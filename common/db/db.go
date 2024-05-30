package db

import (
	"context"

	"gorm.io/gorm"
)

type DBClient interface {
	GetConn(c *context.Context) (*gorm.DB, error)
	// Exec(c *context.Context, Tablename string) error
}

type Database struct {
	addr      string
	port      string
	password  string
	username  string
	dbname    string
	charset   string
	parseTime string
	loc       string
}

func CreateNewDBimpl() *Database {
	// addr string,
	// port string,
	// password string,
	// username string,
	// dbname string,
	// charset string,
	// parseTime string,
	// loc string,
	return &Database{
		addr: "127.0.0.1",
	}
}
