package main

import (
	"database/sql"
	"fmt"
	"log"

	// _ 是为了让编译器导入包，但不将其作为当前文件的导出成员。
	_ "github.com/go-sql-driver/mysql" 
)

func main() {
	// 配置数据库连接字符串
	dsn := "zhou:030507@/BlockChain"

	// 打开数据库连接
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 测试数据库连接
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// 执行查询
	rows, err := db.Query("SELECT * FROM your_table_name")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 遍历结果
	for rows.Next() {
		// 处理每一行数据
		// 这里需要根据你的表结构来定义变量
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}

	// 检查是否有错误发生
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}