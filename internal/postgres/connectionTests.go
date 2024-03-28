package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

var Infos = DBInfos{
	Host:     "localhost",
	Port:     5432,
	User:     "postgres",
	Password: "postgres",
	Dbname:   "eval_test",
}

var globalTestConnection *sqlx.DB = nil

func createTestConnection() *sqlx.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable",
		Infos.Host, Infos.Port, Infos.User, Infos.Password)

	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(fmt.Sprintf("create database %v", Infos.Dbname))
	if err != nil {
		panic(err)
	}
	db.Close()

	psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		Infos.Host, Infos.Port, Infos.User, Infos.Password, Infos.Dbname)

	db, err = sqlx.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func GetTestConnection() *sqlx.DB {

	if globalTestConnection == nil {
		globalTestConnection = createTestConnection()
	} else if err := globalTestConnection.Ping(); err != nil {
		CloseTestConnection()
		globalTestConnection = createTestConnection()
	}

	return globalTestConnection
}

func IsTestConnectionAlive() bool {
	return globalTestConnection == nil || globalTestConnection.Ping() == nil
}

func CloseTestConnection() {
	if globalTestConnection != nil {
		if err := globalTestConnection.Ping(); err != nil {
			globalTestConnection.Close()
		}
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable",
		Infos.Host, Infos.Port, Infos.User, Infos.Password)

	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("drop database if exists %v", Infos.Dbname))
	if err != nil {
		panic(err)
	}

}

func ResetTestTable(tableName string) {
	if !IsConnectionAlive() {
		createConnection(Infos)
		defer CloseConnection()
	}

	_, err := globalConnection.Exec(fmt.Sprintf("drop table if exists %v", tableName))
	if err != nil {
		panic(err)
	}
}
