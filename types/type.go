package types

import (
	"github.com/Zero_cold050903/ShiShan_Cloud_Computing/config"
)

type User struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Permission bool   `json:"permission"`
}
type Task struct {
	TaskID     string `json:"taskid"`
	TaskName   string `json:"taskname"`
	TaskType   string `json:"tasktype"`
	TaskStatus string `json:"taskstatus"`
	TaskFile   string `json:"taskfile"` // 任务文件路径
}
type TaskQueue struct {
	Comming  Task
	Runnning Task
	Finished Task
	Queue    [config.MAX_TASK_NUM]Task
}
type TaskHandler interface {
	AddTask(task TaskQueue) error           // 添加任务
	DeleteTask(task TaskQueue) error        // 删除任务
	UpdateTask(task TaskQueue) error        // 更新任务
	GetTask(task TaskQueue) error           // 获取任务
	GetAllTask(task TaskQueue) error        // 获取所有任务
	RunTask(task TaskQueue) error           // 运行任务
	StopTask(task TaskQueue) error          // 停止任务
	RestartTask(task TaskQueue) error       // 重启任务
	SaveTask(task TaskQueue) error          // 保存任务
	LoadTask(task TaskQueue) error          // 加载任务
	SaveAllTask(task TaskQueue) error       // 保存所有任务
	LoadAllTask(task TaskQueue) error       // 加载所有任务
	SaveTaskToFile(task TaskQueue) error    // 保存任务到文件
	LoadTaskFromFile(task TaskQueue) error  // 从文件加载任务
	SaveAllTaskToFile(task TaskQueue) error // 保存所有任务到文件
}
