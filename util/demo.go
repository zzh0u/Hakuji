package util

import (
	"log"
	"path/filepath"
)

func demo() {
	// 初始化 MinIO 配置
	cfg := &MinIOConfig{
		Endpoint:    "localhost:9000",                           // MinIO 服务地址
		AccessKeyID: "rWi02y2gSXQQ2MDDLr3d",                     // compose.yaml 中设置的用户名
		SecretKey:   "6r5OEp5HcOI1vskdpB2qTXp2RPiuqf43Q35GEACv", // compose.yaml 中设置的密码
		UseSSL:      false,                                      // 本地开发禁用SSL
		BucketName:  "bucket",                                   // 对应创建的存储桶
	}

	// 创建 MinIO 客户端
	client, err := NewMinIOClient(cfg)
	if err != nil {
		log.Fatalf("初始化MinIO客户端失败: %v", err)
	}
	_ = client // 实际使用时可直接用 cfg.Client 操作

	// 上传示例文件
	filePath := filepath.Join("/Users/zhou/Downloads", "产品需求文档写作指南.pdf")
	if err := cfg.UploadFile(filePath); err != nil {
		log.Fatalf("文件上传失败: %v", err)
	}

	log.Printf("文件成功上传至存储桶: %s", cfg.BucketName)
}
