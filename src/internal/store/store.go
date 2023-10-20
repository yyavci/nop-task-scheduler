package store

import (
	"errors"
	"fmt"

	"github.com/yyavci/nop-task-scheduler/internal/database"
)

type Store struct {
	Id  int
	Url string
	Name string
}

func GetStores() ([]Store, error) {

	fmt.Printf("getting stores...\n")

	db, err := database.OpenConnection()
	defer database.CloseConnection(db)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT Id,Url,Name FROM Store")
	if err != nil {
		fmt.Printf("Cannot get stores! Err:%+v\n", err)
		return nil, err
	}
	var stores []Store

	for rows.Next() {
		var store Store
		if err := rows.Scan(&store.Id, &store.Url , &store.Name); err != nil {
			fmt.Printf("Cannot parse stores! Err:%+v\n", err)
			return nil, err
		}
		stores = append(stores, store)
	}

	if len(stores) == 0 {
		return nil, errors.New("no store found")
	}
	return stores, nil
}
