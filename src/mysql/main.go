package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// 是一个连接池操作
var db *sql.DB

func initDB() (err error) {
	// 连接数据库
	dsn := "root:@tcp(127.0.0.1:3306)/gostudy"
	db, err = sql.Open("mysql", dsn)
	// dsn格式不正确
	if err != nil {
		fmt.Printf("open %s failed, err: %v\n", dsn, err)
		return
	}
	// 尝试连接数据库
	err = db.Ping()
	if err != nil {
		fmt.Printf("open %s failed, err: %v\n", dsn, err)
		return
	}
	// 设置数据库连接池的最大连接数
	// db.SetConnMaxLifetime(10)
	return
}

type user struct {
	id   int
	name string
	age  int
}

func query() {
	var u1 user
	// 查询单条sql语句
	sqlStr := `select id, name, age from user where id=?;`
	// 执行
	rowObj := db.QueryRow(sqlStr, 1) // 从连接池里取一个连接去数据库查询数据
	// 取到结果，释放连接
	rowObj.Scan(&u1.id, &u1.name, &u1.age)
	fmt.Printf("u1: %#v\n", u1)
}

func queryMore(n int) {
	sqlStr := `select id, name, age from user where id > ?;`
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Printf("exec %s query failed, err: %v\n", sqlStr, err)
		return
	}

	defer rows.Close()
	for rows.Next() {
		var u1 user
		err := rows.Scan(&u1.id, &u1.name, &u1.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("name: %s age: %d\n", u1.name, u1.age)
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err: %v\n", err)
	}
	fmt.Println("connect db successfully")
}
