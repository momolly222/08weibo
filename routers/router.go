package routers

import (
	"mystudy/gin/08weibo/controllers"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化注册路由
func InitRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("C:/A-NCQ/study/myGo/src/mystudy/gin/08weibo/views/*")
	router.Static("/static", "./static")
	router.StaticFS("/upload", http.Dir("C:/A-NCQ/study/myGo/src/mystudy/gin/08weibo/static"))

	// 设置 session 中间件
	store := cookie.NewStore([]byte("loginuser"))
	router.Use(sessions.Sessions("mysession", store))

	{
		// 注册
		router.GET("/register", controllers.RegisterGet)
		router.POST("/register", controllers.RegisterPost)

		// 登陆
		router.GET("/login", controllers.LoginGet)
		router.POST("/login", controllers.LoginPost)

		// 首页
		router.GET("/", controllers.HomeGet)

		// 退出
		router.GET("/exit", controllers.ExitGet)
	}

	// 文章
	v1 := router.Group("/article")
	{
		// 获取、写文章
		v1.GET("/add", controllers.AddArticleGet)
		v1.POST("/add", controllers.AddArticlePost)

		// 显示文章内容
		v1.GET("/show/:id", controllers.ShowArticleGet)

		// 修改更新文章
		v1.GET("/update", controllers.UpdateArticleGet)
		v1.POST("/update", controllers.UpdateArticlePost)

		v1.GET("/delete", controllers.DeleteArticleGet)
	}

	// 标签
	router.GET("/tags", controllers.TagsGet)

	// 相册
	router.GET("/photo", controllers.PhotoGet)
	router.POST("/photo", controllers.PhotoUploadPost)

	return router
}
