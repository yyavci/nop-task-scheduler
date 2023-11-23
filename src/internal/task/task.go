package task

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/yyavci/nop-task-scheduler/internal/database"
	"github.com/yyavci/nop-task-scheduler/internal/http"
	"github.com/yyavci/nop-task-scheduler/internal/store"
)

// structs
type ScheduleTask struct {
	Id             int
	Name           string
	Type           string
	Enabled        bool
	CronExpression string
}
type ScheduleTaskRunRequest struct {
	TaskType string
}

// errors
var errCannotGetTasks = errors.New("cannot get schedule tasks")
var errCannotParseTasks = errors.New("cannot parse schedule tasks")
var errNoScheduleTasksFound = errors.New("no schedule tasks found")
var errUpdateTaskStarted = errors.New("cannot update schedule task started")
var errParseRequest = errors.New("cannot parse request")
var errSendingRequest = errors.New("error posting request")
var errTaskIdCannotZero = errors.New("task id cannot be zero")

func GetScheduleTasks() ([]ScheduleTask, error) {

	fmt.Println("getting schedule tasks...")

	db, err := database.OpenConnection()
	defer database.CloseConnection(db)
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT Id , Name , Type, Enabled , ISNULL(CronExpression,'') FROM ScheduleTask")
	if err != nil {
		fmt.Printf("error:cannot get schedule tasks! err:%+v\n %+v\n", errCannotGetTasks, err)
		return nil, errCannotGetTasks
	}
	var scheduleTasks []ScheduleTask

	for rows.Next() {
		var task ScheduleTask
		if err := rows.Scan(&task.Id, &task.Name, &task.Type, &task.Enabled, &task.CronExpression); err != nil {
			fmt.Printf("error:cannot parse schedule tasks! err:%+v\n %+v\n", errCannotParseTasks, err)
			return nil, errCannotParseTasks
		}
		if task.Enabled && len(task.CronExpression) > 0 {
			scheduleTasks = append(scheduleTasks, task)
		}

	}

	if len(scheduleTasks) == 0 {
		fmt.Printf("error:no schedule tasks found! err:%+v\n", errNoScheduleTasksFound)
		return nil, errNoScheduleTasksFound
	}
	return scheduleTasks, nil
}

func DoTask(task ScheduleTask, store store.Store) {
	fmt.Printf("[%d]'%s' task for store '%s' started. \n", task.Id, task.Name, store.Name)

	err := UpdateTask(task.Id, true, false)
	if err != nil {
		fmt.Printf("error:error updating task started! err:%+v\n %+v\n", errUpdateTaskStarted, err)
		UpdateTask(task.Id, false, false)
		return
	}

	request := &ScheduleTaskRunRequest{TaskType: task.Type}

	jsonStr, err := json.Marshal(request)
	if err != nil {
		fmt.Printf("error:error parsing request! err:%+v\n %+v\n", errParseRequest, err)
		UpdateTask(task.Id, false, false)
		return
	}

	response, err := http.PostJsonRequest(store.Url+"ScheduleTask/RunTask", string(jsonStr))

	if err != nil {
		UpdateTask(task.Id, false, false)
		return
	}
	fmt.Printf("response status-code:%d status:%s", response.StatusCode, response.Status)

	if response.StatusCode < 200 || response.StatusCode > 400 {
		fmt.Printf("error:error posting request! err:%+v\n %+v\n", errSendingRequest, err)
		UpdateTask(task.Id, false, false)
		return
	}

	UpdateTask(task.Id, false, true)

}

func UpdateTask(id int, start bool, ok bool) error {

	if id == 0 {
		fmt.Printf("error:task id cannot be 0! err:%+v\n", errTaskIdCannotZero)
		return errTaskIdCannotZero
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
