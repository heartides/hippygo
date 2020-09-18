package controller

import (
	"github.com/gin-gonic/gin"
	"learning/app/model"
	"learning/databases"
	"learning/utils"
)

var (
	db = databases.DB
)

// 增
func Create(cxt *utils.Context) {
	title := cxt.PostForm("title")
	post := model.Posts{
		Title: title,
	}
	db.Create(&post)
	cxt.Success("新增成功", gin.H{"title": title})
}

// 改
func Update(cxt *utils.Context) {
	title := cxt.PostForm("title")
	id := cxt.Param("id")
	db.Model(model.Posts{}).Where("id = ?", id).Update("title", title)
	cxt.Success("修改成功", gin.H{"title": title, "id": id})
}

// 删
func Delete(cxt *utils.Context) {
	id := cxt.Param("id")
	// 软删
	db.Where("id = ?", id).Delete(&model.Posts{})

	// 永久删
	// db.Unscoped().Where("id = ?",id).Delete(&model.Posts{})
	cxt.Success("删除成功", gin.H{"id": id})
}

// 查
func Search(cxt *utils.Context) {
	keyword := cxt.Query("keyword")
	result := db.Where("title like ?", "%"+keyword+"%").Find(&model.Posts{})
	if result.RecordNotFound() {
		cxt.Fail(1, "语法或系统发生错误", result.Error.Error())
		return
	}
	cxt.Success("查询结果", gin.H{"result": result.Value})
}
