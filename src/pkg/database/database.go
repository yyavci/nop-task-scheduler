package database

import (
	"database/sql"
	"fmt"
	"nop-task-scheduler/pkg/config"
	"os"

	_ "github.com/microsoft/go-mssqldb"
)

func CheckDatabaseConnection(appConfig config.AppConfig) (*sql.DB, error) {

	fmt.Println("checking database connection...")

	db, err := sql.Open("mssql", os.Getenv("CONN_STR"))
	if err != nil {
		fmt.Printf("Error occured opening database connection! Err:%s\n", err)
		return nil, err
	}
	//defer database.Close()
	err = db.Ping()
	if err != nil {
		fmt.Printf("cannot connect to database! Err:%s\n", err)
		return nil, err
	}

	return db, nil

}
