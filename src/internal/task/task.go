package task

import (
	"database/sql"
	"errors"
	"fmt"
)

type ScheduleTask struct {
	Id             int
	Name           string
	Enabled        bool
	CronExpression string
}

func GetScheduleTasks(database *sql.DB) ([]ScheduleTask, error) {

	fmt.Println("getting schedule tasks...")

	rows, err := database.Query("SELECT Id , Name , Enabled FROM ScheduleTask")
	if err != nil {
		fmt.Printf("Cannot get schedule tasks! Err:%+v\n", err)
		return nil, err
	}
	var scheduleTasks []ScheduleTask

	for rows.Next() {
		var task ScheduleTask
		if err := rows.Scan(&task.Id, &task.Name, &task.Enabled); err != nil {
			fmt.Printf("Cannot get schedule tasks! Err:%+v\n", err)
			return nil, err
		}
		task.CronExpression = "*/1 * * * *" // TODO get it from database
		scheduleTasks = append(scheduleTasks, task)

	}

	if len(scheduleTasks) == 0 {
		fmt.Printf("schedule task count is 0!\n")
		return nil, errors.New("schedule task count is 0")
	}

	return scheduleTasks, nil
}

func DoTask(task ScheduleTask) {
	fmt.Printf("[%d]'%s' task started. \n", task.Id, task.Name)
}
