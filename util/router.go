package util

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/api/upload_book", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(400, gin.H{"error": "未收到文件"})
			return
		}
		tmpPath := "/tmp/" + file.Filename
		if err = c.SaveUploadedFile(file, tmpPath); err != nil {
			c.JSON(500, gin.H{"error": "保存临时文件失败"})
			return
		}
		cfg, err := LoadConfig("config/settings.yaml")
		if err != nil {
			c.JSON(500, gin.H{"error": "配置加载失败"})
			return
		}
		fmt.Println("MinIO config 成功解析")

		_, err = NewMinIOClient(cfg)
		if err != nil {
			c.JSON(500, gin.H{"error": "MinIO初始化失败"})
			return
		}
		fmt.Println("MinIO 初始化成功")

		err = cfg.UploadFile(tmpPath)
		if err != nil {
			c.JSON(500, gin.H{"error": "文件上传失败"})
			return
		}
		c.JSON(200, gin.H{"message": "文件上传成功"})
	})
}
