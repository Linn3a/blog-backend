package controller

import (
	"blog/models"
	"blog/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// GetAllPassages r.GET("/p",controller.GetAllPassages)
func GetAllPassages(c *gin.Context) {
	db := models.GetDB()
	var passages models.Passages
	err := db.Select([]string{"id", "title", "desc", "created_at", "cate_id"}).Find(&passages).Error
	if err != nil {
		response.Fail(c, "获取文章数据失败")
		return
	}
	response.Success(c, gin.H{"passages": passages}, "获取文章数据成功")
}

// GetPassageContent r.GET("/p/:id", controller.GetPassageContent)
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

	response.Response(c, http.StatusOK, true, gin.H{
		"passage": passage,
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
	log.Println(requestPassage)

	err := db.Create(&requestPassage).Error
	log.Println(&requestPassage.Id)
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
	err := db.Delete(&models.Passage{}, PassageId).Error
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
