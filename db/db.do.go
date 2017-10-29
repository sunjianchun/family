package db

import (
	"family/conf"
	"family/util"
)

//Do 数据库操作函数
func (db *DB) Do(operate string, args ...interface{}) {

	if operate == conf.Query {
		rows, err := db.Conn.Query(db.QueryStr)
		util.Dealerr(err, util.Return)

		defer rows.Close()
		db.Scanner.Scan(rows)
	} else {
		stmt, err := db.Conn.Prepare(db.QueryStr)
		util.Dealerr(err, util.Return)

		//if len(args) > 0 {
		_, err = stmt.Exec(args)
		//} else {
		//	_, err = stmt.Exec()
		//}
		util.Dealerr(err, util.Return)
	}

	db.Conn.Close()
}
