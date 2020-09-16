package controllers

import (
	"github.com/gin-gonic/gin"
)

type post struct {
	Name  string `json:"name" binding:"required"`
	Title string `json:"title" binding:"required"`
	Price int    `json:"price" binding:"required"`
}

func Form(c *gin.Context) {
	var (
		postForm post
		response gin.H
	)

	if err := c.ShouldBindJSON(&postForm); err == nil {
		response = gin.H{"code": 0, "msg": postForm}
	} else {
		response = gin.H{"code": 1, "msg": err.Error()}
	}

	c.JSON(200, response)
}
