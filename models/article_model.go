package models

import (
	"fmt"
	"mystudy/gin/08weibo/config"
	mydatabase "mystudy/gin/08weibo/mydatabase"
	"strconv"
)

// Article 创建文章结构体
type Article struct {
	Id         int
	Title      string
	Tags       string
	Short      string // 简介
	Content    string
	Author     string
	Createtime int64
}

// InsertArticle 插入文章
func insertArticle(article Article) (int64, error) {
	return mydatabase.ModifyDB("insert into article(title, tags, short, content, author, createtime) values (?,?,?,?,?,?)",
		article.Title, article.Tags, article.Short, article.Content, article.Author, article.Createtime)
}

// AddArticle 增加新的博客文章
func AddArticle(article Article) (int64, error) {
	i, err := insertArticle(article)
	SetArticleRowsNum()
	return i, err
}

// FindArticleWithPage 根据页码查询文章
func FindArticleWithPage(page int) ([]Article, error) {
	// fmt.Println("----------->page", page)

	// 从配置文件中获取每页的文章数量
	return QueryArticleWithPage(page, config.NUM)
}

// QueryArticleWithPage 根据页码显示 从第几条到第几条 的文章
func QueryArticleWithPage(page, num int) ([]Article, error) {
	sql := fmt.Sprintf("limit %d, %d", page*num, num)
	return QueryArticleWithCon(sql)
}

// QueryArticleWithCon 返回文章列表
func QueryArticleWithCon(sql string) ([]Article, error) {
	sql = "select id, title, tags, short, content, author, createtime from article " + sql
	fmt.Println(sql)
	rows, err := mydatabase.QueryDB(sql)
	if err != nil {
		return nil, err
	}
	var artList []Article
	for rows.Next() {
		id := 0
		title := ""
		tags := ""
		short := ""
		content := ""
		author := ""
		var createtime int64
		createtime = 0
		rows.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
		art := Article{id, title, tags, short, content, author, createtime}
		artList = append(artList, art)
	}
	return artList, nil
}

// QueryArticleWithId 根据 id 查询文章
func QueryArticleWithId(id int) Article {
	row := mydatabase.QueryRowDB("select id, title, tags, short, content, author, createtime from article where id=" + strconv.Itoa(id))
	title := ""
	tags := ""
	short := ""
	content := ""
	author := ""
	var createtime int64
	createtime = 0
	row.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
	art := Article{id, title, tags, short, content, author, createtime}
	return art
}

// FindArticleWithTag 根据 tag 查询文章
func FindArticleWithTag(tag string) ([]Article, error) {
	sql := "where tags like '% " + tag + " %'" // 前后均有其他标签， %占位符
	sql += "or tags like '% " + tag + "'"      // 前有其他标签
	sql += "or tags like '" + tag + " %'"      // 后有其他标签，%代表后面还接其他标签，sql译为 or tags like 'tag %'
	sql += "or tags like '" + tag + "'"        // 前后均没有标签
	fmt.Println(sql)
	return QueryArticleWithCon(sql)
}

// -------------------- 翻页 ------------------

// 存储表的行数，只有自己可以更改，当文章新增或者删除时需要更新这个值
var articleRowsNum = 0

// GetArticleRowsNum 只有首次获取行数的时候采取统计表里的行数
func GetArticleRowsNum() int {
	if articleRowsNum == 0 {
		articleRowsNum = QueryArticleRowNum()
	}
	return articleRowsNum
}

// QueryArticleRowNum 查询文章的总条数
func QueryArticleRowNum() int {
	// count(id) 统计指定列（这里为id列）的总记录数
	row := mydatabase.QueryRowDB("select count(id) from article")
	num := 0
	row.Scan(&num)
	return num
}

// SetAritcleRowsNum 设置页数
func SetArticleRowsNum() {
	articleRowsNum = QueryArticleRowNum()
}

// ------------------ 修改文章数据 --------------------

// UpdateArticle 更新文章
func UpdateArticle(article Article) (int64, error) {
	// 数据库操作
	return mydatabase.ModifyDB("update article set title=?, tags=?, short=?, content=? where id=?",
		article.Title, article.Tags, article.Short, article.Content, article.Id)
}

// DeleteAritcleWithId 根据 id 删除文章
func DeleteArticleWithId(artID int) (int64, error) {
	i, err := mydatabase.ModifyDB("delete from article where id=?", artID)
	SetArticleRowsNum()
	return i, err
}
