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

	// // 上传图像
	// r.POST("/upload", controller.Upload)
	// r.POST("/upload/rich_editor_upload", controller.RichEditorUpload)
	// // 用户信息管理
	// userRoutes := r.Group("/user")
	// userRoutes.Use(middleware.AuthMiddleware())
	// userRoutes.GET("", controller.GetInfo)                         // 验证用户
	// userRoutes.GET("briefInfo/:id", controller.GetBriefInfo)       // 获取用户简要信息
	// userRoutes.GET("detailedInfo/:id", controller.GetDetailedInfo) // 获取用户详细信息
	// userRoutes.PUT("avatar/:id", controller.ModifyAvatar)          // 修改头像
	// userRoutes.PUT("name/:id", controller.ModifyName)              // 修改用户名
	// // 我的收藏

	// colRoutes := r.Group("/collects")
	// colRoutes.Use(middleware.AuthMiddleware())
	// colRoutes.GET(":id", controller.Collects)        // 查询收藏
	// colRoutes.PUT("new/:id", controller.NewCollect)  // 收藏
	// colRoutes.DELETE(":index", controller.UnCollect) // 取消收藏
	// // 我的关注
	// folRoutes := r.Group("/following")
	// folRoutes.Use(middleware.AuthMiddleware())
	// folRoutes.GET(":id", controller.Following)      // 查询关注
	// folRoutes.PUT("new/:id", controller.NewFollow)  // 关注
	// folRoutes.DELETE(":index", controller.UnFollow) // 取消关注
	// // 查询分类
	// r.GET("/category", controller.SearchCategory)         // 查询分类
	// r.GET("/category/:id", controller.SearchCategoryName) // 查询分类名
	// //用户文章的增删查改
	// articleRoutes := r.Group("/article")
	// //articleRoutes.Use(middleware.AuthMiddleware())
	// articleController := controller.NewArticleController()
	// articleRoutes.POST("", middleware.AuthMiddleware(), articleController.Create)      // 发布文章
	// articleRoutes.PUT(":id", middleware.AuthMiddleware(), articleController.Update)    // 修改文章
	// articleRoutes.DELETE(":id", middleware.AuthMiddleware(), articleController.Delete) // 删除文章
	// articleRoutes.GET(":id", articleController.Show)                                   // 查看文章
	// articleRoutes.POST("list", articleController.List)                                 // 显示文章列表
	return r
}
