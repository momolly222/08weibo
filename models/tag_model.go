package models

import (
	"log"
	"mystudy/gin/08weibo/mydatabase"
	"strings"
)

// ----------------------- 标签 ----------------------

// QueryTags 查询所有标签，返回一个字段的列表
func QueryTags() []string {
	rows, err := mydatabase.QueryDB("select tags from article")
	if err != nil {
		log.Println(err)
	}
	var tagsList []string
	for rows.Next() {
		tag := ""
		rows.Scan(&tag)
		tagsList = append(tagsList, tag)
	}
	return tagsList
}

// HandleTagsListData ...
func HandleTagsListData(tagList []string) map[string]int {
	var tagsMap = make(map[string]int)
	for _, tag := range tagList {
		sometags := strings.Split(tag, " ")
		for _, value := range sometags {
			tagsMap[value]++
		}
	}
	return tagsMap
}
