package customer

import (
	"database/sql"
	"fmt"

	"github.com/yyavci/nop-task-scheduler/internal/database"
)

type Customer struct {
	Id         int
	SystemName string
}

func GetScheduleTaskCustomer(systemName string) (*Customer, error) {

	var scheduleTaskCustomer Customer

	db, err := database.OpenConnection()
	defer database.CloseConnection(db)

	if err != nil {
		return nil, err
	}

	fmt.Println("getting schedule task customer...")

	row := db.QueryRow("SELECT Id,SystemName FROM Customer WHERE SystemName = ?", systemName)
	if err := row.Scan(&scheduleTaskCustomer.Id, &scheduleTaskCustomer.SystemName); err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("Cannot get schedule task customer! Err:%+v\n", err)
			return nil, err
		}
	}
	return &scheduleTaskCustomer, nil
}
