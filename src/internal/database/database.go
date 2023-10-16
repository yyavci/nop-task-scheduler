package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/yyavci/nop-task-scheduler/internal/config"

	_ "github.com/microsoft/go-mssqldb"
)

var connStr string

func OpenConnection() (*sql.DB, error) {
	db, err := sql.Open("mssql", connStr)
	if err != nil {
		fmt.Printf("Error occured opening database connection! Err:%+v\n", err)
		return nil, err
	}
	return db,nil
}

func CloseConnection(db *sql.DB) error {
	err := db.Close()
	if err != nil {
		fmt.Printf("Error occured closing database connection! Err:%+v\n", err)
		return err
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
		fmt.Printf("set CONN_STR environment variable first!\n")
		return errors.New("connection string is not set")
	}

	db, err := OpenConnection()
	if err != nil {
		fmt.Printf("Error occured opening database connection! Err:%+v\n", err)
		return err
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("cannot connect to database! Err:%+v\n", err)
		return err
	}

	return nil

}
