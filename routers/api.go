package routers

import (
	"github.com/gin-gonic/gin"
	"learning/controllers"
)

// api 路由定义
func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", controllers.Index)
	router.POST("/upload", controllers.Uploads)
	router.POST("/form",controllers.Form)
	return router
}
