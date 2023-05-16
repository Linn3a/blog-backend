package controller

import (
	"blog/models"
	"blog/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetCommentsbyPassageId(c *gin.Context) {
	db := models.GetDB()
	passageId := c.Param("id")
	pid, _ := strconv.Atoi(passageId)
	var comments models.Comments
	err := db.Where("passage_id=?", uint(pid)).Find(&comments).Error
	if err != nil {
		response.Fail(c, "获取评论数据失败")
	}
	response.Success(c, gin.H{"comments": comments}, "获取评论数据成功")
}

// CreateComment r.POST("/comment",CreateComment)
func CreateComment(c *gin.Context) {
	db := models.GetDB()
	var requestComment models.Comment
	err := c.ShouldBind(&requestComment)
	if err != nil {
		response.Fail(c, "解析请求数据失败")
	}
	err = db.Create(&requestComment).Error
	if err != nil {
		response.Fail(c, "新增评论失败")
	}
	response.Success(c, gin.H{"id": requestComment.Id}, "新增评论成功")
}

// DeleteComment r.DELETE("/comment",DeleteComment)
func DeleteComment(c *gin.Context) {
	db := models.GetDB()
	commentId := c.Param("id")
	err := db.Delete(&models.Comment{}, commentId).Error
	if err != nil {
		response.Fail(c, "删除评论失败")
	}
	response.Success(c, nil, "删除评论成功")
}
