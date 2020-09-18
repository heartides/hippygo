package middleware

import (
	"github.com/gin-gonic/gin"
)

func Example() gin.HandlerFunc {
	return func(cxt *gin.Context) {
		cxt.Next()
	}
}
