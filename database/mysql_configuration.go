package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	log "goapp/log_config"
)

var Db *sql.DB
var err error

func MysqlConfiguration() {
	dsn := "root:password@tcp(localhost:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"

	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err := Db.Ping(); err != nil {
		panic(err)
	}

	log.Info("Connected to MySQL database successfully!")
}
