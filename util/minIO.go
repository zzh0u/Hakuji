package util

import (
	"context"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinIOConfig struct {
	Client *minio.Client // 导出字段以便外部访问

	Endpoint    string
	AccessKeyID string
	SecretKey   string
	UseSSL      bool

	BucketName string
}

// NewMinIOClient 创建并返回配置好的MinIO客户端
func NewMinIOClient(cfg *MinIOConfig) (*minio.Client, error) {
	client, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKeyID, cfg.SecretKey, ""),
		Secure: cfg.UseSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("MinIO连接失败: %w", err)
	}

	ctx := context.Background()
	exists, err := client.BucketExists(ctx, cfg.BucketName)
	if !exists && err == nil {
		// 创建新存储桶，使用空结构体表示默认配置
		err = client.MakeBucket(ctx, cfg.BucketName, minio.MakeBucketOptions{})
	}

	if err != nil {
		return nil, fmt.Errorf("存储桶初始化失败: %w", err)
	}

	cfg.Client = client
	return client, nil
}

func (cfg *MinIOConfig) UploadFile(filePath string, objectName ...string) error {
	if cfg.Client == nil {
		return fmt.Errorf("MinIO客户端未初始化")
	}

	// 获取对象名称（支持可选参数）
	name := filepath.Base(filePath)
	if len(objectName) > 0 {
		name = objectName[0]
	}

	// 上传文件（核心API调用）
	_, err := cfg.Client.FPutObject(
		context.Background(),
		cfg.BucketName,
		name,
		filePath,
		minio.PutObjectOptions{
			ContentType: "application/octet-stream", // 默认二进制类型
		},
	)

	if err != nil {
		return fmt.Errorf("文件上传失败: %w", err)
	}
	return nil
}

func (cfg *MinIOConfig) DownloadFile(objectName string, localFilePath string) error {
	if cfg.Client == nil {
		return fmt.Errorf("MinIO客户端未初始化")
	}

	if err := os.MkdirAll(filepath.Dir(localFilePath), 0755); err != nil {
		return fmt.Errorf("创建本地目录失败: %w", err)
	}

	err := cfg.Client.FGetObject(
		context.Background(),
		cfg.BucketName,           // 从配置中获取存储桶名
		objectName,               // 必需参数：存储桶中的对象名称
		localFilePath,            // 必需参数：本地保存路径
		minio.GetObjectOptions{}, // 下载选项（可空）
	)

	if err != nil {
		return fmt.Errorf("文件下载失败: %w", err)
	}
	return nil
}

func (cfg *MinIOConfig) UploadFolder(localPath string) error {
	if cfg.Client == nil {
		return fmt.Errorf("MinIO客户端未初始化")
	}

	return filepath.WalkDir(localPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err // 跳过目录和错误文件
		}

		// 生成相对路径作为对象名称
		// 重要！！！应提取上传路径的最后一部分
		relPath, _ := filepath.Rel(localPath, path)
		objectName := filepath.Join("./", relPath)

		// 上传单个文件
		_, err = cfg.Client.FPutObject(
			context.Background(),
			cfg.BucketName,
			objectName,
			path,
			minio.PutObjectOptions{},
		)
		return err
	})
}

func (cfg *MinIOConfig) DownloadFolder(minioPrefix string, localPath string) error {
	if cfg.Client == nil {
		return fmt.Errorf("MinIO客户端未初始化")
	}

	if err := os.MkdirAll(localPath, 0755); err != nil {
		return err
	}

	// 递归列出所有对象
	objCh := cfg.Client.ListObjects(context.Background(), cfg.BucketName, minio.ListObjectsOptions{
		Prefix:    minioPrefix,
		Recursive: true,
	})

	for obj := range objCh {
		if obj.Err != nil {
			continue
		}

		relPath := strings.TrimPrefix(obj.Key, minioPrefix)
		localFilePath := filepath.Join(localPath, relPath)

		if err := os.MkdirAll(filepath.Dir(localFilePath), 0755); err != nil {
			return err
		}

		if err := cfg.Client.FGetObject(
			context.Background(),
			cfg.BucketName,
			obj.Key,
			localFilePath,
			minio.GetObjectOptions{},
		); err != nil {
			return err
		}
	}
	return nil
}
