package db

import (
	"database/sql"
	"family/conf"
	"family/util"
)

//Do 数据库操作函数
func (db *DB) Do(operate string, args ...interface{}) []map[string]interface{} {

	defer db.Conn.Close()
	if operate == conf.Query {
		var rows *sql.Rows
		stmt, err := db.Conn.Prepare(db.QueryStr)
		util.Dealerr(err, util.Return)
		defer stmt.Close()

		if len(args) > 0 {
			rows, err = stmt.Query(args...)
		} else {
			rows, err = stmt.Query()
		}
		util.Dealerr(err, util.Return)
		defer rows.Close()

		colnames, err := rows.Columns()
		util.Dealerr(err, util.Return)
		var result = []map[string]interface{}{}
		rawResult := make([]sql.RawBytes, len(colnames))
		dest := make([]interface{}, len(colnames))
		for i, _ := range rawResult {
			dest[i] = &rawResult[i]
		}
		row1 := 0
		for rows.Next() {
			rowResult := make(map[string]interface{}, len(colnames))
			rows.Scan(dest...)
			for i, v := range rawResult {
				if v == nil {
					rowResult[colnames[i]] = ""
				} else {
					rowResult[colnames[i]] = string(v)
				}
			}
			result = append(result, rowResult)
			row1++
		}
		return result
	} else {
		stmt, err := db.Conn.Prepare(db.QueryStr)
		util.Dealerr(err, util.Return)
		if len(args) > 0 {
			_, err = stmt.Exec(args...)
		} else {
			_, err = stmt.Exec()
		}
		util.Dealerr(err, util.Return)
	}
	return nil
}
