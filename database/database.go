package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Book struct {
	gorm.Model
	Isbn   string `gorm:"uniqueIndex"`
	Title  string
	Author string
}

func main() {
	dsn := "host=localhost user=zzh0u password=123456 dbname=postgresql port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	// _ = db.AutoMigrate(&Book{}) //此处会有sql语句

	newBook := Book{Isbn: "9787544244398", Title: "All about Lily Chou-chou", Author: "Shunji Iwai"}

	db.Create(&newBook)

	fmt.Println(newBook.ID)

	books := []Book{
		{Isbn: "9787010181141", Title: "毛泽东选集", Author: "毛泽东"},
		{Isbn: "9787520207669", Title: "Little Prince", Author: "安托万·德·圣埃克苏佩里"},
	}
	result := db.Create(&books)

	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)
}
