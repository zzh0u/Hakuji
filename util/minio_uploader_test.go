package util_test

import (
	"context"
	"hakuji/util"
	"os"
	"testing"
)

func TestUploadFile(t *testing.T) {
	ctx := context.Background()

	// 初始化模拟客户端
	mockClient, err := util.NewMinIOClient(
		"localhost:9000",
		"4Tnp1kcq7J6tCWbEf6KR",
		"nZt4gRgBdhCogDwO5kT4NuXqh392sGZywE3IdI7S",
		"firstbucket",
		false,
	)
	if err!= nil {
		t.Fatalf("初始化模拟客户端失败: %v", err)
	}

	// 创建临时测试文件
	tmpFile, err := os.CreateTemp("", "test-file-*.txt")
	if err != nil {
		t.Fatalf("创建临时文件失败: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.WriteString("test content")
	tmpFile.Close()

	// 测试正常上传
	t.Run("UploadSuccess", func(t *testing.T) {
		err := mockClient.UploadFile(ctx, "test-object.txt", tmpFile.Name())
		if err != nil {
			t.Errorf("预期上传成功，但失败: %v", err)
		}
	})

	// 测试空文件名
	t.Run("UploadEmptyFileName", func(t *testing.T) {
		err := mockClient.UploadFile(ctx, "", tmpFile.Name())
		if err == nil {
			t.Error("预期上传失败（空对象名），但成功")
		}
	})
}
