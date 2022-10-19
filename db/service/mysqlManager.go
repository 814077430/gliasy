package service

import (
	"server/conf"
	"database/sql"
	"github.com/name5566/leaf/log"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func initMysql() {
	if conf.Server.DBMysqlMaxConnNum <= 0 {
		conf.Server.DBMysqlMaxConnNum = 100
	}

	Db, _ = sql.Open("mysql", conf.Server.DBMysqlUrl)

	err := Db.Ping()
	if err != nil {
		log.Debug("Mysql ping failed:", err)
		return
	}

	log.Debug("Db mysql connect success")

	Db.SetMaxOpenConns(conf.Server.DBMysqlMaxConnNum)
	Db.SetMaxIdleConns(conf.Server.DBMysqlMaxConnNum/2)
	Db.SetConnMaxLifetime(0)
}



