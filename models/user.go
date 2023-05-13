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
	Birthday  time.Time `json:"birthday" binding:"-"`
	LastLogin time.Time `json:"last_login" binding:"-"`
	CreatedAt time.Time `json:"created_at" binding:"-"`
	Comments  Comments  `json:"comments"`
	Passages  Passages  `gorm:"many2many:star_passages;"`
}

type Users []*User
