package util

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// LoadConfig 从YAML文件加载配置
func LoadConfig(path string) (*MinIOConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %w", err)
	}

	var config MinIOConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("解析YAML失败: %w", err)
	}
	return &config, nil
}
