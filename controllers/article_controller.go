package controllers

import (
	"fmt"
	"mystudy/gin/08weibo/models"
	"mystudy/gin/08weibo/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// AddArticleGet 获取写博客的页面
func AddArticleGet(c *gin.Context) {
	islogin := GetSession(c)
	c.HTML(http.StatusOK, "write_article.html", gin.H{"Islogin": islogin})
}

func AddArticlePost(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("loginuser")

	// 获取浏览器传输的数据，通过表单的 name 属性获取值
	// 获取表单信息
	title := c.PostForm("title")
	tags := c.PostForm("tags")
	short := c.PostForm("short")
	content := c.PostForm("content")
	fmt.Printf("title: %s, tags: %s\n", title, tags)

	// 实例化 model， 将它输入到数据库中
	art := models.Article{0, title, tags, short, content, user.(string), time.Now().Unix()}
	_, err := models.AddArticle(art)

	// 返回数据给浏览器
	response := gin.H{}
	if err == nil {
		// 无误
		response = gin.H{"code": 1, "message": "ok"}
	} else {
		response = gin.H{"code": 0, "message": "error"}
	}

	c.JSON(http.StatusOK, response)
}

// ShowArticleGet 文章详情页
func ShowArticleGet(c *gin.Context) {
	islogin := GetSession(c)
	idStr := c.Param("id") // 或 c.Query("id")
	id, _ := strconv.Atoi(idStr)
	fmt.Println("id:", id)

	// 获取 id 所对应的文章信息
	art := models.QueryArticleWithId(id)
	// 渲染 HTML
	c.HTML(http.StatusOK, "show_article.html", gin.H{"Islogin": islogin, "Title": art.Title, "Content": utils.SwitchMarkdownToHtml(art.Content), "Author": art.Author, "Tags": art.Tags})
}

// UpdateArticleGet 修改更新文章的获取
func UpdateArticleGet(c *gin.Context) {
	// 获取 session
	islogin := GetSession(c)
	idstr := c.Query("id") // 或 c.Param("id")
	id, _ := strconv.Atoi(idstr)
	fmt.Println(id)

	// 获取 id 所对应的文章信息
	art := models.QueryArticleWithId(id)

	c.HTML(http.StatusOK, "write_article.html", gin.H{"Islogin": islogin, "Title": art.Title, "Tags": art.Tags, "Short": art.Short, "Content": art.Content, "Id": art.Id})
}

// UpdateArticlePost 修改更新文章
func UpdateArticlePost(c *gin.Context) {
	idstr := c.Query("id")
	id, _ := strconv.Atoi(idstr)
	fmt.Println("postid: ", id)

	title := c.PostForm("title")
	tags := c.PostForm("tags")
	short := c.PostForm("short")
	content := c.PostForm("content")

	art := models.Article{id, title, tags, short, content, "", 0}
	_, err := models.UpdateArticle(art)

	// 返回数据给浏览器
	respone := gin.H{}
	if err == nil {
		respone = gin.H{"code": 1, "message": "文章更新成功"}
	} else {
		respone = gin.H{"code": 0, "message": "文章更新失败"}
	}

	c.JSON(http.StatusOK, respone)
}

// DeleteAtricleGet 点击删除后重定向到首页
func DeleteArticleGet(c *gin.Context) {
	islogin := GetSession(c)

	idstr := c.Query("id")
	id, _ := strconv.Atoi(idstr)
	fmt.Println("id: ", id)

	_, err := models.DeleteArticleWithId(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"Islogin": islogin, "code": 0, "message": "文章删除失败"})
	} else {
		c.Redirect(http.StatusMovedPermanently, "/")
	}
}
