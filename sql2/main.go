package main

//go connect mysql

import (
	// "database/sql"
	"fmt"
	"sql2/sqlquery"

	_ "github.com/go-sql-driver/mysql"
	// "sql2/sqlquery"
)

func main() {
	err := sqlquery.InitDB()
	if err != nil {
		fmt.Println("init DB failed,err:%v\n", err)
	}
	fmt.Println("连接数据库成功")

	sqlquery.Insert()
}
