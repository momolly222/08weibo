package utils

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"html/template"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
	"github.com/sourcegraph/syntaxhighlight"
)

// 传入的数据不一样，那么MD5后的32位长度的数据肯定会不一样
func MD5(str string) string {
	md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	return md5str
}

// SwitchTimeStampToData 把数据库存储的 int64 时间戳 转换为格式化时间字符串展示在页面上
func SwitchTimeStampToData(createtime int64) string {
	str := time.Unix(createtime, 0).Format("2006-01-02 15:04:05")
	return str
}

// SwitchMarkdownToHtml Markdown
func SwitchMarkdownToHtml(content string) template.HTML {
	markdown := blackfriday.MarkdownCommon([]byte(content))

	// 获取到 html 文档
	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(markdown))
	/**
	对 document 进程查询，选择器和 css 的语法一样
	第一个参数： i 是查询到的第几个元素
	第二个参数： selection 就是查询到的元素
	**/
	doc.Find("code").Each(func(i int, selection *goquery.Selection) {
		light, _ := syntaxhighlight.AsHTML([]byte(selection.Text()))
		selection.SetHtml(string(light))
		fmt.Println(selection.Html())
		fmt.Println("light:", string(light))
		fmt.Println("\n\n\n")
	})
	htmlString, _ := doc.Html()
	return template.HTML(htmlString)
}
