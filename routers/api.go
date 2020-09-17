package routers

import (
	"github.com/gin-gonic/gin"
	"learning/app/controller"
	"learning/utils"
)

// controller 路由定义
func init() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/", utils.Handle(controller.Post))
	}


	router.Run(":3030")
}
