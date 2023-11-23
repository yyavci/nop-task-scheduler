package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/yyavci/nop-task-scheduler/internal/config"

	_ "github.com/microsoft/go-mssqldb"
)

// variables
var connStr string

// errors
var errOpenConnection = errors.New("open database connection error")
var errCloseConnection = errors.New("close database connection error")
var errCannotConnectDatabase = errors.New("cannot connect database")
var errConnStrIsNotSet = errors.New("connstr is not set")

func OpenConnection() (*sql.DB, error) {
	db, err := sql.Open("mssql", connStr)
	if err != nil {
		fmt.Printf("error: error opening connection. err:%+v\n", err)
		return nil, errOpenConnection
	}
	return db, nil
}

func CloseConnection(db *sql.DB) error {
	err := db.Close()
	if err != nil {
		fmt.Printf("error: error closing connection. err:%+v\n", err)
		return errCloseConnection
	}

	return nil
}

func Init(appConfig config.AppConfig) error {

	fmt.Println("initializing database...")

	connStr = os.Getenv("CONN_STR")

	if len(connStr) == 0 {
		connStr = appConfig.ConnectionString
	}

	if len(connStr) == 0 {
		fmt.Printf("error: set CONN_STR environment variable first!\n")
		return errConnStrIsNotSet
	}

	db, err := OpenConnection()
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("error: connection error. err:%+v\n %+v\n", errCannotConnectDatabase, err)
		return errCannotConnectDatabase
	}

	return nil

}
