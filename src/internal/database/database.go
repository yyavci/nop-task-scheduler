package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/yyavci/nop-task-scheduler/internal/config"

	_ "github.com/microsoft/go-mssqldb"
)

func CheckDatabaseConnection(appConfig config.AppConfig) (*sql.DB, error) {

	fmt.Println("checking database connection...")

	connStr := os.Getenv("CONN_STR")

	if len(connStr) == 0 {
		connStr = appConfig.ConnectionString
	}

	if len(connStr) == 0 {
		fmt.Printf("set CONN_STR environment variable first!\n")
		return nil, errors.New("connection string is not set")
	}

	db, err := sql.Open("mssql", connStr)
	if err != nil {
		fmt.Printf("Error occured opening database connection! Err:%+v\n", err)
		return nil, err
	}
	
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Printf("cannot connect to database! Err:%+v\n", err)
		return nil, err
	}

	return db, nil

}
