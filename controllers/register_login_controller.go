package controllers

import (
	"fmt"
	"mystudy/gin/08weibo/models"
	"mystudy/gin/08weibo/utils"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// RegisterGet 获取注册页
func RegisterGet(c *gin.Context) {
	// 返回 html
	c.HTML(http.StatusOK, "register.html", gin.H{"title": "注册页"})
}

// RegisterPost 处理注册
func RegisterPost(c *gin.Context) {
	// 获取表单信息
	username := c.PostForm("username")
	password := c.PostForm("password")
	repassword := c.PostForm("repassword")
	fmt.Println("打印的内容", username, password, repassword)

	// 注册之前先判断该用户名是否已经被注册，如果已经注册，返回错误
	id := models.QueryUserWithUsername(username)
	fmt.Println("id:", id)
	if id > 0 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "用户名已经存在"})
		return
	}

	// 注册用户名和密码
	// 存储的密码是md5后的数据，那么在登录的验证的时候，也是需要将用户的密码md5之后和数据库里面的密码进行判断
	password = utils.MD5(password)
	fmt.Println("密码md5后：", password)

	user := models.User{0, username, password, 0, time.Now().Unix()}
	_, err := models.InsertUser(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "注册失败"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "注册成功"})
	}
}

// LoginGet 获取登录页
func LoginGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{"title": "登录页"})
}

// LoginPost 登陆账户
func LoginPost(c *gin.Context) {
	// 获取表单信息
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println(username, password)

	id := models.QueryUserWithParam(username, utils.MD5(password))
	fmt.Println("id:", id)
	if id > 0 {

		session := sessions.Default(c)
		session.Set("loginuser", username)
		session.Save() // 无论是 set 一个 session, 还是 delete 一个 session, 都要调用 save 方法进行保存

		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "登陆成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "登陆失败，账户名或密码错误。"})
	}
}
