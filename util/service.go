package util

import "hakuji/dao"

func Service() {
	result, _ := dao.InitDB()
	result.AutoMigrate(&dao.User{}, &dao.Book{}, &dao.Favorite{}, &dao.Comment{}, &dao.Download{}, &dao.Borrow{})	
}
