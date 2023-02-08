package database

import (
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
)

const (
    host     = "DB-postgresql-nyc-9181987-do-user-2929533-0.b.DB.ondigitalocean.com"
    port     = 25060
    user     = "doadmin"
    password = "AVNS_8YY6GhAZgV3ajrMKS2W"
    dbname   = "Baby_Tracker"
)

var DB *sql.DB

func SetUpDatabaseConnection() {
	var err error
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require", host, port, user, password, dbname)

	// open database
	DB, err = sql.Open("postgres", psqlconn)
	CheckError(err)
	
	// check DB
	err = DB.Ping()
	CheckError(err)
}

func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}
