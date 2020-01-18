package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/inwady/easyconfig"
	"log"
	"errors"
)

var localDB *sql.DB

var (
	host = easyconfig.GetString("database.host", "127.0.0.1")
	port = easyconfig.GetInt("database.port", 5432)

	user = easyconfig.GetString("database.user", "postgres")
	password = easyconfig.GetString("database.password", "postgres")
	dbname = easyconfig.GetString("database.dbname", "rhack")
)

var (
	CannotInsertError = errors.New("cannot insert in database")
	CannotSelectError = errors.New("cannot select in database")
	CannotUpdateError = errors.New("cannot update in database")
	BadDataInTupleError = errors.New("bad data in tuple")
)

func init() {
	var dbinfo string
	if (password != "") {
		dbinfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
	}  else {
		dbinfo = fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
			host, port, user, dbname)
	}

	var err error
	localDB, err = sql.Open("postgres", dbinfo)
	if err != nil {
		log.Panicf("cannot connect db, error: %v", err)
	}

	// var version string
	//
	// log.Println("try connect to postgresql")
	// err = localDB.QueryRow("SELECT version()").Scan(&version)
	// if err != nil {
	// 	log.Panicf("bad database, error: %v", err)
	// }
	//
	// log.Println("database version", version)
}

func GetConnection() *sql.DB {
	if localDB == nil {
		panic("cannot get connection")
	}

	return localDB
}

func DataBaseClose() {
	localDB.Close()
}
