package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"blog/models"
)

var DB *gorm.DB

func Initmysql() {
	// 连接数据库
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	// 建表
	err = DB.AutoMigrate(models.User{},models.Passage{}, models.Comment{}, models.Category{}, models.Tag{}, models.TagPassage{}, models.Star{})
	if err != nil {
		panic(err)
	}
	
	var root models.User
	
	err = DB.First(&root).Error
	if err != nil {
		panic(err)
	}

	root = models.User{
		Username:       "root",
		Password:       "rootroot123",
		IsAdmin:        true,
	}

	err = DB.Create(&root).Error
	if err != nil {
		panic(err)
	}
}
