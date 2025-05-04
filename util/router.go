package util

import (
	"archive/zip"
	"fmt"
	"hakuji/dao"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/api/upload_book", func(c *gin.Context) {
		// 获取上传的文件
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

		// 加载MinIO配置
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

		// 上传文件到MinIO
		objectName := file.Filename
		err = cfg.UploadFile(tmpPath, objectName)
		if err != nil {
			c.JSON(500, gin.H{"error": "文件上传失败"})
			return
		}

		// 获取书籍信息并保存到数据库
		db, err := dao.InitDB()
		if err != nil {
			c.JSON(500, gin.H{"error": "数据库连接失败"})
			return
		}

		// 从表单获取书籍信息
		isbn := c.PostForm("isbn")
		title := c.PostForm("title")
		author := c.PostForm("author")
		publisher := c.PostForm("publisher")
		category := c.PostForm("category")
		contentSummary := c.PostForm("contentSummary")

		// 处理日期
		var publishedDate time.Time
		publishedDateStr := c.PostForm("publishedDate")
		if publishedDateStr != "" {
			publishedDate, err = time.Parse(time.RFC3339, publishedDateStr)
			if err != nil {
				publishedDate = time.Now() // 如果解析失败，使用当前时间
			}
		} else {
			publishedDate = time.Now()
		}

		// 处理封面文件
		var coverURL string

		// 检查是否有上传的封面文件
		coverFile, err := c.FormFile("coverFile")
		if err == nil && coverFile != nil {
			coverTmpPath := "/tmp/" + coverFile.Filename
			if err = c.SaveUploadedFile(coverFile, coverTmpPath); err == nil {
				// 上传封面到MinIO
				coverObjectName := "covers/" + coverFile.Filename
				if err = cfg.UploadFile(coverTmpPath, coverObjectName); err == nil {
					coverURL = cfg.Minio.Endpoint + "/" + cfg.Minio.BucketName + "/" + coverObjectName
				}
			}
		} else {
			// 根据文件类型处理封面
			fileExt := strings.ToLower(filepath.Ext(file.Filename))

			// 如果是EPUB文件，尝试从中提取封面
			if fileExt == ".epub" {
				coverURL = extractCoverFromEpub(tmpPath, cfg)
			}
			// 对于PDF文件使用默认处理方式，不需要额外操作
		}

		// 创建书籍记录
		book := &dao.Book{
			ISBN:           isbn,
			Title:          title,
			Author:         author,
			CoverURL:       coverURL,
			Hash:           "", // 可以添加文件哈希计算
			PreHash:        "",
			Publisher:      publisher,
			PublishedDate:  publishedDate,
			Category:       category,
			ContentSummary: contentSummary,
			Rating:         0,
			DownloadCount:  0,
			CreatedAt:      time.Now(),
		}

		// 保存到数据库
		err = dao.CreateBook(db, book)
		if err != nil {
			c.JSON(500, gin.H{"error": "保存书籍信息到数据库失败"})
			return
		}

		c.JSON(200, gin.H{"message": "文件上传成功，书籍信息已保存", "bookId": book.ID})
	})

	// 新增下载书籍API
	r.GET("/api/download_book", func(c *gin.Context) {
		bookName := c.Query("book_name")
		if bookName == "" {
			c.JSON(400, gin.H{"error": "缺少书籍名称参数"})
			return
		}

		cfg, err := LoadConfig("config/settings.yaml")
		if err != nil {
			c.JSON(500, gin.H{"error": "配置加载失败"})
			return
		}

		_, err = NewMinIOClient(cfg)
		if err != nil {
			c.JSON(500, gin.H{"error": "MinIO初始化失败"})
			return
		}

		localPath := "/Users/zhou/Downloads/" + bookName
		if err := cfg.DownloadFile(bookName, localPath); err != nil {
			c.JSON(500, gin.H{"error": "文件下载失败"})
			return
		}

		c.JSON(200, gin.H{"message": "文件下载成功", "path": localPath})
	})
	// 新增获取书籍列表API
	r.GET("/api/books", func(c *gin.Context) {
		db, err := dao.InitDB()
		if err != nil {
			c.JSON(500, gin.H{"error": "数据库连接失败"})
			return
		}
		books, err := dao.GetAllBooks(db)
		if err != nil {
			c.JSON(500, gin.H{"error": "获取书籍列表失败"})
			return
		}
		c.JSON(200, gin.H{"books": books})
		fmt.Println("获取书籍列表成功")
		fmt.Println(books)
	})
}

// extractCoverFromEpub 从EPUB文件中提取封面图片
func extractCoverFromEpub(epubPath string, cfg *MinIOConfig) string {
	// 打开EPUB文件（实际上是ZIP文件）
	readCloser, err := zip.OpenReader(epubPath)
	if err != nil {
		fmt.Println("打开EPUB文件失败:", err)
		return ""
	}
	defer readCloser.Close()

	// 创建临时目录存放解压文件
	tmpDir := "/tmp/epub_extract_" + filepath.Base(epubPath)
	os.MkdirAll(tmpDir, 0755)
	defer os.RemoveAll(tmpDir) // 处理完成后删除临时目录

	// 遍历ZIP文件中的所有文件
	var coverFile *zip.File
	for _, file := range readCloser.File {
		// 查找文件名中包含"cover"的图片文件
		fileName := strings.ToLower(file.Name)
		if strings.Contains(fileName, "cover") {
			coverFile = file
			break
		}
	}

	// 如果找不到封面文件，返回空字符串
	if coverFile == nil {
		fmt.Println("未在EPUB中找到封面图片")
		return ""
	}

	// 提取封面文件
	coverPath := filepath.Join(tmpDir, filepath.Base(coverFile.Name))
	coverReader, err := coverFile.Open()
	if err != nil {
		fmt.Println("打开封面文件失败:", err)
		return ""
	}
	defer coverReader.Close()

	// 创建本地文件
	coverOutput, err := os.Create(coverPath)
	if err != nil {
		fmt.Println("创建本地封面文件失败:", err)
		return ""
	}
	defer coverOutput.Close()

	// 复制文件内容
	_, err = io.Copy(coverOutput, coverReader)
	if err != nil {
		fmt.Println("复制封面文件内容失败:", err)
		return ""
	}

	// 上传封面到MinIO
	coverObjectName := "covers/" + filepath.Base(coverFile.Name)
	err = cfg.UploadFile(coverPath, coverObjectName)
	if err != nil {
		fmt.Println("上传封面到MinIO失败:", err)
		return ""
	}

	// 返回封面URL
	return cfg.Minio.Endpoint + "/" + cfg.Minio.BucketName + "/" + coverObjectName
}
