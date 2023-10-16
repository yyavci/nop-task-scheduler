package task

import (
	"errors"
	"fmt"
	"github.com/yyavci/nop-task-scheduler/internal/database"
	"github.com/yyavci/nop-task-scheduler/internal/config"
	"github.com/yyavci/nop-task-scheduler/internal/http"
)

type ScheduleTask struct {
	Id             int
	Name           string
	Enabled        bool
	CronExpression string
}

func GetScheduleTasks() ([]ScheduleTask, error) {

	fmt.Println("getting schedule tasks...")

	db, err := database.OpenConnection()
	if err != nil {
		fmt.Printf("Cannot open db connection! Err:%+v\n", err)
		return nil, err
	}

	rows, err := db.Query("SELECT Id , Name , Enabled FROM ScheduleTask")
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

	defer database.CloseConnection(db)

	return scheduleTasks, nil
}

func DoTask(task ScheduleTask, conf config.AppConfig) {
	fmt.Printf("[%d]'%s' task started. \n", task.Id, task.Name)

	fmt.Println(conf.StoreUrl)
	response, err := http.PostJsonRequest(conf.StoreUrl+"/ScheduleTask/Run", "{}")

	if err != nil {
		fmt.Printf("error posting request! err:%s\n", err.Error())
		return
	}

	fmt.Printf("response status-code:%d status:%s", response.StatusCode, response.Status)

}
