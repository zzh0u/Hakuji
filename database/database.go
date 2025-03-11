package database

import (

	// _ 是为了让编译器导入包，但不将其作为当前文件的导出成员。
	_ "github.com/go-sql-driver/mysql" 
)

func Database() {
}