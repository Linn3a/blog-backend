package models

import (
	_ "gorm.io/driver/mysql"
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
	Birthday  time.Time `json:"birthday" `
	LastLogin time.Time `json:"last_login" binding:"-"`
	CreatedAt time.Time `json:"created_at" binding:"-"`
	UserTag   string    `json:"user_tag" binging:"-" gorm:"default:游客"`
	TagColor  string    `json:"tag_color" binging:"-" gorm:"default:#f3f4f6"`
	Comments  Comments  `json:"comments"`
	Passages  Passages  `gorm:"many2many:star_passages;"`
}

type Users []*User
