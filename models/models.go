package models

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"hellas/common/setting"
)

var db *gorm.DB

func Setup() {
	var (
		err error
		dbType, path string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String()
	//dbName = sec.Key("NAME").String()
	path = sec.Key("PATH").String()
	db, err = gorm.Open(dbType, path)

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