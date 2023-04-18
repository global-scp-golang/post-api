package databases

import (
	"database/sql"
	"fmt"
	"post-api/src/agent/settings"
	"time"
)

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

func OpenDB() *sql.DB {
	setting, err := settings.LoadSetting()
	if err != nil {
		panic(err)
	}

	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		setting.Database.Host, setting.Database.Port, setting.Database.User, setting.Database.Password, setting.Database.Dbname)
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxIdleConns(20) // DB Connection을 사용하지 않을때에도 Connection 풀에 최대 Connection 20개를 둔다
	db.SetMaxOpenConns(10) // DB Connection을 최대 10개까지 동시 사용가능하다

	return db
}

func OpenTestDB() *sql.DB {
	setting, err := settings.LoadTestSetting()
	if err != nil {
		panic(err)
	}

	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		setting.Database.Host, setting.Database.Port, setting.Database.User, setting.Database.Password, setting.Database.Dbname)
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxIdleConns(20) // DB Connection을 사용하지 않을때에도 Connection 풀에 최대 Connection 20개를 둔다
	db.SetMaxOpenConns(10) // DB Connection을 최대 10개까지 동시 사용가능하다

	return db
}
