package log

import (
	"fmt"
	"log"
	"os"
	"time"
)

type LogHandler interface {
	LogInfo(message string) error    // 记录信息
	LogError(message string) error   // 记录错误
	LogWarning(message string) error // 记录警告
	LogDebug(message string) error   // 记录调试信息
}
type Log struct {
	LogID        string `json:"logid"`
	LogError     string `json:"logerror"`
	LogTimeStamp string `json:"logtimestamp"`
	LogLevel     string `json:"loglevel"`
	LogMessage   string `json:"logmessage"`
}

func InitLog() {
	homedir := os.Getenv("HOME")
	err := os.Mkdir(homedir + "/log")
	if err != nil {
		log.Fatal(err)
		os.WriteFile(homedir+"/log/error.log", []byte(err.Error()), 0644)
	}
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	fmt.Println(formattedTime + " Log directory initialized\n")
}
func (l *Log) LogInfo(message string) error {
	log.Println(message)
	return nil
}
