package controller

import (
	"blog/models"
	"blog/response"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// GetAllTags r.GET("/tag", controller.GetAllTags)

type responseTag struct {
	Id            uint   `json:"id"`
	Name          string `json:"name"`
	Color         string `json:"color"`
	CateId        uint   `json:"cate_id"`
	PassageAmount int    `json:"passage_amount"`
}

// GetAllTags r.GET("/tag", controller.GetAllTags)
func GetAllTags(c *gin.Context) {
	db := models.GetDB()
	var tags models.Tags
	err := db.Model(&models.Tag{}).Preload("Passages").Find(&tags).Error
	if err != nil {
		response.Fail(c, "获取标签数据失败")
		return
	}
	var responseBody []responseTag
	for i := 0; i < len(tags); i++ {
		responseBody = append(responseBody, responseTag{
			Id:            tags[i].Id,
			Name:          tags[i].Name,
			Color:         tags[i].Color,
			CateId:        tags[i].CateId,
			PassageAmount: len(tags[i].Passages),
		})
	}

	response.Success(c, gin.H{"tags": responseBody}, "获取所有标签成功")
}

// GetTag r.GET("/tag/:id", controller.GetTag)
func GetTag(c *gin.Context) {
	db := models.GetDB()
	tagId := c.Param("id")
	tid, _ := strconv.Atoi(tagId)
	var tag models.Tag
	err := db.Model(&models.Tag{}).Preload("Passages").Preload("Passages.Tags").Find(&tag, uint(tid)).Error
	if err != nil {
		log.Println(err)
		response.Fail(c, "查找标签数据失败")
		return
	}

	for i := 0; i < len(tag.Passages); i++ {
		tag.Passages[i].Content = ""
	}
	log.Println(tag)
	response.Success(c, gin.H{"tag": tag}, "获取标签数据成功")
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
		return
	}
	response.Success(c, nil, "删除标签成功")
}
