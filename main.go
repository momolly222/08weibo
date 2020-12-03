package main

import (
	mydatabase "mystudy/gin/08weibo/mydatabase"
	routers "mystudy/gin/08weibo/routers"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mydatabase.InitMysql()

	router := routers.InitRouter()
	// 静态资源
	// router.Static("/static", "./static")
	router.Run("127.0.0.1:8080")
}
