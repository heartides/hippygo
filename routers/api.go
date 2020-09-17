package routers

import (
	"github.com/gin-gonic/gin"
	"learning/app/controller"
)

// controller 路由定义
func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", controller.Post)
	return router
}
