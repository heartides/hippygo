package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"path"
	"strings"
	"time"
)

// 上传处理
// 单文件处理
func Uploads(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(200, gin.H{"code": 1, "msg": err})
		return
	}
	//获取文件后缀
	filenameWithSuffix := path.Ext(file.Filename)

	//定义文件名
	fileDstFullName := GetRandomString(16) + filenameWithSuffix

	//上传处理
	fileError := c.SaveUploadedFile(file, strings.Join([]string{"uploads", fileDstFullName}, "/"))

	if fileError != nil {
		log.Println(fileError)
		c.JSON(200, gin.H{"code": 1, "msg": fileError.Error()})
		return
	}
	c.JSON(200, gin.H{"code": 0, "msg": "ok"})
}

//随机字符串
func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
