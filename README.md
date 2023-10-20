# nopcommerce task scheduler
## introduction
this is my *"hello world"* app for go language. (for nopcommerce 4.2. might work for others too 😀)

## instructions
- set your *"CONN_STR"* environment variable or *"ConnectionString"* field and *"StoreUrl"* in your config.json file.
- add *"CronExpression"* column to your *"ScheduleTask"* table.
- return at the beginning of your *"Nop.Services/Tasks/TaskManager.Initialize()"* method like below. (or make it configurable)
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
- create new controller method routed at *"/ScheduleTask/Run"* with a parameter *TaskId* which calls something like *"Nop.Web/Areas/Admin/Controllers/ScheduleTaskController.RunNow()"* and returns *"ScheduleTaskResponse"* as json.

```csharp
ScheduleTaskResponse Run(int TaskId);
```
```csharp
ScheduleTaskResponse
{
  public bool Success {get; set;}
  public string Message {get; set;}
}
```
