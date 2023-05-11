package controller

import (
	"blog/models"
	"blog/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
	// "strconv"
)

// Register 注册
//
//	{
//	  "Username": "余霞",
//	  "Password": "dolore",
//	  "Desc": "cillum cupidatat",
//	  "Gender": 2
//	}
func Register(c *gin.Context) {

	db := models.GetDB()
	var requestUser models.User
	c.Bind(&requestUser)
	username := requestUser.Username
	password := requestUser.Password
	gender := requestUser.Gender
	log.Println(username)
	log.Println(password)
	log.Println(gender)
	// 数据验证

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	models.DB.Where("username=?", username).First(&user)
	if user.Id != 0 {
		response.Response(c, http.StatusOK, false, nil, "用户名已存在")
		return
	}
	newUser := models.User{
		Username:  username,
		Password:  string(hashedPassword),
		Gender:    gender,
		Avatar:    "https://img.js.design/assets/img/645650aca30d747a6da0787b.jpg#208a42317334f6b46f2c03fcf9c101bd",
		IsAdmin:   false,
		Birthday:  time.Now(),
		LastLogin: time.Now(),
	}

	err = db.Create(&newUser).Error
	log.Println(&newUser.Id)
	if err != nil {
		response.Response(c, http.StatusOK, false, nil, "注册失败")
		return
	}
	// 返回结果
	response.Response(c, http.StatusOK, true, gin.H{"id": newUser.Id}, "注册成功")
}

func Login(c *gin.Context) {
	db := models.GetDB()

	var requestUser models.User
	c.Bind(&requestUser)
	username := requestUser.Username
	password := requestUser.Password
	var user models.User
	db.Where("username=?", username).First(&user)
	if user.Id == 0 {
		response.Response(c, http.StatusOK, false, nil, "用户不存在")
		return
	}
	//hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	//log.Println(hashedPassword)

	// Compare the hashed password with the stored hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(c, http.StatusOK, false, nil, "密码错误")
		return
	}

	response.Response(c, http.StatusOK, false, gin.H{"user": user}, "登录成功")
	return
}
