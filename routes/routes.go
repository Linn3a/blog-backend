package routes

import (
	"blog/controller"
	"blog/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoutes(r *gin.Engine) *gin.Engine {
	// 允许跨域访问
	r.Use(middleware.CORSMiddleware())
	// 注册
	r.POST("/user", controller.Register)
	// 登录
	r.POST("/login", controller.Login)
	r.POST("/autologin", controller.Autologin)
	r.PUT("user/:id", controller.UpdateUser)
	r.DELETE("/user/:id", controller.DeleteUser)
	r.GET("/user", controller.GetAllUsers)
	r.GET("user/:id", controller.GetUser)
	r.PUT("admin/:id", controller.ChangeAdminState)

	r.GET("/cate", controller.GetAllCates)
	r.GET("/cate/:id", controller.GetCate)
	r.POST("/cate", controller.CreateCate)
	r.PUT("/cate/:id", controller.UpdateCate)
	r.DELETE("/cate/:id", controller.DeleteCate)

	r.GET("/p", controller.GetAllPassages)
	r.GET("/p/:id", controller.GetPassageContent)
	r.POST("/p", controller.CreatePassage)
	r.DELETE("/p/:id", controller.DeletePassage)
	r.PUT("/p/:id", controller.UpdatePassage)

	r.GET("/tag", controller.GetAllTags)
	r.GET("/tag/:id", controller.GetTag)
	r.POST("/tag", controller.CreateTag)
	r.DELETE("/tag/:id", controller.DeleteTag)

	r.POST("/comment", controller.CreateComment)
	r.DELETE("/comment/:id", controller.DeleteComment)

	r.POST("/upload", controller.Upload)

	return r
}
