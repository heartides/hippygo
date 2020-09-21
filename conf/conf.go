package conf

import (
	"github.com/joho/godotenv"
	"learning/databases"
	"log"
	"os"
)

func init() {
	// 读取环境变量
	godotenv.Load()

	//连接数据库
	if os.Getenv("DB_CONNON") == "ON" {
		databases.InitMysql(ConfMysql())
	}

	// 连接Redis
	if os.Getenv("REDIS_CONNON") == "ON" {
		_, redisConnErr := ConfRedis().Ping().Result()
		if redisConnErr != nil {
			log.Panic("连接Redis不成功", redisConnErr.Error())
		}
	}

}
