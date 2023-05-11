package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	// 连接数据库
	dsn := "root:123456@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	// 建表
	err = db.AutoMigrate(User{}, Passage{}, Comment{}, Category{}, Tag{}, TagPassage{}, Star{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	bir, _ := time.ParseInLocation("2006-01-02 15:04:05", "2003-06-15 00:00:00", time.Local)
	root := User{
		Username:  "root",
		Password:  "123123",
		Gender:    1,
		Avatar:    "https://img.js.design/assets/img/645650aca30d747a6da0787b.jpg#208a42317334f6b46f2c03fcf9c101bd",
		IsAdmin:   true,
		Birthday:  bir,
		LastLogin: time.Now(),
	}
	err = db.Create(&root).Error
	log.Println(&root.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to create root user: %w", err)
	}

	DB = db
	return db, nil

}

func GetDB() *gorm.DB {
	return DB
}
