package main

import (
	"log"
	"path/filepath"
	"sync"

	"hakuji/util"
)

func main() {
	cfg := &util.MinIOConfig{
		Endpoint:    "localhost:9000",                           // MinIO 服务地址
		AccessKeyID: "rWi02y2gSXQQ2MDDLr3d",                     // compose.yaml 中设置的用户名
		SecretKey:   "6r5OEp5HcOI1vskdpB2qTXp2RPiuqf43Q35GEACv", // compose.yaml 中设置的密码
		UseSSL:      false,                                      // 本地开发禁用SSL
		BucketName:  "bucket",                                   // 对应创建的存储桶
	}

	_, err := util.NewMinIOClient(cfg)
	if err != nil {
		log.Fatalf("Init MinIO client failed: %v", err)
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		filePath := filepath.Join("/Users/zhou/Downloads", "3.epub")
		if err := cfg.UploadFolder(filePath); err != nil {
			log.Printf("File upload failed: %v", err)
			// log.Fatalf("File upload failed: %v", err) // 避免静态终止？？？
			return
		}
		log.Printf("File successfully upload to bucket: %s", cfg.BucketName)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		// if err = cfg.DownloadFile("filename.format", "dir/to/your/path/filename.format"); err != nil {
		if err = cfg.DownloadFile("产品需求文档写作指南.pdf", "/Users/zhou/Downloads/产品需求文档写作指南.pdf"); err != nil {
			log.Fatalf("File download failed: %v", err)
			return
		}
		// log.Printf("File download successfully to local: %s", "dir/to/your/path/")
		log.Printf("File download successfully to local: %s", "/Users/zhou/Downloads/")
	}()

	wg.Wait()
}
