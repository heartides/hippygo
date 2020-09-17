package controller

import (
	"learning/utils"
)

func Post(cxt *utils.Context)  {
	cxt.Fail("登录失败",422)
}