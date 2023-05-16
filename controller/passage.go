package controller

import (
	"blog/models"
	"blog/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"log"
	"net/http"
	"strconv"
	"time"
)

// GetAllPassages r.GET("/p",controller.GetAllPassages)
func GetAllPassages(c *gin.Context) {
	db := models.GetDB()
	var passages models.Passages
	err := db.Model(&models.Passage{}).Preload("Tags").Select([]string{"id", "title", "desc", "created_at", "cate_id"}).Find(&passages).Error
	if err != nil {
		response.Fail(c, "获取文章数据失败")
		return
	}
	response.Success(c, gin.H{"passages": passages}, "获取文章数据成功")
}

// GetPassageContent r.GET("/p/:id", controller.GetPassageContent)

type responsecomment struct {
	Id        uint      `json:"id"`
	Content   string    `json:"content"`
	UserId    uint      `json:"user_id"`
	Username  string    `json:"username"`
	Avatar    string    `json:"avatar"`
	PassageId uint      `json:"passage_id"`
	CreatedAt time.Time `json:"created_at"`
}

func GetPassageContent(c *gin.Context) {
	db := models.GetDB()
	PassageId := c.Param("id")
	var passage models.Passage
	//err := db.Find(&passage, PassageId)
	err := db.Model(&models.Passage{}).Preload("Tags").Preload("Comments").Find(&passage, PassageId).Error
	if err != nil {
		response.Response(c, http.StatusOK, false, nil, "获取文章内容失败")
		return
	}
	var responsecomments []responsecomment

	for i := 0; i < len(passage.Comments); i++ {
		log.Println(passage.Comments[i].UserId)
		var result models.User
		db.First(&result, passage.Comments[i].UserId)
		log.Println(result.Username)
		responsecomments = append(responsecomments, responsecomment{
			Id:        passage.Comments[i].Id,
			Content:   passage.Comments[i].Content,
			UserId:    passage.Comments[i].UserId,
			Username:  result.Username,
			Avatar:    result.Avatar,
			PassageId: passage.Comments[i].PassageId,
			CreatedAt: passage.Comments[i].CreatedAt,
		})
	}
	response.Response(c, http.StatusOK, true, gin.H{
		"passage": gin.H{
			"id":         passage.Id,
			"title":      passage.Title,
			"content":    passage.Content,
			"desc":       passage.Desc,
			"created_at": passage.CreatedAt,
			"comments":   responsecomments,
			"tags":       passage.Tags,
			"cate_id":    passage.CateId,
		},
	}, "获取文章内容成功")

}

// CreatePassage r.POST("/p",controller.CreatePassage)
func CreatePassage(c *gin.Context) {
	db := models.GetDB()
	var requestPassage models.Passage
	bindErr := c.ShouldBind(&requestPassage)
	if bindErr != nil {
		response.Response(c, http.StatusOK, false, gin.H{"error": bindErr}, "解析请求数据失败")
		log.Println(bindErr)
	}

	err := db.Create(&requestPassage).Error
	if err != nil {
		response.Response(c, http.StatusOK, false, gin.H{"error": err}, "添加文章失败")
		return
	}

	response.Response(c, http.StatusOK, true, gin.H{"id": requestPassage.Id}, "添加文章成功")
}

// DeletePassage r.DELETE("/p/:id",controller.DeletePassage)
func DeletePassage(c *gin.Context) {
	db := models.GetDB()
	PassageId := c.Param("id")
	pid, _ := strconv.Atoi(PassageId)
	passage := models.Passage{
		Id: uint(pid),
	}
	err := db.Select(clause.Associations).Delete(&passage).Error
	if err != nil {
		response.Response(c, http.StatusOK, false, gin.H{"error": err}, "删除文章失败")
		return
	}
	response.Response(c, http.StatusOK, true, nil, "删除文章成功")
	return
}

// UpdatePassage r.PUT("/p/:id",controller.UpdatePassage)

func UpdatePassage(c *gin.Context) {

	db := models.GetDB()
	PassageId := c.Param("id")
	var requestPassage models.Passage
	bindErr := c.ShouldBind(&requestPassage)
	if bindErr != nil {
		response.Response(c, http.StatusOK, false, nil, "解析请求数据失败")
		return
	}

	//err := db.Where("id=?", PassageId).Updates(&requestPassage).Error
	err := db.Model(&models.Passage{}).Where("id=?", PassageId).Updates(&requestPassage).Error
	if err != nil {
		log.Println(err)
		response.Response(c, http.StatusOK, false, nil, "文章信息更新失败")
		return
	}

	response.Response(c, http.StatusOK, true, nil, "文章信息更新成功")
	return
}
