package customer

import (
	"database/sql"
	"fmt"
)

type Customer struct {
	Id         int
	SystemName string
}

func GetScheduleTaskCustomer(database *sql.DB, systemName string) (*Customer, error) {

	var scheduleTaskCustomer Customer

	fmt.Println("getting schedule task customer...")

	row := database.QueryRow("SELECT Id,SystemName FROM Customer WHERE SystemName = ?", systemName)
	if err := row.Scan(&scheduleTaskCustomer.Id, &scheduleTaskCustomer.SystemName); err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("Cannot get schedule task customer! Err:%+v\n", err)
			return nil, err
		}
	}
	//fmt.Printf("schedule task customerId:%d\n", scheduleTaskCustomer.Id)

	return &scheduleTaskCustomer, nil
}
