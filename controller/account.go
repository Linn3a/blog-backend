package controller

import (
	"blog/models"
	"blog/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/clause"
	"log"
	"net/http"
	"strconv"
	"time"
	// "strconv"
)

// Register 用户注册：增加一个用户
func Register(c *gin.Context) {

	db := models.GetDB()
	var requestUser models.User
	bindErr := c.ShouldBind(&requestUser)
	if bindErr != nil {
		log.Println(bindErr)
		response.Response(c, http.StatusOK, false, nil, "解析请求数据失败")
		return
	}
	log.Println(requestUser)
	username := requestUser.Username
	password := requestUser.Password
	gender := requestUser.Gender

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

	response.Response(c, http.StatusOK, true, gin.H{"id": newUser.Id}, "注册成功")
}

// Login 用户登录
func Login(c *gin.Context) {
	db := models.GetDB()

	var requestUser models.User
	bindErr := c.Bind(&requestUser)
	if bindErr != nil {
		response.Fail(c, "解析请求数据失败")
		return
	}
	username := requestUser.Username
	password := requestUser.Password
	var user models.User
	db.Where("username=?", username).First(&user)
	if user.Id == 0 {
		response.Fail(c, "用户不存在")
		return
	}

	// Compare the hashed password with the stored hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Fail(c, "密码错误")
		return
	}
	Token, err := GenToken(user.Id, user.Username, user.Avatar)
	if err != nil {
		response.Fail(c, "生成token失败")
	}
	response.Success(c, gin.H{
		"user":  user,
		"Token": Token,
	}, "登录成功")
	return
}

func DeleteUser(c *gin.Context) {
	db := models.GetDB()
	userid := c.Param("id")
	uid, _ := strconv.Atoi(userid)
	//err := db.Where("user_id = ?", uint(uid)).Delete(&models.Comment{}).Error
	//if err != nil {
	//	response.Fail(c, "删除该用户发出评论失败")
	//	return
	//}
	user := models.User{
		Id: uint(uid),
	}
	err := db.Select(clause.Associations).Delete(&user).Error
	if err != nil {
		response.Response(c, http.StatusOK, false, nil, "注销用户失败")
		return
	}
	response.Response(c, http.StatusOK, true, nil, "注销成功")
	return
}

func UpdateUser(c *gin.Context) {
	db := models.GetDB()
	userid := c.Param("id")
	var requestUser models.User
	bindErr := c.Bind(&requestUser)
	if bindErr != nil {
		response.Response(c, http.StatusOK, false, nil, "解析请求数据失败")
		return
	}
	log.Println(requestUser)
	if requestUser.Password != "" {
		hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(requestUser.Password), bcrypt.DefaultCost)
		if err != nil {
			response.Response(c, http.StatusOK, false, nil, "密码加密失败")
			return
		}
		requestUser.Password = string(hashedNewPassword)
		err = db.Model(&models.User{}).Where("Id=?", userid).Updates(requestUser).Error
		if err != nil {
			response.Response(c, http.StatusOK, false, nil, "用户信息更新失败")
			return
		}
		response.Response(c, http.StatusOK, true, nil, "用户信息更新成功")
		return
	}

	err := db.Model(&models.User{}).Where("Id=?", userid).Updates(requestUser).Error
	if err != nil {
		response.Response(c, http.StatusOK, false, nil, "用户信息更新失败")
		return
	}
	response.Response(c, http.StatusOK, true, nil, "用户信息更新成功")
	return
}

// 由于gorm存在false字段拒绝更新的特性 把更改管理员权限单独拿出来写
type adminState struct {
	IsAdmin bool `json:"is_admin"`
}

func ChangeAdminState(c *gin.Context) {
	db := models.GetDB()
	userid := c.Param("id")
	var requestState adminState
	bindErr := c.ShouldBind(&requestState)
	if bindErr != nil {
		response.Response(c, http.StatusOK, false, nil, "解析请求数据失败")
		return
	}
	isAdmin := requestState.IsAdmin
	err := db.Model(&models.User{}).Where("Id=?", userid).Updates(map[string]interface{}{"IsAdmin": isAdmin}).Error
	if err != nil {
		response.Response(c, http.StatusOK, false, nil, "用户管理权限更新失败")
		return
	}
	response.Response(c, http.StatusOK, true, nil, "用户管理权限更新成功")
}

type userinfo struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	Gender    int    `json:"gender"`
	LastLogin string `json:"last_login"`
	CreatedAt string `json:"created_at"`
	IsAdmin   bool   `json:"is_admin"`
}

func GetAllUsers(c *gin.Context) {
	db := models.GetDB()
	var users []userinfo
	err := db.Model(models.User{}).Find(&users).Error
	if err != nil {
		response.Response(c, http.StatusOK, false, nil, "获取用户信息失败")
		return
	}
	log.Println(users)
	response.Response(c, http.StatusOK, true, gin.H{"users": users}, "获取成功")
	return
}

func GetUser(c *gin.Context) {
	db := models.GetDB()
	userid := c.Param("id")
	var user models.User
	err := db.Model(&models.User{}).Preload("Passages").Preload("Passages.Tags").Preload("Comments").Find(&user, userid).Error
	if err != nil {
		response.Response(c, http.StatusOK, false, nil, "获取用户信息失败")
		return
	}
	//var stars models.Stars
	//err = db.Find(&stars, "UserId=?", userid)
	//if err != nil {
	//	response.Response(c, http.StatusOK, false, nil, "获取用户收藏文章失败")
	//	return
	//}
	//log.Println(stars)
	//var comments models.Comments
	//err = db.Find(&comments, "UserId=?", userid)
	//if err != nil {
	//	response.Response(c, http.StatusOK, false, nil, "获取用户发出评论失败")
	//	return
	//}
	//log.Println(comments)
	response.Response(c, http.StatusOK, true, gin.H{
		"user": user,
	}, "获取用户信息成功")

}

type requsettoken struct {
	Token string
}

func Autologin(c *gin.Context) {
	var currentToken requsettoken
	err := c.ShouldBind(&currentToken)
	if err != nil {
		response.Fail(c, "解析请求数据失败")
		log.Println(err)
		return
	}
	currentUser, err := ParseToken(currentToken.Token)

	if err != nil {
		log.Println(err)
		log.Println(currentUser)
		response.Fail(c, "解析Token失败")
		return
	}
	response.Success(c, gin.H{
		"id":       currentUser.UserID,
		"username": currentUser.Username,
		"avatar":   currentUser.Avatar,
	}, "获取用户数据成功")
}
