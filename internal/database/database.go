package database

import (
	"backend/internal/config"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// DB the List repository
type DB struct {
	DB *sql.DB
}

// NewDB returns a Repository
func NewDB(conf config.DBConfig) DB {

	connectionString := fmt.Sprintf("%s:%s@%s/%s", conf.Username, conf.Password, conf.URL, conf.DatabaseName)
	d, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error())
	}
	return DB{DB: d}
}
