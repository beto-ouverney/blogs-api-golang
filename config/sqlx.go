package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func GetReaderSqlx() *sqlx.DB {
	return sqlx.MustConnect("mysql", DB_CONNECTION)
}

func GetWriterSqlx() *sqlx.DB {
	return sqlx.MustConnect("mysql", DB_CONNECTION)
}
