package models

import (
	"hellas/dtos"
	"hellas/dtos/user"
	"log"
)


type Login struct {
	Id string
}

type Product struct {
	Code string
	Price uint
}

// 新用户创建
func CreateNewUser(param user.RegisterParameter) user.RegisterResult {
	// 开启事务
	tx := db.Begin()
	var result user.RegisterResult
	var userList []dtos.User
	// 用户ID不为空的时候，做唯一性验证
	if param.AccountId != "" {
		//var userStruct dtos.User
		//userStruct.AccountId = param.AccountId
		//db.
		db.Where(&dtos.User{AccountId: param.AccountId}).Find(&userList)
	}


	// 提交事务
	tx.Commit()

	//db.Debug().Find(&login)
	//log.Printf("%+v", login)
	////
	////db.AutoMigrate(&Product{})
	////
	////db.Create(&Product{Code: "L1212", Price: 1000})
	//var product []Product
	//db.Find(&product) // find product with id 1
	//log.Printf("%+v", product)
	return result
}


func GetUser() []Product {
	var login []Login

	db.Debug().Find(&login)
	log.Printf("%+v", login)
	//
	//db.AutoMigrate(&Product{})
	//
	//db.Create(&Product{Code: "L1212", Price: 1000})
	var product []Product
	db.Find(&product) // find product with id 1
	log.Printf("%+v", product)
	return product
}