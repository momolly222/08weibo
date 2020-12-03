package models

import (
	"fmt"
	"mystudy/gin/08weibo/mydatabase"
)

// Photo 创建相片结构体
type Photo struct {
	Id         int
	Filepath   string
	Filename   string
	Status     int
	Createtime int64
}

// InsertPhoto 插入图片到数据库
func InsertPhoto(photo Photo) (int64, error) {
	return mydatabase.ModifyDB("insert into photo(filepath, filename, status, createtime) values (?,?,?,?)",
		photo.Filepath, photo.Filename, photo.Status, photo.Createtime)
}

// QueryPhoto 查看图片
func QueryPhoto() ([]Photo, error) {
	rows, err := mydatabase.QueryDB("select id, filepath, filename, status, createtime from photo")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var photoList []Photo
	for rows.Next() {
		id := 0
		filepath := ""
		filename := ""
		status := 0
		var createtime int64
		createtime = 0
		rows.Scan(&id, &filepath, &filename, &status, &createtime)
		photo := Photo{id, filepath, filename, status, createtime}
		photoList = append(photoList, photo)
	}
	return photoList, nil
}
