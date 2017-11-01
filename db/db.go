package db

import (
	"database/sql"
)

//查询
//type Rows interface {
//	Next() bool
//	Scan(dest ...interface{}) error
//	Close() error
//}

//type DBScanner interface {
//	Scan([]map[string]interface{})
//}

type DB struct {
	Conn     *sql.DB
	QueryStr string
	//	Scanner  DBScanner
}

//func NewDB(querystr string, scanner DBScanner) *DB {
func NewDB(querystr string) *DB {
	return &DB{
		dbConn(),
		querystr,
		//		scanner,
	}
}
