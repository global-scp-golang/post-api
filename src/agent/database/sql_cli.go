package database

import "database/sql"

type Db struct {
	DB *sql.DB
}

var initMySQLDB = new(Db)

func SetDataBase(db *sql.DB) {
	initMySQLDB.DB = db
}

func Connect() *Db {
	return initMySQLDB
}
