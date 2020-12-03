package mydatabase

import (
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB

// InitMysql 初始化连接数据库
func InitMysql() {
	fmt.Println("InitMysql...")
	if db == nil {
		tmpdb, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golearn01")
		if err != nil {
			fmt.Printf("sql open failed, err: %v, tmpdb: %v\n", err, tmpdb)
			return
		}
		db = tmpdb
		CreateTableWithUser()
		CreatTableWithArticle()
		_, err = CreateTableWithPhoto()
		if err != nil {
			fmt.Println(err)
		}
	}
}

// CreateTableWithUser 创建用户表
func CreateTableWithUser() {
	sql := `create table if not exists user(
		id int(4) primary key auto_increment not null,
		username varchar(64),
		password varchar(64),
		status int(4),
		createtime int(10)
		);`

	ModifyDB(sql)
}

// CreatTableWithArticle 创建博客文章表单
func CreatTableWithArticle() {
	sql := `create table if not exists article(
		id int(4) primary key auto_increment not null,
		title varchar(30),
		author varchar(20),
		tags varchar(30),
		short varchar(255),
		content longtext,
		createtime int(10)
		);`

	ModifyDB(sql)
}

// CreateTableWithPhoto 创建相册表单
func CreateTableWithPhoto() (int64, error) {
	sql := `create table if not exists photo(
		id int(4) primary key auto_increment not null,
		filepath varchar(255),
		filename varchar(64),
		status int(4),
		createtime int(10)
		);`

	d, err := ModifyDB(sql)
	return d, err
}

// ModifyDB 操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

// QueryRowDB 单行查询
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}

// QueryDB 多行查询
func QueryDB(sql string) (*sql.Rows, error) {
	return db.Query(sql)
}
