package db

import (
	"database/sql"
	"family/conf"
	"family/util"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func dbConn() *sql.DB {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", conf.BC.Data["db"]["user"], conf.BC.Data["db"]["pass"], conf.BC.Data["db"]["host"], conf.BC.Data["db"]["port"], conf.BC.Data["db"]["db"])
	db, err := sql.Open("mysql", url)
	util.Dealerr(err, util.Return)
	return db
}
