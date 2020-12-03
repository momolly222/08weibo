package models

import (
	"fmt"
	"mystudy/gin/08weibo/mydatabase"
)

// User 创建用户结构体
type User struct {
	Id         int
	Username   string
	Password   string
	Status     int // 0 正常状态, 1 删除
	Createtime int64
}

// InsertUser 插入用户
func InsertUser(user User) (int64, error) {
	return mydatabase.ModifyDB("insert into user(username, password, status, createtime) values (?,?,?,?)",
		user.Username, user.Password, user.Status, user.Createtime)
}

// QueryUserWightCon 按条件查询
func QueryUserWightCon(con string) int {
	sql := fmt.Sprintf("select id from user %s", con)
	fmt.Println(sql)
	row := mydatabase.QueryRowDB(sql)
	// id := 0
	var id int
	row.Scan(&id)
	return id
}

// QueryUserWithUsername 根据用户名查询 id
func QueryUserWithUsername(username string) int {
	sql := fmt.Sprintf("where username='%s'", username)
	return QueryUserWightCon(sql)
}

// QueryUserWithParam 根据用户名、密码查询用户，必须两者均正确，才返回用户 id
func QueryUserWithParam(username, password string) int {
	sql := fmt.Sprintf("where username='%s' and password='%s'", username, password)
	return QueryUserWightCon(sql)
}
