package app

import (
	"fmt"
	"nop-task-scheduler/pkg/config"
	"nop-task-scheduler/pkg/customer"
	"nop-task-scheduler/pkg/database"
	"nop-task-scheduler/pkg/scheduler"
	"nop-task-scheduler/pkg/task"

	"github.com/go-co-op/gocron"
)

var schedule *gocron.Scheduler

func Run() {
	fmt.Println("app started")

	conf, err := config.ReadConfiguration()
	if err != nil {
		panic(err)
	}

	db, err := database.CheckDatabaseConnection(*conf)
	if err != nil {
		panic(err)
	}

	cust, err := customer.GetScheduleTaskCustomer(db, "BackgroundTask")
	if err != nil {
		panic(err)
	}
	fmt.Printf("schedule task customerId:%d\n", cust.Id)

	tasks, err := task.GetScheduleTasks(db)
	if err != nil {
		panic(err)
	}

	sch, err := scheduler.InitScheduler(tasks)
	if err != nil {
		panic(err)
	}
	schedule = sch
}
