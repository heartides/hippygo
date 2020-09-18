package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//实现 Context 接口
type Context struct {
	*gin.Context
}

type HandlerFunc func(cxt *Context)

func Handle(handle HandlerFunc) gin.HandlerFunc {
	return func(cxt *gin.Context) {
		ctx := &Context{
			cxt,
		}
		handle(ctx)
	}
}

// 成功响应
func (c *Context) Success(msg string, data map[string]interface{}) {
	res := Response{
		Data: data,
		Msg:  msg,
		Code: 0,
	}
	c.JSON(http.StatusOK, res)
}

// 失败响应
func (c *Context) Fail(errCode int, msg string, err ...interface{}) {
	res := Response{
		Code: errCode,
		Msg:  msg,
	}

	if err != nil {
		// 非DEV环境下隐藏ErrDesc
		if gin.Mode() != gin.ReleaseMode && err[0] != "" {
			res.ErrDesc = err
		}
	}
	c.JSON(http.StatusOK, res)
}
