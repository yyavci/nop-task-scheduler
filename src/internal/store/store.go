package store

import (
	"errors"
	"fmt"

	"github.com/yyavci/nop-task-scheduler/internal/database"
)

// structs
type Store struct {
	Id   int
	Url  string
	Name string
}

// errors
var errCannotGetStores = errors.New("cannot get stores")
var errCannotParseStores = errors.New("cannot parse stores")
var errNoStoreFound = errors.New("no store found")

func GetStores() ([]Store, error) {

	fmt.Printf("getting stores...\n")

	db, err := database.OpenConnection()
	defer database.CloseConnection(db)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT Id,Url,Name FROM Store")
	if err != nil {
		fmt.Printf("error: cannot get stores! err:%+v\n %+v\n", errCannotGetStores, err)
		return nil, errCannotGetStores
	}
	var stores []Store

	for rows.Next() {
		var store Store
		if err := rows.Scan(&store.Id, &store.Url, &store.Name); err != nil {
			fmt.Printf("error: cannot parse stores! err:%+v\n %+v\n", errCannotParseStores, err)
			return nil, errCannotParseStores
		}
		stores = append(stores, store)
	}

	if len(stores) == 0 {
		fmt.Printf("error: cannot parse stores! err:%+v\n %+v\n", errNoStoreFound, err)
		return nil, errNoStoreFound
	}
	return stores, nil
}
