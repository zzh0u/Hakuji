package database

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// var newBook = []Book{
// 	{ISBN: "9787544244398", Title: "All about Lily Chou-chou", Author: "Shunji Iwai"},
// 	{},
// }

// 数据库初始化相关函数
func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=zzh0u password=123456 dbname=postgresql port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Warn,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                                   newLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	return db, err
}

// func AutoMigrateModels(db *gorm.DB) error {
// 	// 自动迁移所有模型（包含被注释的AutoMigrate功能）
// }

// func CreateBook(db *gorm.DB, book *Book) error {
// 	// 创建单个书籍记录
// }

// func BatchCreateBooks(db *gorm.DB, books []Book) (int64, error) {
// 	// 批量创建书籍
// }

// func GetBookByISBN(db *gorm.DB, isbn string) (*Book, error) {
// 	// 通过ISBN查询书籍
// }

// func UpdateBook(db *gorm.DB, book *Book) error {
// 	// 更新书籍信息
// }

// func DeleteBook(db *gorm.DB, id uint) error {
// 	// 根据ID删除书籍
// }

// // 事务处理
// func WithTransaction(db *gorm.DB, txFunc func(*gorm.DB) error) error {
// 	// 事务封装处理
// }

// // 错误处理
// func IsRecordNotFoundError(err error) bool {
// 	// 判断是否为"记录未找到"错误
// }
