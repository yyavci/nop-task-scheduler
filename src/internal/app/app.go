package app

import (
	"fmt"

	"github.com/yyavci/nop-task-scheduler/internal/config"
	"github.com/yyavci/nop-task-scheduler/internal/database"
	"github.com/yyavci/nop-task-scheduler/internal/scheduler"
	"github.com/yyavci/nop-task-scheduler/internal/store"
	"github.com/yyavci/nop-task-scheduler/internal/task"
)

func Run() {
	fmt.Println("app started")

	conf, err := config.ReadConfiguration("config.json")
	if err != nil {
		return
	}

	err = database.Init(*conf)
	if err != nil {
		return
	}

	tasks, err := task.GetScheduleTasks()
	if err != nil {
		return
	}

	stores, err := store.GetStores()
	if err != nil {
		return
	}

	_, err = scheduler.InitScheduler(tasks, stores)
	if err != nil {
		return
	}
}
