package main

import (
	gin "github.com/gin-gonic/gin"
	"blog/routes"
) 


func main() {

	r := gin.Default()
	// 配置静态文件路径
	// 启动路由
	routes.CollectRoutes(r)
	// 启动服务
	panic(r.Run(":8080"))
}

