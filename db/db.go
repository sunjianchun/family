package db

import (
	"database/sql"
)

//查询
type Rows interface {
	Next() bool
	Scan(dest ...interface{}) error
	Close() error
}

type DBScanner interface {
	Scan(Rows)
}

type DB struct {
	Conn     *sql.DB
	QueryStr string
	Scanner  DBScanner
}

func NewDB(querystr string, scanner DBScanner) *DB {
	return &DB{
		dbConn(),
		querystr,
		scanner,
	}
}
