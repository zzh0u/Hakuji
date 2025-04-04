package dao

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 数据库初始化相关函数
func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=zzh0u password=123456 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
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

	// cleanData(db)

	return db, err
}

func AutoMigrateModels(db *gorm.DB) error {
	return db.AutoMigrate(&Book{}, &User{}, &Favorite{}, &Comment{}, &Download{}, &Borrow{})
}

// 有没有比较通用一些的检索方法，譬如传入一个 book *Book，然后返回所有满足条件的书
// 但怎么知道是通过 book 中的哪一个参数搜索

// CRUD operations
// get all books
func GetAllBooks(db *gorm.DB) ([]Book, error) {
	var books []Book
	result := db.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

// create a new book
func CreateBook(db *gorm.DB, book *Book) error {
	result := db.Create(book)
	fmt.Println(result.RowsAffected)
	return result.Error
}

// create books in batches
func BatchCreateBooks(db *gorm.DB, books []Book) error {
	result := db.CreateInBatches(books, 40)
	return result.Error
}

// search book via ISBN
func GetBookByISBN(db *gorm.DB, isbn string) (*Book, error) {
	var book Book
	result := db.Where("isbn = ?", isbn).Find(&book)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}

// update book info
func UpdateBook(db *gorm.DB, book *Book) error {
	// 更新书籍信息
	result := db.Save(book)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// delete book via id
func DeleteBook(db *gorm.DB, id uint) error {
	now := time.Now()
	result := db.Model(&Book{}).Where("id = ", id).Update("deleted_at", now)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 测试时使用，保证每次启动服务都不会有上次的数据
func cleanData(db *gorm.DB) error {
	models := []interface{}{
		&Comment{},
		&Favorite{},
		&Download{},
		&Borrow{},
		&Book{},
		&User{},
	}

	return db.Transaction(func(tx *gorm.DB) error {
		for _, model := range models {
			err := tx.Unscoped().Where("1 = 1").Delete(model).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}
