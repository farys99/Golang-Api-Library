package app

import (
	"ApiLibrary/helper"
	"database/sql"
	"os"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
