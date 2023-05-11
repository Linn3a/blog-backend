package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	err = DB.AutoMigrate(User{},Passage{}, Comment{}, Category{}, Tag{}, TagPassage{}, Star{})
	if err != nil {
		panic(err)
	}
	
	var root User
	
	err = DB.First(&root).Error
	if err != nil {
		panic(err)
	}

	DB.Save(&User{
		Username:       "root",
		Password:       "rootroot123",
		IsAdmin:        true,
	})
}
