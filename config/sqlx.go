package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func GetSqlx() *sqlx.DB {
	return sqlx.MustConnect("mysql", DB_CONNECTION)
}
