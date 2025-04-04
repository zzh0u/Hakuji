package util

// 初始化 MinIO 配置
var cfg = &MinIOConfig{
	Endpoint:    "localhost:9000",                           // MinIO 服务地址
	AccessKeyID: "rWi02y2gSXQQ2MDDLr3d",                     // compose.yaml 中设置的用户名
	SecretKey:   "6r5OEp5HcOI1vskdpB2qTXp2RPiuqf43Q35GEACv", // compose.yaml 中设置的密码
	UseSSL:      false,                                      // 本地开发禁用SSL
	BucketName:  "bucket",                                   // 对应创建的存储桶
}
