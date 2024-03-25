package postgres

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "eval"
)

var globalPsqlConnection *sql.DB = nil

func GetConnection() *sql.DB {

	if globalPsqlConnection == nil {
		globalPsqlConnection = createConnection()
	} else if err := globalPsqlConnection.Ping(); err != nil {
		CloseConnection()
		globalPsqlConnection = createConnection()
	}

	return globalPsqlConnection
}

func createConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func CloseConnection() {
	if globalPsqlConnection != nil {
		if err := globalPsqlConnection.Ping(); err != nil {
			globalPsqlConnection.Close()
		}
	}
}
