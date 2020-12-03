package controllers

import (
	"fmt"
	"mystudy/gin/08weibo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TagsGet(c *gin.Context) {
	islogin := GetSession(c)

	tags := models.QueryTags()
	fmt.Println(tags)
	fmt.Println(models.HandleTagsListData(tags))

	c.HTML(http.StatusOK, "tags.html", gin.H{"Tags": models.HandleTagsListData(tags), "Islogin": islogin})
}
