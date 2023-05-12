package controller

import (
	"blog/models"
	"blog/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// GetPassageContent r.GET("/p/:id", controller.GetPassageContent)
func GetPassageContent(c *gin.Context) {
	db := models.GetDB()
	PassageId := c.Param("id")
	var passage models.Passage
	err := db.Find(&passage, PassageId)
	if err != nil {
		response.Response(c, http.StatusOK, false, gin.H{"error": err}, "获取文章内容失败")
		return
	}
	var tagsId []uint
	err = db.Model(&models.Tag{}).Preload("Passages").Find(&tagsId)
	//err = db.Model(models.TagPassages{}).Find(&tagsId, "PassageId=?", PassageId)
	if err != nil {
		response.Response(c, http.StatusOK, false, gin.H{"error": err}, "获取文章标签失败")
		return
	}

	var comments models.Comments
	err = db.Find(&comments, "PassageId=?", PassageId)
	if err != nil {
		response.Response(c, http.StatusOK, false, gin.H{"error": err}, "获取文章评论失败")
		return
	}

	response.Response(c, http.StatusOK, true, gin.H{
		"id":        passage.Id,
		"title":     passage.Title,
		"content":   passage.Content,
		"desc":      passage.Desc,
		"create_at": passage.CreatedAt,
		"cate_id":   passage.CateId,
		"tags":      tagsId,
		"comments":  comments,
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

	passageTitle := requestPassage.Title
	passageContent := requestPassage.Content
	passageDesc := requestPassage.Desc
	passageCateId := requestPassage.CateId

	newPassage := models.Passage{
		Title:   passageTitle,
		Content: passageContent,
		Desc:    passageDesc,
		CateId:  passageCateId,
	}

	err := db.Create(&newPassage).Error
	log.Println(&newPassage.Id)
	if err != nil {
		response.Response(c, http.StatusOK, false, gin.H{"error": err}, "添加文章失败")
		return
	}

	response.Response(c, http.StatusOK, true, gin.H{"id": newPassage.Id}, "添加文章成功")
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
type updatedPassage struct {
	Passage models.Passage `json:"passage"`
	Tags    []uint         `json:"tags"`
}

func UpdatePassage(c *gin.Context) {

	db := models.GetDB()
	PassageId := c.Param("id")
	var requestPassage updatedPassage
	bindErr := c.ShouldBind(&requestPassage)
	if bindErr != nil {
		response.Response(c, http.StatusOK, false, nil, "解析请求数据失败")
		return
	}
	log.Print(requestPassage)
	//tags := requestPassage.Tags
	newPassage := requestPassage.Passage
	log.Println(newPassage)
	//if bindErr != nil {
	//	response.Response(c, http.StatusOK, false, nil, "解析请求数据失败")
	//	return
	//}
	err := db.Model(&models.Passage{}).Where("Id=?", PassageId).Updates(newPassage).Error
	if err != nil {
		response.Response(c, http.StatusOK, false, nil, "文章信息更新失败")
		return
	}
	//var tagedpassages []models.TagPassage
	//passageId, _ := strconv.Atoi(PassageId)
	//for _, value := range tags {
	//	tagedpassages = append(tagedpassages, models.TagPassage{
	//		PassageId: passageId,
	//		TagId:     value,
	//	})
	//}
	//err = db.Save(&tagedpassages).Error
	//if err != nil {
	//	response.Response(c, http.StatusOK, false, nil, "文章标签插入失败")
	//	return
	//}
	response.Response(c, http.StatusOK, true, nil, "文章信息更新成功")
	return
}
