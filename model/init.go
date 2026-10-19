package model

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dsn string) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("数据库连接成功...")
	// err = db.AutoMigrate() // 建表
	if err != nil {
		log.Println(err)
	}
	log.Println("表结构初始化成功...")
	DB = db
}
