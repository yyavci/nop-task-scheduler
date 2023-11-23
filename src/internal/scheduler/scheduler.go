package scheduler

import (
	"errors"
	"fmt"
	"time"

	"github.com/yyavci/nop-task-scheduler/internal/store"
	"github.com/yyavci/nop-task-scheduler/internal/task"

	"github.com/go-co-op/gocron"
)

// errors
var errCannotRunTask = errors.New("cannot run task")

func InitScheduler(tasks []task.ScheduleTask, stores []store.Store) (*gocron.Scheduler, error) {

	fmt.Println("initializing scheduler...")

	sch := gocron.NewScheduler(time.UTC)

	for y := 0; y < len(stores); y++ {

		for i := 0; i < len(tasks); i++ {
			_, err := sch.Cron(tasks[i].CronExpression).Do(task.DoTask, tasks[i], stores[y])
			if err != nil {
				fmt.Printf("error: cannot run task! taskId:%d err:%+v\n %+v\n", tasks[i].Id, errCannotRunTask, err)
				return nil, errCannotRunTask
			}
		}

	}

	sch.StartBlocking()

	return sch, nil

}
