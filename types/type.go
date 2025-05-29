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
	AddTask() error           // 添加任务
	DeleteTask() error        // 删除任务
	UpdateTask() error        // 更新任务
	GetTask() error           // 获取任务
	GetAllTask() error        // 获取所有任务
	RunTask() error           // 运行任务
	StopTask() error          // 停止任务
	RestartTask() error       // 重启任务
	SaveTask() error          // 保存任务
	LoadTask() error          // 加载任务
	SaveAllTask() error       // 保存所有任务
	LoadAllTask() error       // 加载所有任务
	SaveTaskToFile() error    // 保存任务到文件
	LoadTaskFromFile() error  // 从文件加载任务
	SaveAllTaskToFile() error // 保存所有任务到文件
}
