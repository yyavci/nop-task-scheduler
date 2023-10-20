# NopCommerce Task Scheduler
## Introduction
this is my *"hello world"* app for go language. (for nopcommerce 4.2. might work for others too 😀)

## Instructions
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
