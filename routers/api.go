package routers

import (
	"github.com/gin-gonic/gin"
	"learning/app/controller"
	"learning/utils"
)

// controller 路由定义
func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", utils.Handle(controller.Post))
	return router
}
