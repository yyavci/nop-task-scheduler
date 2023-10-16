package scheduler

import (
	"fmt"
	"time"

	"github.com/yyavci/nop-task-scheduler/internal/task"
	conf "github.com/yyavci/nop-task-scheduler/internal/config"

	"github.com/go-co-op/gocron"
)

func InitScheduler(tasks []task.ScheduleTask , config conf.AppConfig) (*gocron.Scheduler, error) {

	fmt.Println("initializing scheduler...")

	sch := gocron.NewScheduler(time.UTC)

	for i := 0; i < len(tasks); i++ {
		_, err := sch.Cron(tasks[i].CronExpression).Do(task.DoTask, tasks[i] , config)
		if err != nil {
			fmt.Printf("Cannot run task! taskId:%d Err:%+v\n", tasks[i].Id, err)
			return nil, err
		}
	}

	sch.StartBlocking()

	return sch, nil

}
