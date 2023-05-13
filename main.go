package main

import (
	"blog/models"
	"blog/routes"
	gin "github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	db, err := models.InitDB()
	// 延迟关闭数据库
	if err != nil {
		log.Println(err)
		return
	}
	a, _ := db.DB()
	defer a.Close()
	r := gin.Default()
	// 配置静态文件路径
	r.StaticFS("/images", http.Dir("./static/images"))
	// 启动路由
	routes.CollectRoutes(r)
	// 启动服务
	panic(r.Run(":8080"))
}
