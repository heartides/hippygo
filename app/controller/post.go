package controller

import "github.com/gin-gonic/gin"

func Post(c *gin.Context)  {
	c.String(200,"hello gin")
}