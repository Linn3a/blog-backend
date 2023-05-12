package controller

import (
	"blog/models"
	"blog/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Cateinfo struct {
	Id    int         `json:"id"`
	Name  string      `json:"name"`
	Cover string      `json:"cover"`
	Tags  models.Tags `json:"tags"`
}

// GetAllCates 	r.GET("/cate", controller.GetAllCates)
func GetAllCates(c *gin.Context) {
	db := models.GetDB()
	var cates Cateinfo
	err := db.Model(models.Cate{}).Preload("Tags").Find(&cates)
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
	//	TODO
}

// CreateCate 	r.POST("/cate", controller.CreateCate)
func CreateCate(c *gin.Context) {
	//	TODO
}

// UpdateCate 	r.PUT("/cate/:id", controller.UpdateCate)
func UpdateCate(c *gin.Context) {
	//	TODO
}

// DeleteCate 	r.DELETE("/cate/:id", controller.DeleteCate)
func DeleteCate(c *gin.Context) {
	//	TODO
}
