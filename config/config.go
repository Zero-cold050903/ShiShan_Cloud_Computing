package config

const DEFAULT_HOST = "127.0.0.1"
const MAX_TASK_NUM = 100
const DEFAULT_PORT = 8080

const RUNNING = 1
const STOPPED = 0
const WAITING = -1

var STATUS_CODE = map[string]int{
	"AVAILABLE": 1,
	"OCCUPIED":  -1,
	"RUNNING":   0,
}

const POOL_SIZE = 10

func InitConfig() {
	//init the config
}
