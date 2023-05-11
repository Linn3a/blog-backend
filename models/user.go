package models

import (
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        uint      `json:"id" binging:"-"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Avatar    string    `json:"avatar" binding:"-"`
	Desc      string    `json:"desc" binding:"-"`
	Gender    uint      `json:"gender" binding:"-"`
	IsAdmin   bool      `json:"is_admin" binding:"-"`
	Birthday  time.Time `json:"birthday" binding:"-"`
	LastLogin time.Time `json:"last_login" binding:"-"`
	CreatedAt time.Time `json:"created_at" binding:"-"`
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
