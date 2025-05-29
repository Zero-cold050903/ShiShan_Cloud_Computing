package log
import (
	"log"
	"os"
)
type LogHandler interface {
	LogInfo(message string) error // 记录信息
	LogError(message string) error // 记录错误
	LogWarning(message string) error // 记录警告
	LogDebug(message string) error // 记录调试信息
}
type Log struct {
	LogID   string `json:"logid"`
	LogError string `json:"logerror"`
	LogTimeStamp string `json:"logtimestamp"`
	LogLevel string `json:"loglevel"`
	LogMessage string `json:"logmessage"`
}
func distributeID (l *Log) {
	// 生成一个随机的ID
	randomID := 
}
func (l *log) LogInfo(message string) error {
	log.Println(message)
	return nil
}
