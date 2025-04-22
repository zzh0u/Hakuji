package util

import (
	"hakuji/dao"
	"github.com/gin-gonic/gin"
)

func Service() *gin.Engine {
	result, _ := dao.InitDB()
	result.AutoMigrate(&dao.User{}, &dao.Book{}, &dao.Favorite{}, &dao.Comment{}, &dao.Download{}, &dao.Borrow{})
	
	r := gin.Default()
	SetupRouter(r)
	return r
}
