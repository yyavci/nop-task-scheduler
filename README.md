# NopCommerce Task Scheduler
## What is it?
Nopcommerce does only support timer based tasks like every n seconds. This app will fix that missing feature.

It is working fine using Nopcommerce 4.2 version. should function for other versions too. (might need small modifications)

This is my *"hello world"* app for go language.

## Setup
- Set your *"CONN_STR"* environment variable or *"ConnectionString"* field in your config.json file.
- Add *"CronExpression"* column to your *"ScheduleTask"* table.
```sql
ALTER TABLE [ScheduleTask] 
	ADD [CronExpression] nvarchar(512) NULL;
```
- Add return statement at the beginning of your *"Nop.Services/Tasks/TaskManager.Initialize()"* method like below. (or make it configurable)
```csharp

public void Initialize()
{
    //here
    return;

    _taskThreads.Clear();

    var taskService = EngineContext.Current.Resolve<IScheduleTaskService>();
    
    .
    .
    .
}
```
## Build & Run

You can use following command in your terminal.
```shell
# build app
go build ./cmd/task-scheduler/

# run
./task-scheduler
```