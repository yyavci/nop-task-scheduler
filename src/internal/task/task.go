package task

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/yyavci/nop-task-scheduler/internal/config"
	"github.com/yyavci/nop-task-scheduler/internal/database"
	"github.com/yyavci/nop-task-scheduler/internal/http"
)

type ScheduleTask struct {
	Id             int
	Name           string
	Enabled        bool
	CronExpression string
}

type ScheduleTaskRunResponse struct {
	Success bool
	Message string
}
type ScheduleTaskRunRequest struct {
	TaskId int
}

func GetScheduleTasks() ([]ScheduleTask, error) {

	fmt.Println("getting schedule tasks...")

	db, err := database.OpenConnection()
	defer database.CloseConnection(db)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT Id , Name , Enabled , CronExpression FROM ScheduleTask")
	if err != nil {
		fmt.Printf("Cannot get schedule tasks! Err:%+v\n", err)
		return nil, err
	}
	var scheduleTasks []ScheduleTask

	for rows.Next() {
		var task ScheduleTask
		if err := rows.Scan(&task.Id, &task.Name, &task.Enabled , &task.CronExpression); err != nil {
			fmt.Printf("Cannot parse schedule tasks! Err:%+v\n", err)
			return nil, err
		}
		scheduleTasks = append(scheduleTasks, task)

	}

	if len(scheduleTasks) == 0 {
		return nil, errors.New("schedule task count is 0")
	}
	return scheduleTasks, nil
}

func DoTask(task ScheduleTask, conf config.AppConfig) {
	fmt.Printf("[%d]'%s' task started. \n", task.Id, task.Name)

	err := UpdateTask(task.Id, true, false)
	if err != nil {
		fmt.Printf("error updating task! err:%+v\n", err)
		UpdateTask(task.Id, false, false)
		return
	}

	request := &ScheduleTaskRunRequest{TaskId: task.Id}

	jsonStr, err := json.Marshal(request)
	if err != nil {
		fmt.Printf("error parsing request! err:%+v\n", err)
		UpdateTask(task.Id, false, false)
		return
	}

	fmt.Println(conf.StoreUrl)
	response, err := http.PostJsonRequest(conf.StoreUrl+"/ScheduleTask/Run", string(jsonStr))

	if err != nil {
		UpdateTask(task.Id, false, false)
		return
	}
	fmt.Printf("response status-code:%d status:%s", response.StatusCode, response.Status)

	if response.StatusCode < 200 || response.StatusCode > 400 {
		fmt.Printf("error posting request! err:%s\n", errors.New("response error"))
		UpdateTask(task.Id, false, false)
		return
	}

	var taskResponse ScheduleTaskRunResponse

	err = json.Unmarshal([]byte(response.Data), &taskResponse)

	if err != nil {
		fmt.Printf("cannot parse json response! err:%+v\n", err)
		UpdateTask(task.Id, false, false)
		return
	}

	if !taskResponse.Success {
		fmt.Printf("failed response! message:%s\n", taskResponse.Message)
		UpdateTask(task.Id, false, false)
		return
	}

	UpdateTask(task.Id, false, true)

}

func UpdateTask(id int, start bool, ok bool) error {

	if id == 0 {
		return errors.New("id cannot be zero")
	}

	db, err := database.OpenConnection()
	defer database.CloseConnection(db)
	if err != nil {
		return err
	}

	whereClause := "WHERE Id = ?"

	var query string = ""

	if start {
		query = "UPDATE ScheduleTask SET LastStartUtc = GETUTCDATE()"
	} else {
		query = "UPDATE ScheduleTask SET LastEndUtc = GETUTCDATE()"
		if ok {
			query += " LastSuccessDate = GETUTCDATE()"
		}
	}
	query += whereClause

	_, err = db.Exec(query, id)

	return err
}
