package db

import (
	"database/sql"
	"family/conf"
	"family/util"
	"strconv"
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
		}
		return result
	} else {
		stmt, err := db.Conn.Prepare(db.QueryStr)
		util.Dealerr(err, util.Return)
		if len(args) > 0 {
			var res sql.Result
			res, err := stmt.Exec(args...)
			util.Dealerr(err, util.Return)
			if operate == conf.Insert {
				id, err := res.LastInsertId()
				util.Dealerr(err, util.Return)
				var result = []map[string]interface{}{}
				var innerRes = make(map[string]interface{})
				innerRes["lastInsertId"] = strconv.FormatInt(id, 10)
				result = append(result, innerRes)
				return result
			}
		} else {
			var res sql.Result
			res, err := stmt.Exec()
			util.Dealerr(err, util.Return)
			if operate == conf.Insert {
				id, err := res.LastInsertId()
				util.Dealerr(err, util.Return)
				var result = []map[string]interface{}{}
				var innerRes = make(map[string]interface{})
				innerRes["lastInsertId"] = strconv.FormatInt(id, 10)
				result = append(result, innerRes)
				return result
			}
		}

	}
	return nil
}
