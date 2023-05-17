package controller

import (
	"blog/models"
	"blog/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Cateinfo struct {
	Id    int    `json:"id" binging:"-"`
	Name  string `json:"name"`
	Cover string `json:"cover"`
}

// GetAllCates 	r.GET("/cate", controller.GetAllCates)
func GetAllCates(c *gin.Context) {
	db := models.GetDB()
	var cates []models.Cate
	err := db.Model(models.Cate{}).Preload("Tags").Find(&cates).Error
	if err != nil {
		response.Response(c, http.StatusOK, false, nil, "读取数据库失败")
		return
	}
	log.Println(cates)
	response.Response(c, http.StatusOK, true, gin.H{
		"cates": cates,
	}, "获取成功")
	return
}

// GetCate 	r.GET("/cate/:id", controller.GetCate)
func GetCate(c *gin.Context) {
	db := models.GetDB()
	cateid := c.Param("id")
	var cate models.Cate
	err := db.Model(models.Cate{}).Preload("Passages").Preload("Passages.Tags").Preload("Tags").Find(&cate, cateid).Error
	if err != nil {
		response.Response(c, http.StatusOK, false, nil, "读取数据库失败")
		return
	}
	log.Println(cate)
	for i := 0; i < len(cate.Passages); i++ {
		cate.Passages[i].Content = ""
	}
	response.Response(c, http.StatusOK, true, gin.H{
		"cate": cate,
	}, "获取成功")
	return

}

// CreateCate 	r.POST("/cate", controller.CreateCate)
func CreateCate(c *gin.Context) {
	db := models.GetDB()
	var requestCate models.Cate
	err := c.ShouldBind(&requestCate)
	if err != nil {
		response.Response(c, http.StatusOK, false, nil, "解析请求数据失败")
	}
	catename := requestCate.Name
	var cate models.Cate
	models.DB.Where("name=?", catename).First(&cate)
	if cate.Id != 0 {
		response.Response(c, http.StatusOK, false, nil, "该类已存在")
		return
	}

	err = db.Create(&requestCate).Error
	log.Println(&requestCate.Id)
	if err != nil {
		response.Response(c, http.StatusOK, false, nil, "添加新类别失败")
		return
	}
	response.Response(c, http.StatusOK, true, gin.H{"id": requestCate.Id}, "添加新类别成功")
}

// UpdateCate 	r.PUT("/cate/:id", controller.UpdateCate)
func UpdateCate(c *gin.Context) {
	db := models.GetDB()
	cate_id := c.Param("id")
	var requestCate models.Cate
	err := c.ShouldBind(&requestCate)
	if err != nil {
		response.Response(c, http.StatusOK, false, nil, "解析请求数据失败")
		return
	}
	err = db.Model(&models.Cate{}).Where("id=?", cate_id).Updates(&requestCate).Error
	if err != nil {
		response.Response(c, http.StatusOK, false, nil, "更新数据库失败")
		log.Println(err)
		return
	}
	response.Response(c, http.StatusOK, true, nil, "更新类别数据成功")

}

// DeleteCate 	r.DELETE("/cate/:id", controller.DeleteCate)
func DeleteCate(c *gin.Context) {
	db := models.GetDB()
	cateid := c.Param("id")
	err := db.Delete(&models.Cate{}, cateid).Error
	if err != nil {
		response.Response(c, http.StatusOK, false, nil, "删除类别失败")
		return
	}
	response.Response(c, http.StatusOK, true, nil, "删除类别成功")
	return
}
