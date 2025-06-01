package cmd
import(
	"os"
	"github.com/Zero_cold050903/ShiShan_Cloud_Computing/config"
	"github.com/Zero_cold050903/ShiShan_Cloud_Computing/log"
	"github.com/Zero_cold050903/ShiShan_Cloud_Computing/types"
	"github.com/Zero_cold050903/ShiShan_Cloud_Computing/db"
	"github.com/Zero_cold050903/ShiShan_Cloud_Computing/cmd/api"
	"github.com/gin-gonic/gin"
)
func init(){
	db.InitDB() // 初始化数据库
	gin.SetMode(gin.ReleaseMode) // 设置为生产模式
	// 初始化日志
	log.InitLog()
	// 初始化配置
	config.InitConfig()
	// 初始化任务队列
	types.InitTaskQueue()
}
func main() {

}
