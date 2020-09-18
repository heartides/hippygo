package conf

import (
	"os"
)

type mysql struct {
	Host     string
	Database string
	Port     string //需要字符串拼接，不要设定为整型
	Username string
	Password string
	Charset  string
}

func ConfMysql() string {

	confMap := mysql{
		Host:     os.Getenv("DB_HOST"),
		Database: os.Getenv("DB_DATABASE"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Charset:  os.Getenv("DB_CHARSET"),
	}

	confMapString := confMap.Username

	// 非空密码
	if confMap.Password != "" {
		confMapString += ":" + confMap.Password
	}
	// 地址 + 端口
	confMapString += "@tcp(" + confMap.Host + ":" + confMap.Port + ")"

	confMapString += "/" + confMap.Database
	confMapString += "?charset=" + confMap.Charset + "&parseTime=True&loc=Local"

	return confMapString
}
