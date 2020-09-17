package databases

import (
	"github.com/jinzhu/gorm"
	"learning/utils"
	"time"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

// 初始化Mysql
func InitMysql(connSetting string)  {
	db,err := gorm.Open("mysql",connSetting)
	db.LogMode(true)
	if err != nil{
		utils.Log().Panic("数据库连接失败",err)
	}

	// 空闲
	db.DB().SetMaxIdleConns(50)

	// 打开
	db.DB().SetMaxOpenConns(100)

	// 超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db

	// 运行迁移
	migration()

}