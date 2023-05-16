package controller

import (
	"blog/models"
	"blog/response"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func CreateStar(c *gin.Context) {
	db := models.GetDB()
	userId := c.Param("userId")
	passageId := c.Param("passageId")
	uid, _ := strconv.Atoi(userId)
	log.Println(uid)
	pid, _ := strconv.Atoi(passageId)
	log.Println(pid)
	user := models.User{
		Id: uint(uid),
	}
	log.Println(user)

	err := db.Model(&user).Association("Passages").Append(&models.Passage{Id: uint(pid)})
	if err != nil {
		response.Fail(c, "插入失败")
		return
	}
	response.Success(c, nil, "收藏文章成功")
}

func DeleteStar(c *gin.Context) {
	db := models.GetDB()
	userId := c.Param("userId")
	passageId := c.Param("passageId")
	uid, _ := strconv.Atoi(userId)
	log.Println(uid)
	pid, _ := strconv.Atoi(passageId)
	log.Println(pid)
	user := models.User{
		Id: uint(uid),
	}
	log.Println(user)
	passage := models.User{
		Id: uint(pid),
	}
	log.Println(user)
	err := db.Model(&user).Association("Passages").Delete(&passage)
	if err != nil {
		response.Fail(c, "删除失败")
		return
	}
	response.Success(c, nil, "清除收藏成功")
}
