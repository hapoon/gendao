package database

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Open(host string, port int, user string, password string, dbname string) error {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbname)
	dbp, err := sql.Open("mysql", dns)
	if err != nil {
		return err
	}
	db = dbp
	return nil
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	if db == nil {
		return nil, errors.New("db is empty")
	}
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
