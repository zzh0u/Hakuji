package util

import (
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// MinIOClient 封装 MinIO 客户端
type MinIOClient struct {
	client *minio.Client
	bucket string
}

// NewMinIOClient 初始化 MinIO 客户端
func NewMinIOClient(endpoint, accessKey, secretKey, bucket string, useSSL bool) (*MinIOClient, error) {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}
	return &MinIOClient{client: client, bucket: bucket}, nil
}

// UploadFile 上传文件到 MinIO
func (m *MinIOClient) UploadFile(ctx context.Context, objectName, filePath string) error {
	_, err := m.client.FPutObject(ctx, m.bucket, objectName, filePath, minio.PutObjectOptions{})
	return err
}
