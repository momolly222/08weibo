package controllers

import (
	"fmt"
	"log"
	"mime/multipart"
	"mystudy/gin/08weibo/models"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

var f *multipart.FileHeader

// 获取相册页面
func PhotoGet(c *gin.Context) {
	islogin := GetSession(c)

	photo, _ := models.QueryPhoto()
	// fmt.Println(photo)

	c.HTML(http.StatusOK, "photo.html", gin.H{"Islogin": islogin, "photo": photo})
}

func PhotoUploadPost(c *gin.Context) {
	fmt.Println("上传文件中...")

	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err)
		return
	}
	f = file
	fmt.Println("f:", f.Filename, f.Size)
	fmt.Println("name:", file.Filename, file.Size)

	now := time.Now()

	// 判断后缀为图片的文件，如果是图片我们才存入到数据库中
	fmt.Println("ext:", filepath.Ext(file.Filename)) // 获取后缀
	fileType := "other"
	fileExt := filepath.Ext(file.Filename)
	if fileExt == ".jpg" || fileExt == ".png" || fileExt == ".gif" || fileExt == ".jpeg" {
		fileType = "img"
	}

	// 文件夹路径
	fileDir := fmt.Sprintf("static/upload/%s/%d/%d/%d", fileType, now.Year(), now.Month(), now.Day())

	// ModePerm是0777，这样拥有该文件夹路径的执行权限
	err = os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		log.Println(err)
		return
	}

	// 文件路径
	timeStamp := time.Now().Unix()
	fileName := fmt.Sprintf("%d-%s", timeStamp, file.Filename)
	// filePathStr := filepath.Join(fileDir, fileName)
	filePathStr := fileDir + fileName

	// 将浏览器客户端上传的文件拷贝到本地路径的文件里面， 此处也可以使用 io 操作
	c.SaveUploadedFile(file, filePathStr)

	if fileType == "img" {
		photo := models.Photo{0, filePathStr, fileName, 0, timeStamp}
		_, err := models.InsertPhoto(photo)
		if err != nil {
			fmt.Println(err)
		}
	}

	c.JSON(http.StatusOK, gin.H{"code": 1, "message": "上传相片成功"})
}
