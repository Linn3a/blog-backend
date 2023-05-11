package models

import (
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        uint      `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Avatar    string    `json:"avatar"`
	Desc      string    `json:"desc"`
	Gender    uint      `json:"gender"`
	IsAdmin   bool      `json:"is_admin"`
	Birthday  time.Time `json:"birthday"`
	LastLogin time.Time `json:"last_login"`
	CreatedAt time.Time `json:"created_at"`
}

type Users []*User

func CreateUser(user User, db *gorm.DB) (err error) {
	err = db.Create(&user).Error
	return err
}

func GetUserById(id int, db *gorm.DB) (user User, err error) {
	err = db.Where("id = ?", id).First(&user).Error
	return user, err
}
