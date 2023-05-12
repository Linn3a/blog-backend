package controller

import (
	"blog/models"
	"blog/response"
	"github.com/gin-gonic/gin"
)

// GetAllTags r.GET("/tag", controller.GetAllTags)
func GetAllTags(c *gin.Context) {
	db := models.GetDB()
	var tags models.Tags
	err := db.Find(&tags).Error
	if err != nil {
		response.Fail(c, "获取标签数据失败")
	}
	response.Success(c, gin.H{"tags": tags}, "获取所有标签成功")
}

// GetTag r.GET("/tag/:id", controller.GetTag)
func GetTag(c *gin.Context) {
	//	TODO
}

// CreateTag r.POST("/tag", controller.CreateTag)
func CreateTag(c *gin.Context) {
	db := models.GetDB()
	var requestTag models.Tag
	err := c.ShouldBind(&requestTag)
	if err != nil {
		response.Fail(c, "解析请求数据失败")
		return
	}
	err = db.Create(&requestTag).Error
	if err != nil {
		response.Fail(c, "新建标签失败")
		return
	}
	response.Success(c, gin.H{"id": requestTag.Id}, "新建标签成功")
}

// DeleteTag r.DELETE("/tag/:id",controller.DeleteTag)
func DeleteTag(c *gin.Context) {
	db := models.GetDB()
	tagId := c.Param("id")
	err := db.Delete(&models.Tag{}, tagId).Error
	if err != nil {
		response.Fail(c, "删除标签失败")
	}
	response.Success(c, nil, "删除标签成功")
}
