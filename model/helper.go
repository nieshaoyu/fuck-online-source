package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/course?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Panic("connect db failed")
	}
	if err := DB.DB().Ping(); err != nil {
		log.Panic("ping db failed")
	}
}
