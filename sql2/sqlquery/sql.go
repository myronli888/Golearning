package sqlquery

import (
	"database/sql"
	"fmt"
)

type User struct {
	Id   int
	Name string
	Age  int
}

var db *sql.DB //连接池对象.

func InitDB() (err error) {
	//数据库信息,字符串存储
	dsn := "root:ebopark@2020@tcp(49.4.2.102:3306)/sql_test"
	//连接数据库
	db, err = sql.Open("mysql", dsn)
	// defer db.Close()
	if err != nil {
		return
	}
	return
}

func Query(n int) {
	var u1 User
	sqlStr := `select id,name,age from user where id=?;`
	//执行 使用db.Query()来发送查询到数据库，获取结果集Rows(一般db查询出来的实例化对象叫做row或者rows)，并检查错误。
	rowObj := db.QueryRow(sqlStr, n)
	//拿出结果
	rowObj.Scan(&u1.Id, &u1.Name, &u1.Age)
	//  %#v 结构体格式化输出的一种格式
	fmt.Printf("u1:%#v\n", u1)
}

func QueryMore(n int) {
	var u2 User
	sqlStr := `select id,name,age from user where id>?;`
	// Queryrow是查询一条, Query是查询多条.
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Println("query failed err:%v\n", err)
		return
	}
	defer rows.Close()
	//联系之前定义的结构体user
	// rows.Next(),for 后面判断一个boolean的值进行循环, 当rows还有数据的时候就继续循环
	for rows.Next() {
		err := rows.Scan(&u2.Id, &u2.Name, &u2.Age)
		if err != nil {
			fmt.Println("scan failed, err:%v\n", err)
		}
		fmt.Printf("u2:%#v\n", u2)
	}
}

func Add(a, b int) int {

	return (a + b)
}

func Insert() {
	//执行了插入操作
	// sqlStr := `insert into user(name,age) values("发财33","48")`
	stmt, _ := db.Prepare("insert into user values(default,?,?)")
	_, err := stmt.Exec("老王", "11")
	if err != nil {
		fmt.Println("insert failed, err:%v\n", err)
	}
	fmt.Println("插入成功")
}
