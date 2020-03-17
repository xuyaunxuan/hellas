package models

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"hellas/common/setting"
)

var db *gorm.DB

func Setup() {
	var err error
	// 数据库链接
	db, err = gorm.Open(setting.DataBaseType, setting.DataBasePath)
	if err != nil {
		log.Println(err)
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}