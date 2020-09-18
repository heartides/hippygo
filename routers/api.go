package routers

import (
	"github.com/gin-gonic/gin"
	"learning/app/controller"
	"learning/utils"
)

// 路由定义
func InitApiRouter() *gin.Engine {
	router := gin.Default()

	// CURD
	posts := router.Group("/post")
	{
		posts.POST("/create", utils.Handle(controller.Create))
		posts.PUT(":id/update", utils.Handle(controller.Update))
		posts.DELETE(":id/delete", utils.Handle(controller.Delete))
		posts.GET("/search", utils.Handle(controller.Search))
	}

	return router
}
