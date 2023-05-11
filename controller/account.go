package controller

import (
	// "blog/models"
	// "blog/response"
	"github.com/gin-gonic/gin"
	// "golang.org/x/crypto/bcrypt"
	"net/http"
	"encoding/json"
	// "strconv"
)

// Register 注册
// {
//   "Username": "余霞",
//   "Password": "dolore",
//   "Desc": "cillum cupidatat",
//   "Gender": 2
// }
func Register(c *gin.Context) {
	// ,db *gorm.DB
	// db := common.GetDB()
	// 获取参数
	// var requestUser models.User
	
	b,_ := c.GetRawData()
	var newer map[string]interface{}
// 用interface{} 表示接收任意对象
// 返回值是一个接口
	_ = json.Unmarshal(b,&newer)
	// c.Bind(&requestUser)
	// username := newer.username
	// password := newer.password
	// gender	 := newer.gender
	// 数据验证
	// var user model.User
	// db.Where("username=?",username).First(&user)
	// if user.ID != 0 {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"code": 422,
	// 		"msg":  "用户已存在",
	// 	})
	// 	return
	// }
	
	// 创建用户
	// newUser := model.User{
	// 	Username: 	username,	
	// 	Password: 	password,
	// 	Avatar:		"/images/default_avatar.png"
	// 	IsAdmin:	false
	// }
	// err := db.Omit("Username", "Password", "Avatar","IsAdmin").Create(&newUser).Error
	// if err != nil {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"code": 500,
	// 		"msg":  "注册失败",
	// 	})
	// 	return
	// }
	// 返回结果
	c.JSON(http.StatusOK,
		gin.H{"status":  
		gin.H{"code": true,
		"msg":  "注册成功"},
		"data":	newer,
	})
}
