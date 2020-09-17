package conf

import (
	"github.com/joho/godotenv"
	"learning/databases"
	"learning/utils"
	"os"
)

func init()  {
	// 读取环境变量
	godotenv.Load()

	//设置日志级别
	utils.BuildLogger(os.Getenv("LOG_LEVEL"))

	//连接数据库
	databases.InitMysql(os.Getenv("MYSQL_DSN"))

}
