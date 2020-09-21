package controller

import (
	"github.com/gin-gonic/gin"
	"learning/conf"
	"learning/utils"
	"time"
)

var cache = conf.Cache

// 赋值
func RedisSet(cxt *utils.Context) {
	key := cxt.PostForm("key")
	_, err := cache.Set("key", key, time.Hour).Result()

	if err == nil {
		cxt.Success("写入成功", gin.H{"key": key})
		return
	}

	cxt.Fail(1, "写入失败", err)
}

// 取值
func RedisGet(cxt *utils.Context) {
	key := cxt.PostForm("key")
	value, err := cache.Get(key).Result()

	if err != nil {
		cxt.Fail(1, "key 不存在")
		return
	}
	cxt.Success("获取成功", gin.H{key: value})
}
