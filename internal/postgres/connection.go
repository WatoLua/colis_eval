package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DBInfos struct {
	Host     string
	Port     uint16
	User     string
	Password string
	Dbname   string
}

var globalConnection *sqlx.DB = nil

func createConnection(infos DBInfos) *sqlx.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		infos.Host, infos.Port, infos.User, infos.Password, infos.Dbname)

	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func GetConnection(infos DBInfos) *sqlx.DB {

	if globalConnection == nil || globalConnection.Ping() != nil {
		globalConnection = createConnection(infos)
	}

	return globalConnection
}

func IsConnectionAlive() bool {
	return globalConnection == nil || globalConnection.Ping() == nil
}

func CloseConnection() {
	if globalConnection != nil {
		if err := globalConnection.Ping(); err != nil {
			globalConnection.Close()
		}
	}
}

func ResetTable(infos DBInfos, tableName string) {
	if !IsConnectionAlive() {
		createConnection(infos)
		defer CloseConnection()
	}

	_, err := globalConnection.Exec(fmt.Sprintf("drop table if exists %v", tableName))
	if err != nil {
		panic(err)
	}
}
