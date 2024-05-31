package db

import (
	"context"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBClient interface {
	GetConn(c *context.Context) (*gorm.DB, error)
	// Exec(c *context.Context, Tablename string) error
}

type Database struct {
	Addr      string `yaml: "addr"`
	Port      int    `yaml: "port"`
	Password  string `yaml: "password"`
	Username  string `yaml: "username"`
	DBname    string `yaml: "dbname"`
	Charset   string `yaml: "charset"`
	ParseTime string `yaml: "parsetime"`
	Loc       string `yaml: "loc"`
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
		Addr:      "127.0.0.1",
		Port:      3306,
		Password:  "123456",
		Username:  "admin",
		DBname:    "user",
		Charset:   "utf8mb4",
		ParseTime: "True",
		Loc:       "Loc",
	}
}

func (d *Database) DSN() string {
	connstr := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=%v&loc=%v",
		d.Username, d.Password, d.Addr, d.Port, d.DBname, d.Charset, d.ParseTime, d.Loc)
	return connstr
}

func (d *Database) GetConn(c *context.Context) (*gorm.DB, error) {
	dsn := d.DSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Get db conn failed:\n", err)
	}
	return db, nil
}
