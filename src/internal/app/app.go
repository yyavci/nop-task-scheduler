package app

import (
	"fmt"

	"github.com/yyavci/nop-task-scheduler/internal/config"
	"github.com/yyavci/nop-task-scheduler/internal/customer"
	"github.com/yyavci/nop-task-scheduler/internal/database"
	"github.com/yyavci/nop-task-scheduler/internal/scheduler"
	"github.com/yyavci/nop-task-scheduler/internal/task"

	"github.com/go-co-op/gocron"
)

var schedule *gocron.Scheduler

func Run() {
	fmt.Println("app started")

	conf, err := config.ReadConfiguration()
	if err != nil {
		panic(err)
	}

	err = database.Init(*conf)
	if err != nil {
		panic(err)
	}

	cust, err := customer.GetScheduleTaskCustomer("BackgroundTask")
	if err != nil {
		panic(err)
	}
	fmt.Printf("schedule task customerId:%d\n", cust.Id)

	tasks, err := task.GetScheduleTasks()
	if err != nil {
		panic(err)
	}

	sch, err := scheduler.InitScheduler(tasks, *conf)
	if err != nil {
		panic(err)
	}
	schedule = sch
}
