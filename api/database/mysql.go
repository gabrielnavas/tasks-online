package database

import (
	"database/sql"
	"fmt"
	"time"
)

func OpenMysqlConnection(
	DB_MYSQL_USERNAME,
	DB_MYSQL_PASSWORD,
	DB_MYSQL_HOST,
	DB_MYSQL_PORT,
	DB_MYSQL_DATABASE_NAME string,
) *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=UTC",
		DB_MYSQL_USERNAME,
		DB_MYSQL_PASSWORD,
		DB_MYSQL_HOST,
		DB_MYSQL_PORT,
		DB_MYSQL_DATABASE_NAME,
	))
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
