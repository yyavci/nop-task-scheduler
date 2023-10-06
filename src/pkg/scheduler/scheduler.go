package scheduler

import (
	"fmt"
	"nop-task-scheduler/pkg/task"
	"time"

	"github.com/go-co-op/gocron"
)

func InitScheduler(tasks []task.ScheduleTask) (*gocron.Scheduler, error) {

	fmt.Println("initializing scheduler...")

	sch := gocron.NewScheduler(time.UTC)

	for i := 0; i < len(tasks); i++ {
		_, err := sch.Cron(tasks[i].CronExpression).Do(task.DoTask, tasks[i])
		if err != nil {
			fmt.Printf("Cannot run task! taskId:%d Err:%s\n", tasks[i].Id, err)
			return nil, err
		}
	}

	sch.StartBlocking()

	return sch, nil

}
