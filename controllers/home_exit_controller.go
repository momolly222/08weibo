package controllers

import (
	"fmt"
	"mystudy/gin/08weibo/models"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// GetSession ...
func GetSession(c *gin.Context) bool {
	session := sessions.Default(c)
	loginuser := session.Get("loginuser")
	// fmt.Println("loginuser:", loginuser)
	if loginuser != nil {
		return true
	} else {
		return false
	}
}

// HomeGet 首页
func HomeGet(c *gin.Context) {
	// 获取 session, 判断用户是否登录
	islogin := GetSession(c)

	tag := c.Query("tag")
	page, _ := strconv.Atoi(c.Query("page"))

	var artList []models.Article
	var hasFooter bool

	if len(tag) > 0 {
		// 按指定的标签搜索
		artList, _ = models.FindArticleWithTag(tag)
		hasFooter = false
	} else {
		if page < 0 {
			page = 0
		}
		artList, _ = models.FindArticleWithPage(page)
		hasFooter = true
	}

	html := models.MakeHomeBlocks(artList, islogin)
	homeFooterPageCode := models.ConfigHomeFooterPageCode(page)

	c.HTML(http.StatusOK, "home.html", gin.H{"Islogin": islogin, "Content": html, "HasFooter": hasFooter, "PageCode": homeFooterPageCode})
}

// ExitGet 退出
func ExitGet(c *gin.Context) {
	// 清楚该用户登陆状态的数据
	session := sessions.Default(c)
	session.Delete("loginuser")
	session.Save()

	fmt.Println("delete session...", session.Get("loginuser"))
	c.Redirect(http.StatusMovedPermanently, "/")
}
