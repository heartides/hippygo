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

func (c *Context) Success(data map[string]interface{}, msg ...string) {
	res := Response{
		Data: data,
		Code: 0,
	}
	if msg == nil {
		res.Msg = "OK"
	} else {
		res.Msg = msg[0]
	}
	c.JSON(http.StatusOK, res)
}

func (c *Context) Fail(msg string, errCode int, err ...string)  {
	res := Response{
		Code: errCode,
		Msg:  msg,
	}

	if err != nil {
		// 非DEV环境下隐藏ErrDesc
		if gin.Mode() != gin.ReleaseMode && err[0] != "" {
			res.ErrDesc = err[0]
		}
	}
	c.JSON(http.StatusOK, res)
}
